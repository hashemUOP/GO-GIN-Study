package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct that matches the JSON body
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginTest(c *gin.Context) {
	var user User

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Now user.Name and user.Age contain the data
	c.JSON(http.StatusOK, gin.H{
		"message": "User received",
		"user":    user,
	})
	fmt.Printf("Received user: %+v\n", user)
}
