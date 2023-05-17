package middlewares

import (
	"campus_pharmacy_share/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}

		user, err := models.GetUserByID(int(claims.(*models.Claims).ID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}

		if user.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}

		c.Next()
	}
}
