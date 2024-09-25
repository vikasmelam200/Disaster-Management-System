package routes

import (
	"DisasterManagement/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/atms/disaster/report", controllers.ReportDisaster)
	api.GET("/atms/disasters", controllers.GetDisasters)
	api.GET("/atms/phases", controllers.AddPhase)
	api.POST("/atms/phases", controllers.GetPhases)
	api.GET("/atms/services", controllers.GetATMSServices)
	api.POST("/atms/service", controllers.AddATMSService)
	api.GET("/atms/highway/status", controllers.GetHighWayStatus)
	api.POST("/atms/disaster/info", controllers.AddDisasterInfo)
	api.GET("/atms/ndrf/route-info", controllers.GetNDRFRouteInfo)
	api.POST("/atms/enhancement", controllers.AddATMSEnhancement)
	api.GET("/atms/enhancements", controllers.GetATMSEnhancements)
}
