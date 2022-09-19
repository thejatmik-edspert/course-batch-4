package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func WithLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		userID := c.Request.Context().Value("user_id")
		log.Println("=====", userID)
	}
}
