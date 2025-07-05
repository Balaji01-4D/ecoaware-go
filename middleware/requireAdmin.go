package middleware

import (
	"net/http"

	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)

func requireAdmin(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	if user.Role != models.RoleAdmin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}

func RequireAdmin() gin.HandlerFunc {
	return requireAdmin
}