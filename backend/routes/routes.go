package routes

import (
	"kos-management/controllers"
	"kos-management/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userController *controllers.UserController, kamarController *controllers.KamarController, penyewaController *controllers.PenyewaController, tagihanController *controllers.TagihanController, transaksiController *controllers.TransaksiController) {
	// Ping endpoint for health check and auto-port detection
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Server is running",
		})
	})

	// Public routes
	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Register)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// User
		api.POST("/users", userController.Create)
		api.GET("/users", userController.GetAll)
		api.GET("/users/:id", userController.GetByID)
		api.PUT("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)

		// Kamar
		api.POST("/kamars", kamarController.Create)
		api.GET("/kamars", kamarController.GetAll)
		api.GET("/kamars/:id", kamarController.GetByID)
		api.PUT("/kamars/:id", kamarController.Update)
		api.DELETE("/kamars/:id", kamarController.Delete)

		// Penyewa
		api.POST("/penyewas", penyewaController.Create)
		api.GET("/penyewas", penyewaController.GetAll)
		api.GET("/penyewas/:id", penyewaController.GetByID)
		api.PUT("/penyewas/:id", penyewaController.Update)
		api.DELETE("/penyewas/:id", penyewaController.Delete)

		// Tagihan
		api.POST("/tagihans", tagihanController.Create)
		api.GET("/tagihans", tagihanController.GetAll)
		api.GET("/tagihans/:id", tagihanController.GetByID)
		api.PUT("/tagihans/:id", tagihanController.Update)
		api.DELETE("/tagihans/:id", tagihanController.Delete)

		// Transaksi
		api.POST("/transaksis", transaksiController.Create)
		api.GET("/transaksis", transaksiController.GetAll)
		api.GET("/transaksis/:id", transaksiController.GetByID)
		api.PUT("/transaksis/:id", transaksiController.Update)
		api.DELETE("/transaksis/:id", transaksiController.Delete)
	}
}
