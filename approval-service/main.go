// main.go

package main

import (
	"schoolhina/approval-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize and setup routes for the approval-service

	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", routes.Login)
	}

	router.Run(":8081")
}
