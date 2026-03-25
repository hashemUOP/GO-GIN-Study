// handler/user_handler.go
package handler

import (
	"net/http"

	"login-app/model"

	"github.com/gin-gonic/gin"
)

// fake user (instead of DB)
var user = model.User{
	ID:       1,
	Username: "admin",
	Password: "1234",
}

// 🔐 LOGIN
func Login(c *gin.Context) {
	var input model.User

	// read JSON from request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// check credentials
	if input.Username == user.Username && input.Password == user.Password {
		// return fake token
		c.JSON(http.StatusOK, gin.H{
			"token": "valid-token",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// 👤 PROFILE (protected)
func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}
