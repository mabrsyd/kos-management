package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"kos-management/models"
	"kos-management/repositories"

	"github.com/gin-gonic/gin"
)

type PenyewaController struct {
	Repo        *repositories.PenyewaRepository
	TagihanRepo *repositories.TagihanRepository
	KamarRepo   *repositories.KamarRepository
}

func (c *PenyewaController) Create(ctx *gin.Context) {
	var penyewa models.Penyewa
	if err := ctx.ShouldBindJSON(&penyewa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the tenant
	if err := c.Repo.Create(&penyewa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate monthly bills for the tenant's contract period
	if err := c.generateMonthlyBills(&penyewa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Tenant created but failed to generate bills: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, penyewa)
}

func (c *PenyewaController) GetAll(ctx *gin.Context) {
	penyewas, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penyewas)
}

func (c *PenyewaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	penyewa, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "penyewa not found"})
		return
	}
	ctx.JSON(http.StatusOK, penyewa)
}

func (c *PenyewaController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var penyewa models.Penyewa
	if err := ctx.ShouldBindJSON(&penyewa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	penyewa.ID = uint(id)
	if err := c.Repo.Update(&penyewa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penyewa)
}

func (c *PenyewaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := c.Repo.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetActiveTenants returns all active tenants
func (c *PenyewaController) GetActiveTenants(ctx *gin.Context) {
	penyewas, err := c.Repo.FindActive()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penyewas)
}

// GetTenantsWithDebt returns tenants with outstanding debts
func (c *PenyewaController) GetTenantsWithDebt(ctx *gin.Context) {
	penyewas, err := c.Repo.FindWithDebt()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penyewas)
}

// GetTenantWithBills returns a tenant along with their bills
func (c *PenyewaController) GetTenantWithBills(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	penyewa, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "penyewa not found"})
		return
	}

	tagihans, err := c.TagihanRepo.FindByPenyewaID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalDebt, err := c.TagihanRepo.GetTenantTotalDebt(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"penyewa":    penyewa,
		"tagihans":   tagihans,
		"total_debt": totalDebt,
	})
}

// Helper function to generate monthly bills for a tenant
func (c *PenyewaController) generateMonthlyBills(penyewa *models.Penyewa) error {
	// Get room information to determine the bill amount
	kamar, err := c.KamarRepo.FindByID(penyewa.KamarID)
	if err != nil {
		return err
	}

	// Parse start and end dates
	startDate, err := time.Parse("2006-01-02", penyewa.TanggalMulai)
	if err != nil {
		return err
	}

	endDate, err := time.Parse("2006-01-02", penyewa.TanggalSelesai)
	if err != nil {
		return err
	}

	// Generate bills for each month in the contract period
	currentDate := startDate
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Create bill for this month
		dueDate := time.Date(currentDate.Year(), currentDate.Month(), 15, 0, 0, 0, 0, time.Local)
		if dueDate.Before(time.Now()) {
			dueDate = time.Now().AddDate(0, 0, 5) // If already past the 15th, set due date to 5 days from now
		}

		tagihan := models.Tagihan{
			PenyewaID:         penyewa.ID,
			Bulan:             int(currentDate.Month()),
			Tahun:             currentDate.Year(),
			JumlahTagihan:     kamar.Harga,
			TanggalJatuhTempo: dueDate,
			Status:            models.StatusBelumLunas,
		}

		if err := c.TagihanRepo.Create(&tagihan); err != nil {
			return err
		}

		// Move to next month
		currentDate = currentDate.AddDate(0, 1, 0)
	}

	return nil
}
