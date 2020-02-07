package models

import (
	"errors"
)

// ErrNoRecord should be used when a record can't be found
var ErrNoRecord = errors.New("models: no matching record found")

// Plant model
type Plant struct {
	ID   int
	Name string
}
