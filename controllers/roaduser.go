package controllers

import (
	"DisasterManagement/database"
	"DisasterManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetATMSServices(c *gin.Context) {
	var services []models.AtmsService
	//just checking database connection is available or not
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not intialized"})
	}
	err := database.DB.Find(&services).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}
	c.JSON(http.StatusOK, services)
}
func AddATMSService(c *gin.Context) {
	var service models.AtmsService
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}
	err := c.ShouldBindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data " + err.Error()})
		return
	}
	err = database.DB.Create(&service).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Failed to create atms service": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}
