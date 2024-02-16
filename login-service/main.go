// main.go
package main

import (
	"schoolhina/login-service/db"
	"schoolhina/login-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the database
	db.InitDB("root", "Admin@123", "schoolhina")
	defer db.CloseDB()

	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", routes.Register)
		authGroup.POST("/login", routes.Login)
	}

	router.Run(":8080")
}
