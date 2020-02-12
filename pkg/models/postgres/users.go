package postgres

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"mattmeinzer.com/plants/pkg/models"
)

// UserModel wraps a sql connection pool for querying
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := "INSERT INTO users(name, email, hashed_password, created) VALUES($1, $2, $3, now())"

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code.Name() == "unique_violation" {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

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
