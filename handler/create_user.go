package handler

import (
	"login-app/database"
	"login-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.CreateUser

	// 1. Bind JSON to Struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// 2. Insert into PostgreSQL
	result := database.DB.Create(&user)

	if result.Error != nil {
		// This will now tell you EXACTLY what PostgreSQL is complaining about
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"details": result.Error.Error(),
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// 3. Success Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
