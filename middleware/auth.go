package middleware

import (
	database "booking-konstruksi/database/migration"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		// Check the prefix header as a Bearer Token
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"messages": "Authorization token not provided",
			})

			c.Abort()
			return
		}

		// Check token in table auth
		var auth database.AuthUser

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		err := db.Where("token = ? AND expires_at > ?", tokenStr, time.Now()).First(&auth).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"messages": "Unauthorization",
			})

			c.Abort()
			return
		}

		c.Set("userID", auth.UserID)
		c.Next()
	}
}
