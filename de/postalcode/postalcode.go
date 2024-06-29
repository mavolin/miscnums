// Package postalcode provides parsing and validation for German postal codes
// (Postleitzahl).
package postalcode

import (
	"encoding"
	"errors"
	"regexp"
)

var ErrSyntax = errors.New("postal code: must be 5-digit number")

// PostalCode represents a German postal code code (Postleitzahl).
//
// It is a 5-digit string.
type PostalCode string

func (c PostalCode) String() string {
	return string(c)
}

func (c PostalCode) Compact() string {
	return string(c)
}

var _ encoding.TextMarshaler = PostalCode("")

func (c PostalCode) MarshalText() ([]byte, error) {
	return []byte(c.Compact()), nil
}

var _ encoding.TextUnmarshaler = (*PostalCode)(nil)

func (c *PostalCode) UnmarshalText(text []byte) error {
	parsed, err := Parse(string(text))
	if err != nil {
		return err
	}

	*c = parsed
	return nil
}

var postalCodeRegexp = regexp.MustCompile(`\d{5}`)

// Parse parses the passed postal code.
//
// If Parse returns without an error, the postal code is considered
// syntactically valid.
//
// Note that while checking if the postal code actually exists is possible,
// it is also error-prone, because as soon as a new postal code is assigned,
// or one is removed, the validation automatically becomes incorrect.
// Therefore, this package only checks for syntactical validity.
func Parse(s string) (PostalCode, error) {
	if !postalCodeRegexp.MatchString(s) {
		return "", ErrSyntax
	}

	return PostalCode(s), nil
}

// IsValid returns true if the postal code is valid.
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
