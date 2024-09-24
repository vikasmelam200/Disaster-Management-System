package controllers

import (
	"DisasterManagement/database"
	"DisasterManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPhases retrieves all phases from the database and returns them as JSON
func GetPhases(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var phases []models.Phase
	err := database.DB.Find(&phases).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve phases: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, phases)
}

// AddPhase adds a new phase to the database and returns the created phase as JSON
func AddPhase(c *gin.Context) {
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	var phase models.Phase
	err := c.ShouldBindJSON(&phase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Insert the phase into the database
	err = database.DB.Create(&phase).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add phase: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, phase)
}
