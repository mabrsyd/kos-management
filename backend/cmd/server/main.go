package main

import (
	"fmt"
	"kos-management/config"
	"kos-management/controllers"
	"kos-management/models"
	"kos-management/repositories"
	"kos-management/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Connect to database
	config.ConnectDatabase()
	db := config.DB

	// Drop and recreate tables for fresh start
	fmt.Println("Migrating database schemas...")
	db.Migrator().DropTable(
		&models.User{},
		&models.Kamar{},
		&models.Penyewa{},
		&models.Tagihan{},
		&models.Transaksi{},
	)

	// Auto migrate models
	err := db.AutoMigrate(
		&models.User{},
		&models.Kamar{},
		&models.Penyewa{},
		&models.Tagihan{},
		&models.Transaksi{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed admin user if no users exist
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		// Create default admin user
		fmt.Println("Creating admin user...")
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		adminUser := models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Name:     "Administrator",
			Role:     "admin",
		}
		result := db.Create(&adminUser)
		if result.Error != nil {
			log.Fatalf("Failed to create admin user: %v", result.Error)
		}
		fmt.Println("Admin user created successfully")
	}

	// Initialize repositories
	userRepo := &repositories.UserRepository{DB: db}
	kamarRepo := &repositories.KamarRepository{DB: db}
	penyewaRepo := &repositories.PenyewaRepository{DB: db}
	tagihanRepo := &repositories.TagihanRepository{DB: db}
	transaksiRepo := &repositories.TransaksiRepository{DB: db}

	// Initialize controllers
	userController := &controllers.UserController{Repo: userRepo}
	kamarController := &controllers.KamarController{Repo: kamarRepo}
	penyewaController := &controllers.PenyewaController{Repo: penyewaRepo}
	tagihanController := &controllers.TagihanController{Repo: tagihanRepo}
	transaksiController := &controllers.TransaksiController{Repo: transaksiRepo}

	// Setup Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.RegisterRoutes(r, userController, kamarController, penyewaController, tagihanController, transaksiController)

	// Try to start the server on port 8080 first
	fmt.Println("Starting server on :8080")
	err = r.Run(":8080")
	if err != nil {
		// If port 8080 is already in use, try port 8081
		fmt.Println("Port 8080 already in use, trying port 8081 instead")
		err = r.Run(":8081")
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}
}
