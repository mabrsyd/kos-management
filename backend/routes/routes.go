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

	// Profile routes
	api.POST("/profile", userController.UpdateProfile)

	// User routes
	api.GET("/users", userController.GetAll)
	api.GET("/users/:id", userController.GetByID)
	api.POST("/users", userController.Create)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.Delete)

	// Kamar routes
	api.GET("/kamars", kamarController.GetAll)
	api.GET("/kamars/:id", kamarController.GetByID)
	api.POST("/kamars", kamarController.Create)
	api.PUT("/kamars/:id", kamarController.Update)
	api.DELETE("/kamars/:id", kamarController.Delete)

	// Penyewa routes
	api.GET("/penyewas/active", penyewaController.GetActiveTenants)
	api.GET("/penyewas/with-debt", penyewaController.GetTenantsWithDebt)
	api.GET("/penyewas", penyewaController.GetAll)
	api.GET("/penyewas/:id", penyewaController.GetByID)
	api.GET("/penyewas/:id/bills", penyewaController.GetTenantWithBills)
	api.POST("/penyewas", penyewaController.Create)
	api.PUT("/penyewas/:id", penyewaController.Update)
	api.DELETE("/penyewas/:id", penyewaController.Delete)

	// Tagihan routes
	api.GET("/tagihans/by-penyewa/:id", tagihanController.GetByPenyewaID)
	api.GET("/tagihans/unpaid/by-penyewa/:id", tagihanController.GetUnpaidByPenyewaID)
	api.GET("/tagihans/by-month", tagihanController.GetBillsByMonth)
	api.GET("/tagihans/debt/by-penyewa/:id", tagihanController.GetTenantDebt)
	api.PUT("/tagihans/update-overdue", tagihanController.UpdateOverdueBills)
	api.GET("/tagihans", tagihanController.GetAll)
	api.GET("/tagihans/:id", tagihanController.GetByID)
	api.POST("/tagihans", tagihanController.Create)
	api.PUT("/tagihans/:id/mark-paid", tagihanController.MarkAsPaid)
	api.PUT("/tagihans/:id", tagihanController.Update)
	api.DELETE("/tagihans/:id", tagihanController.Delete)

	// Transaksi routes
	api.GET("/transaksis", transaksiController.GetAll)
	api.GET("/transaksis/:id", transaksiController.GetByID)
	api.POST("/transaksis", transaksiController.Create)
	api.PUT("/transaksis/:id", transaksiController.Update)
	api.DELETE("/transaksis/:id", transaksiController.Delete)
}
