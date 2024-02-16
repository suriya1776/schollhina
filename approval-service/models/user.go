// models/user.go
package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Role     string
	Approved bool // New field to track approval status
}
