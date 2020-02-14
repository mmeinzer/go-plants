package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is used when a record cannot be found in the DB
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials means the login credentials were not valid
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail means that a given email is already registered
	ErrDuplicateEmail = errors.New("models: duplicate email")
	// ErrNotActive means the user has the active flag set to false
	ErrNotActive = errors.New("models: the user is no longer active")
)

// Plant model
type Plant struct {
	ID    int
	Name  string
	Owner User
}

// User model
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
	Plants         []Plant
}
