package handler

import (
	// "login-app/database"
	"login-app/database"
	"login-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// 1. Get the username from the URL parameter
	username := c.Param("username")

	// 2. Execute the Delete query
	// Notice we pass an empty &models.User{} to tell GORM which table to use
	result := database.DB.Where("username = ?", username).Delete(&model.User{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": result.Error.Error()})
		return
	}

	// 3. Check if the user actually existed
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 4. Success Response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
