package models

import (
	"time"
	"webpage/db" // Access the shared DB connection
)

// User represents a row in the "users" table
type User struct {
	ID           int       `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Userame      string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	TimeStamp    time.Time `db:"time_stamp"`
}

func GetUserByUsername(username string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, username, password_hash, role, time_stamp
		FROM users
		WHERE username = $1
	`

	var u User

	// Scan each selected field into the corresponding struct field
	err := db.DB.QueryRow(query, username).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Userame,
		&u.PasswordHash,
		&u.Role,
		&u.TimeStamp,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
