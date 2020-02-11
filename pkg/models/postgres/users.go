package postgres

import (
	"database/sql"

	"mattmeinzer.com/plants/pkg/models"
)

// UserModel wraps a sql connection pool for querying
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate verifies a user exists and credentials match
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get method to fetch details for a specific user
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
