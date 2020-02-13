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
	stmt := "SELECT id, hashed_password, active FROM users WHERE email = $1"

	var id int
	var hashedPassword []byte
	var active bool
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword, &active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	if !active {
		return 0, models.ErrNotActive
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

// Get method to fetch details for a specific user
func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}

	stmt := "SELECT id, name, email, created, active FROM users WHERE id = $1"
	err := m.DB.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created, &u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return u, nil
}
