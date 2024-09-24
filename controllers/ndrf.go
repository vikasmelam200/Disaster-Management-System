package controllers

import (
	"DisasterManagement/database"
	"DisasterManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHighWayStatus(c *gin.Context) {
	var disasterInfo []models.DisasterInfo
	err := database.DB.Find(&disasterInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retreive high way status"})
		return
	}
	c.JSON(http.StatusOK, disasterInfo)
}
func AddDisasterInfo(c *gin.Context) {
	var disaster models.DisasterInfo
	err := c.ShouldBindJSON(&disaster)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data:" + err.Error()})
		return
	}
	err = database.DB.Create(&disaster).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add disaster info" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, disaster)
}
func GetNDRFRouteInfo(c *gin.Context) {
	var disasterInfo models.DisasterInfo
	affectedArea := c.Query("affected_area")
	if affectedArea == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "affected_area query parameter is required"})
		return
	}
	err := database.DB.Where("affected_area=?", affectedArea).First(&disasterInfo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "information not available for the affected area"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"road_status": disasterInfo,
		"best_routes":  "Best routes for NDRF to reach affected areas",
		"contact_atms": "Audio communication channel with ATMS control centre"})
}
