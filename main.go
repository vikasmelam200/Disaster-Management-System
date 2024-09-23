package main

import (
	"DisasterManagement/database"
	"DisasterManagement/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDB()
	routes.InitializeRouter(r)
	r.Run(":8080")
}
