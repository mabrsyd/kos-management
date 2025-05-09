package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"kos-management/models"
	"kos-management/repositories"

	"github.com/gin-gonic/gin"
)

type KamarController struct {
	Repo *repositories.KamarRepository
}

func (c *KamarController) Create(ctx *gin.Context) {
	var kamar models.Kamar
	if err := ctx.ShouldBindJSON(&kamar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Format data tidak valid: %s", err.Error())})
		return
	}

	// Validasi data
	if kamar.Nomer == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nomor kamar tidak boleh kosong"})
		return
	}

	if kamar.Harga <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Harga kamar harus lebih dari 0"})
		return
	}

	// Jika user_id tidak diset, coba ambil dari context
	if kamar.UserID == 0 {
		userID, exists := ctx.Get("userID")
		if exists {
			kamar.UserID = userID.(uint)
		}
	}

	// Log untuk debugging
	log.Printf("Creating kamar: %+v", kamar)

	if err := c.Repo.Create(&kamar); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal menyimpan ke database: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusCreated, kamar)
}

func (c *KamarController) GetAll(ctx *gin.Context) {
	kamars, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal mengambil data kamar: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, kamars)
}

func (c *KamarController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID kamar tidak valid"})
		return
	}
	kamar, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kamar tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, kamar)
}

func (c *KamarController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID kamar tidak valid"})
		return
	}

	// Dapatkan kamar yang ada
	existingKamar, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kamar tidak ditemukan"})
		return
	}

	var updatedKamar models.Kamar
	if err := ctx.ShouldBindJSON(&updatedKamar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Format data tidak valid: %s", err.Error())})
		return
	}

	// Validasi data
	if updatedKamar.Nomer == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nomor kamar tidak boleh kosong"})
		return
	}

	if updatedKamar.Harga <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Harga kamar harus lebih dari 0"})
		return
	}

	// Set ID dari URL path
	updatedKamar.ID = uint(id)

	// Jika user_id kosong, gunakan user_id dari kamar yang sudah ada
	if updatedKamar.UserID == 0 {
		updatedKamar.UserID = existingKamar.UserID
	}

	// Log untuk debugging
	log.Printf("Updating kamar: %+v", updatedKamar)

	if err := c.Repo.Update(&updatedKamar); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal memperbaharui data: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, updatedKamar)
}

func (c *KamarController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID kamar tidak valid"})
		return
	}

	// Periksa apakah kamar ada
	_, err = c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kamar tidak ditemukan"})
		return
	}

	if err := c.Repo.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal menghapus kamar: %s", err.Error())})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Kamar berhasil dihapus"})
}
