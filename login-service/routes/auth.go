// routes/auth.go
package routes

import (
	"net/http"

	"schoolhina/login-service/db"
	"schoolhina/login-service/models"
	"schoolhina/login-service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Assuming your models.User has a 'Role' field
func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user has the required role to register
	if !isSpecialRole(newUser.Role) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role for registration"})
		return
	}

	// For regular users, hash the password and save to 'users' table
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if isSpecialRole(newUser.Role) {
		// Register unapproved user and redirect to a page for additional school information
		unapprovedUserID, err := db.InsertUnapprovedUser(newUser.Username, string(hashedPassword), newUser.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusSeeOther, gin.H{"message": "Redirect to school information page", "unapprovedUserID": unapprovedUserID})
		return
	}

	userID, err := db.InsertUser(newUser.Username, string(hashedPassword), newUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Generate and return an authentication token
	authToken, err := utils.GenerateAuthToken(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate auth token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered", "token": authToken})
}

// isValidRegistrationRole checks if the role is valid for registration
// isSpecialRole checks if the role is chairperson or CEO
func isSpecialRole(role string) bool {
	specialRoles := []string{"chairperson", "CEO"}
	for _, specialRole := range specialRoles {
		if role == specialRole {
			return true
		}
	}
	return false
}

func Login(c *gin.Context) {
	// Parse JSON request body
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user from the database by username (use your database logic here)
	storedUser, err := db.GetUserByUsername(loginUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user from the database"})
		return
	}

	if storedUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate and return an authentication token (JWT) if login is successful
	authToken, err := utils.GenerateAuthToken(storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating authentication token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": authToken})
}
