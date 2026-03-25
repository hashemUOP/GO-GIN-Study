package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("create user middleware called")
		c.Next() // call the next handler
	}
}
