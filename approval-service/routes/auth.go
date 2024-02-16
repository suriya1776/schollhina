// routes/auth.go

package routes

import (
	"net/http"

	"schoolhina/approval-service/utils"

	"github.com/gin-gonic/gin"
)

// Login is the login endpoint for the master user
func Login(c *gin.Context) {

	// var loginUser models.User
	// Validate master credentials
	if !utils.ValidateMasterUserCredentials(c.PostForm("username"), c.PostForm("password")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Assuming utils.GenerateAuthToken is a function that generates your authentication token
	authToken, err := utils.GenerateAuthToken(0) // 0 is a placeholder for the master user ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate auth token"})
		return
	}

	// Send a custom response with the authentication token
	c.JSON(http.StatusOK, gin.H{"message": "Master user logged in", "token": authToken})
}

// ApproveUser approves a user in the approval-service
// func ApproveUser(c *gin.Context) {
// 	// Validate master token
// 	masterToken := c.GetHeader("Authorization")
// 	if !utils.IsValidMasterToken(masterToken) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid master token"})
// 		return
// 	}

// 	// Extract user ID from the request
// 	userID, err := utils.GetUserIDFromRequest(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user ID"})
// 		return
// 	}

// 	// Implement your approval logic here

// 	c.JSON(http.StatusOK, gin.H{"message": "User approved"})
// }
