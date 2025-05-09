package controllers

import (
	"net/http"
	"strconv"

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
