package routes

import (
	"DisasterManagement/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/disaster/report", controllers.ReportDisaster)
	api.GET("/disasters", controllers.GetDisasters)
	api.GET("/phases", controllers.AddPhase)
	api.POST("/phases", controllers.GetPhases)
}
