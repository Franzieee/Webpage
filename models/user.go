package models

import (
	"time"
	"webpage/db" // Access the shared DB connection

	"golang.org/x/crypto/bcrypt" //hashing package
)

// User represents a row in the "users" table
type User struct {
	ID           int       `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	CreateAt    time.Time `db:"created_at"`
}

// Login form that gets the username before login
func GetUserByUsername(username string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, username, password_hash, role, created_at
		FROM users
		WHERE username = $1
	`

	var u User

	// Scan each selected field into the corresponding struct field
	err := db.DB.QueryRow(query, username).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.PasswordHash,
		&u.Role,
		&u.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Register area
// Registers a new user information in the database
func CreateUser(FirstName, LastName, Username, Password, Role string) error {
	// Hashing the password to keep it safe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// For Query
	query := `
		INSERT INTO users (first_name, last_name, username, password_hash, role)
		VALUES ($1, $2, $3, $4, $5)
	`

	// Execution of query
	_, err = db.DB.Exec(query, FirstName, LastName, Username, string(hashedPassword), Role)
	if err != nil {
		return err
	}

	return nil

}
