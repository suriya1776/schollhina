// utils/auth_utils.go

package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key") // Replace with your actual secret key

// ValidateMasterUserCredentials validates the master user credentials
func ValidateMasterUserCredentials(username, password string) bool {
	// Implement master user credential validation logic here
	// For simplicity, assume all credentials are valid in this example
	return true
}

// IsValidMasterToken checks if the provided token is a valid master token
func IsValidMasterToken(tokenString string) bool {
	// Implement master token validation logic here
	// For simplicity, assume all tokens are valid in this example
	return true
}

func GenerateAuthToken(userID uint) (string, error) {
	// Set the token claims
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token with RSA 256 encryption
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
