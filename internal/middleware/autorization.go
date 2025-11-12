package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OwnData() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		authId, exists := c.Get("id")
		if !exists || authId != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			c.Abort()
			return
		}
	}
}
