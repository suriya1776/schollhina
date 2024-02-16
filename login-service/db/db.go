// db/db.go

package db

import (
	"database/sql"
	"fmt"
	"log"

	"schoolhina/login-service/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(username, password, dbName string) {
	// Create a database connection
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbName)
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Closed the database connection")
	}
}

func InsertUser(username, hashedPassword, role string) (int64, error) {
	result, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", username, hashedPassword, role)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	// Create a query to select a user by username
	query := "SELECT id, username, password FROM users WHERE username = ?"

	// Execute the query
	row := db.QueryRow(query, username)

	// Create a user object to store the result
	var user models.User

	// Scan the result into the user object
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		// No user found with the given username
		return nil, nil
	} else if err != nil {
		// An error occurred during the query
		return nil, err
	}

	// User found, return the user object
	return &user, nil
}

func InsertUnapprovedUser(username, hashedPassword, role string) (int64, error) {
	// Insert user details into 'users' table without providing a password
	result, err := db.Exec("INSERT INTO users (username,password, role) VALUES (?,?,?)", username, hashedPassword, role)
	if err != nil {
		fmt.Println("Error inserting into users table:", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return 0, err
	}

	// Insert unapproved user details into 'unapproved_users' table
	_, err = db.Exec("INSERT INTO unapproved_users (user_id) VALUES (?)", lastInsertID)
	if err != nil {
		fmt.Println("Error inserting into unapproved_users table:", err)
		return 0, err
	}

	return lastInsertID, nil
}
