// main.go
package main

import (
	"login-app/handler"
	"login-app/middleware"

	// "login-app/middleware"

	"login-app/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Database Connection
	database.ConnectDatabase()

	r := gin.Default()
	// // public route
	// r.POST("/login", handler.Login)

	// // protected route
	// r.GET("/profile", middleware.Auth(), handler.GetProfile)

	// new login route for testing
	r.POST("/login/test", handler.LoginTest)

	// new route to create user
	r.POST("/create", middleware.CreateUserMiddleware(), handler.CreateUser)

	//new route to delete user
	r.DELETE("/users/:username", middleware.DeleteUserMiddleware(), handler.DeleteUser)

	//new route to read user
	r.GET("/users/:username", middleware.ReadUserMiddleware(), handler.ReadUser)

	r.Run(":8081")
}
