package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// EmailRX is a regular expression for determining if an email is valid
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Form type for setting form values and errors
type Form struct {
	url.Values
	Errors errors
}

// New creates a new Form from a set of url.Values
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required takes in required fields as strings and adds errors to the Form for any that are missing
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MatchesPattern accepts a regex to determine if the given field matches
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

// MinLength sets a minimum allowed length for a given field
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}
}

// MaxLength applies a max length validation to a given Form field
func (f *Form) MaxLength(field string, max int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > max {
		f.Errors.Add(field, fmt.Sprintf("This field cannot be longer than %d characters", max))
	}
}

// PermittedValues sets the allowable values for a given Form field
func (f *Form) PermittedValues(field string, allowed ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, allowedValue := range allowed {
		if value == allowedValue {
			return
		}
	}

	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if there are no validation errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
