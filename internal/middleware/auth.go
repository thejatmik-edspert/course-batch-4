package middleware

import (
	"context"
	"course/internal/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	bearer = "Bearer "
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, map[string]interface{}{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, bearer) {
			c.JSON(401, map[string]interface{}{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		auth := strings.Split(authHeader, " ")
		// todo melakukan validasi/decyprt jwt tokennya
		user := domain.User{}
		data, err := user.DecryptJwt(auth[1])
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		userID := int(data["user_id"].(float64))
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userID)
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
