package middleware

import (
	"fmt"

	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// RoleGuard restricts endpoint access to specific roles from JWT claims
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			response.Unauthorized(c, "Otorisasi peran gagal: Role tidak ditemukan di token")
			c.Abort()
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
			response.Forbidden(c, fmt.Sprintf("Akses terlarang. Diperlukan salah satu dari peran: %v, Anda masuk sebagai: %s", allowedRoles, userRole))
			c.Abort()
			return
		}

		c.Next()
	}
}
