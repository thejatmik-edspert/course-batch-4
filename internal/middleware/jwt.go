package middleware

import (
	"context"
	"strings"

	"course/internal/user"

	"github.com/gin-gonic/gin"
)

const bearer = "Bearer"

func UseClaims(us *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, bearer) {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
		}
		token := strings.Split(auth, " ")[1]
		if token == "" {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
		}

		claims, err := us.DecriptJWT(token)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", claims["user_id"])
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
