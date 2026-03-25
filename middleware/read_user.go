package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func ReadUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for read user middleware
		fmt.Println("Reading user middleware")

		c.Next() // call the next handler
	}
}