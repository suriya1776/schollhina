// db/db.go
package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(dialect, connectionStr string) {
	var err error
	DB, err = gorm.Open(dialect, connectionStr)
	if err != nil {
		fmt.Println("Failed to connect to the database. Error:", err)
		panic("Failed to connect to the database")
	}

	// Enable singular table names
	DB.SingularTable(true)

	// AutoMigrate your models here
	// e.g., DB.AutoMigrate(&YourModel{})
}
