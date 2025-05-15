package controllers

import (
	"net/http"
	"strconv"

	"kos-management/models"
	"kos-management/repositories"
	"kos-management/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Repo.FindByUsername(loginRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"role":     user.Role,
		},
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := c.Repo.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Don't return the password
	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := c.Repo.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Don't return the password
	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Don't return the password
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Get existing user first
	existingUser, err := c.Repo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Parse request body
	var updateRequest struct {
		Username    string `json:"username"`
		Name        string `json:"name"`
		Password    string `json:"password"`    // Current password for verification
		NewPassword string `json:"newPassword"` // New password if changing
		Role        string `json:"role"`
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If current password is provided, verify it
	if updateRequest.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(updateRequest.Password))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password saat ini tidak valid"})
			return
		}
	} else {
		// For profile updates, password verification is required
		currentUserID, exists := ctx.Get("userID")
		if exists && currentUserID.(uint) == existingUser.ID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password diperlukan untuk memperbarui profil"})
			return
		}
	}

	// Update fields if they are provided and not empty
	if updateRequest.Name != "" {
		existingUser.Name = updateRequest.Name
	}

	if updateRequest.Role != "" {
		// Only let admins change roles
		currentUserID, exists := ctx.Get("userID")
		if exists {
			currentUser, err := c.Repo.FindByID(currentUserID.(uint))
			if err == nil && currentUser.Role == "admin" {
				existingUser.Role = updateRequest.Role
			}
		}
	}

	// Update password if a new one is provided
	if updateRequest.NewPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateRequest.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	if err := c.Repo.Update(existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Don't return the password
	existingUser.Password = ""
	ctx.JSON(http.StatusOK, existingUser)
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	// Get user ID from the JWT token
	userID, _ := ctx.Get("userID")

	// Get existing user from database
	existingUser, err := c.Repo.FindByID(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Parse request body
	var updateRequest struct {
		Name        string `json:"name"`
		Password    string `json:"password"`    // Current password
		NewPassword string `json:"newPassword"` // New password if changing
	}

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify current password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(updateRequest.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password saat ini tidak valid"})
		return
	}

	// Update name if provided
	if updateRequest.Name != "" {
		existingUser.Name = updateRequest.Name
	}

	// Update password if a new one is provided
	if updateRequest.NewPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateRequest.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui password"})
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	// Save the changes
	if err := c.Repo.Update(existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Don't return the password
	existingUser.Password = ""
	ctx.JSON(http.StatusOK, existingUser)
}

func (c *UserController) Delete(ctx *gin.Context) {
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
