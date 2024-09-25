package controllers

import (
	"DisasterManagement/database"
	"DisasterManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetATMSEnhancements(c *gin.Context) {
	var enhancements []models.AtmsEnhancement
	err := database.DB.Find(&enhancements).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retreive enhancement"})
		return
	}
	c.JSON(http.StatusOK, enhancements)
}
func AddATMSEnhancement(c *gin.Context) {
	var enhancement models.AtmsEnhancement
	err := c.ShouldBindJSON(&enhancement)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&enhancement).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add enhancement details"})
		return
	}
	c.JSON(http.StatusCreated, enhancement)
}
