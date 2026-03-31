package middleware

import (
	"fmt"
	"os"
	"strings"

	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthGuard validates JWT Bearer tokens and sets user claims in context
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(c, "Akses ditolak: Token tidak ditemukan atau format salah")
			c.Abort()
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
			response.Unauthorized(c, "Token kadaluarsa atau tidak valid")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["id"])
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			response.Unauthorized(c, "Gagal membaca claims JWT")
			c.Abort()
		}
	}
}
