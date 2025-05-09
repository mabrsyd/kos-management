package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Get database connection string from environment variable
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		// Default connection string if not set in environment
		dsn = "host=localhost user=postgres password=1234 dbname=kosmanagement port=5432 sslmode=disable"
		fmt.Printf("Using default database configuration: %s\n", dsn)
	}

	// Configure GORM logger for better debugging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Don't ignore record not found errors
			Colorful:                  true,        // Enable color
		},
	)

	// Try to connect to the database
	fmt.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// If the initial connection fails, attempt to connect to a different database
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		fmt.Println("Attempting to connect to 'postgres' database instead...")

		// Try connecting to the default postgres database
		dsn = "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Fatal("Failed to connect to any database:", err)
		}

		// Create the kosmanagement database if it doesn't exist
		fmt.Println("Creating 'kosmanagement' database...")
		db.Exec("CREATE DATABASE kosmanagement")

		// Now connect to the newly created database
		dsn = "host=localhost user=postgres password=1234 dbname=kosmanagement port=5432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Fatal("Failed to connect to the created database:", err)
		}
	}

	fmt.Println("Database connection established successfully")
	DB = db
}
