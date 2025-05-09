package controllers

import (
	"net/http"
	"strconv"

	"kos-management/models"
	"kos-management/repositories"

	"github.com/gin-gonic/gin"
)

type PenyewaController struct {
	Repo *repositories.PenyewaRepository
}

func (c *PenyewaController) Create(ctx *gin.Context) {
	var penyewa models.Penyewa
	if err := ctx.ShouldBindJSON(&penyewa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Repo.Create(&penyewa); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
