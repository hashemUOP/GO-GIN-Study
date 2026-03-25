// main.go
package main

import (
	"login-app/handler"
	// "login-app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// // public route
	// r.POST("/login", handler.Login)

	// // protected route
	// r.GET("/profile", middleware.Auth(), handler.GetProfile)

	// new login route for testing
	r.POST("/login/test", handler.LoginTest)

	r.Run(":8081")
}
