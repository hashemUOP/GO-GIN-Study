package handler

import (
	"login-app/database"
	"login-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadUser(c *gin.Context) {
	// 1. Extract the username from the URL parameter
	username := c.Param("username")
	var user model.User

	// 2. Query the database
	// This translates to: SELECT * FROM users WHERE username = 'sara_dev';
	result := database.DB.Where("username = ?", username).Take(&user)

	// 3. Handle Errors (e.g., User doesn't exist)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": result.Error.Error()})
		return
	}

	// 4. Return the user data
	c.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"user":    user,
	})
}
