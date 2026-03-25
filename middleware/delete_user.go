package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DeleteUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for delete user middleware
		fmt.Println("delete user middleware called")
		c.Next() // call the next handler
	}
}
