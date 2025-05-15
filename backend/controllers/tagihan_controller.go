package controllers

import (
	"net/http"
	"strconv"
	"time"

	"kos-management/models"
	"kos-management/repositories"

	"github.com/gin-gonic/gin"
)

type TagihanController struct {
	Repo *repositories.TagihanRepository
}

func (c *TagihanController) Create(ctx *gin.Context) {
	var tagihan models.Tagihan
	if err := ctx.ShouldBindJSON(&tagihan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Repo.Create(&tagihan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, tagihan)
}

func (c *TagihanController) GetAll(ctx *gin.Context) {
	tagihans, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tagihans)
}

func (c *TagihanController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	tagihan, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "tagihan not found"})
		return
	}
	ctx.JSON(http.StatusOK, tagihan)
}

func (c *TagihanController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var tagihan models.Tagihan
	if err := ctx.ShouldBindJSON(&tagihan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagihan.ID = uint(id)
	if err := c.Repo.Update(&tagihan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tagihan)
}

func (c *TagihanController) Delete(ctx *gin.Context) {
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

// GetByPenyewaID returns all bills for a tenant
func (c *TagihanController) GetByPenyewaID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	tagihans, err := c.Repo.FindByPenyewaID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tagihans)
}

// GetUnpaidByPenyewaID returns all unpaid bills for a tenant
func (c *TagihanController) GetUnpaidByPenyewaID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	tagihans, err := c.Repo.FindUnpaidByPenyewaID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tagihans)
}

// GetBillsByMonth returns all bills for a specific month and year
func (c *TagihanController) GetBillsByMonth(ctx *gin.Context) {
	bulan, err := strconv.Atoi(ctx.Query("bulan"))
	if err != nil || bulan < 1 || bulan > 12 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	tahun, err := strconv.Atoi(ctx.Query("tahun"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}

	tagihans, err := c.Repo.FindBillsByMonth(bulan, tahun)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tagihans)
}

// MarkAsPaid marks a bill as paid
func (c *TagihanController) MarkAsPaid(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	tagihan, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "bill not found"})
		return
	}

	tagihan.Status = models.StatusLunas
	tagihan.UpdatedAt = time.Now()

	if err := c.Repo.Update(tagihan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tagihan)
}

// UpdateOverdueBills updates the status of bills that are past due date
func (c *TagihanController) UpdateOverdueBills(ctx *gin.Context) {
	if err := c.Repo.UpdateOverdueBills(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "overdue bills updated"})
}

// GetTenantDebt returns the total debt for a tenant
func (c *TagihanController) GetTenantDebt(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	totalDebt, err := c.Repo.GetTenantTotalDebt(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"total_debt": totalDebt})
}
