package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthGuard validasi token JWT Bearer
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Akses ditolak: Token tidak ditemukan atau format salah",
			})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "super_secret_siul_pbj_key_2026"
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Token kadaluarsa atau tidak valid",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Simpan data _claim_ ke Gin Context
			c.Set("user_id", claims["id"])
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Gagal membaca claims JWT",
			})
		}
	}
}

// RoleGuard membatasi akses endpoint untuk role spesifik
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Otorisasi peran gagal: Role tidak ditemukan di token",
			})
			return
		}

		userRole := roleValue.(string)

		isAllowed := false
		for _, allowed := range allowedRoles {
			if allowed == userRole {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": fmt.Sprintf("Akses terlarang. Diperlukan salah satu dari peran: %v, Anda masuk sebagai: %s", allowedRoles, userRole),
			})
			return
		}

		c.Next()
	}
}
