package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"siul-pbj-api/config"
	"siul-pbj-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi gagal, pastikan username dan password terisi",
		})
		return
	}

	var user models.User
	// Fetch user from DB
	result := config.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Username atau password salah",
		})
		return
	}

	// NOTE: In production you MUST use bcrypt to compare password!
	// For testing with raw SQL DBeaver directly we just use plain text or dummy hash match.
	if input.Password != user.PasswordHash { // TODO: Replace with bcrypt.CompareHashAndPassword
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Username atau password salah",
		})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 1 day expiry
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super_secret_siul_pbj_key_2026"
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("JWT Generate Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal men-generate token auth",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login berhasil",
		"data": gin.H{
			"token": tokenString,
			"user": gin.H{
				"id":       user.ID,
				"nama":     user.Nama,
				"username": user.Username,
				"role":     user.Role,
			},
		},
	})
}
