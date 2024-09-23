package controllers

import (
	"DisasterManagement/database"
	"DisasterManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReportDisaster creates a new disaster record in the database

func ReportDisaster(c *gin.Context) {
	// Check if the database connection is initialized
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var disaster models.Disaster
	// Bind the request JSON to the disaster model
	err := c.ShouldBindJSON(&disaster)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Create the disaster record in the database
	err = database.DB.Create(&disaster).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to report disaster: " + err.Error()})
		return
	}

	// Return the newly created disaster record
	c.JSON(http.StatusOK, disaster)
}

// GetDisasters retrieves all disasters from the database
func GetDisasters(c *gin.Context) {
	// Check if the database connection is initialized
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var disasters []models.Disaster
	// Retrieve all disasters from the database
	err := database.DB.Find(&disasters).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve disasters: " + err.Error()})
		return
	}

	// Return the list of disasters
	c.JSON(http.StatusOK, disasters)
}
