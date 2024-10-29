package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}
