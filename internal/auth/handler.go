package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"siul-pbj-api/internal/domain"
	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Handler holds dependencies for auth endpoints
type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Validasi gagal, pastikan username dan password terisi")
		return
	}

	var user domain.User
	result := h.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		response.Unauthorized(c, "Username atau password salah")
		return
	}

	// TODO: Replace with bcrypt.CompareHashAndPassword in production
	if input.Password != user.PasswordHash {
		response.Unauthorized(c, "Username atau password salah")
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super_secret_siul_pbj_key_2026"
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("JWT Generate Error:", err)
		response.InternalError(c, "Gagal men-generate token auth")
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
