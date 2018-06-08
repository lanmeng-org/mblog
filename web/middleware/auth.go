package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "%s", "test m")
		c.Abort()
	}
}