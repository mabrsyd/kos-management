package controllers

import (
	"net/http"
	"strconv"

	"kos-management/models"
	"kos-management/repositories"

	"github.com/gin-gonic/gin"
)

type TransaksiController struct {
	Repo *repositories.TransaksiRepository
}

func (c *TransaksiController) Create(ctx *gin.Context) {
	var transaksi models.Transaksi
	if err := ctx.ShouldBindJSON(&transaksi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Repo.Create(&transaksi); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, transaksi)
}

func (c *TransaksiController) GetAll(ctx *gin.Context) {
	transaksis, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaksis)
}

func (c *TransaksiController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	transaksi, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "transaksi not found"})
		return
	}
	ctx.JSON(http.StatusOK, transaksi)
}

func (c *TransaksiController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var transaksi models.Transaksi
	if err := ctx.ShouldBindJSON(&transaksi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaksi.ID = uint(id)
	if err := c.Repo.Update(&transaksi); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaksi)
}

func (c *TransaksiController) Delete(ctx *gin.Context) {
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
