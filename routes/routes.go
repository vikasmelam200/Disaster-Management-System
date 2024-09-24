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
	api.GET("/atms/services", controllers.GetATMSServices)
	api.POST("/atms/service", controllers.AddATMSService)
	api.GET("/highway/status", controllers.GetHighWayStatus)
	api.POST("/disaster/info", controllers.AddDisasterInfo)
	api.GET("/ndrf/route-info", controllers.GetNDRFRouteInfo)

}
