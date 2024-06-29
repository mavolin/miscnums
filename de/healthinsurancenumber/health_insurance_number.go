// Package healthinsurancenumber provides parsing and validation for German
// health insurance numbers.
package healthinsurancenumber

import (
	"encoding"
	"errors"
	"regexp"
)

var (
	ErrSyntax     = errors.New("health insurance number: invalid health insurance number")
	ErrCheckDigit = errors.New("health insurance number: invalid check digit")
)

// HealthInsuranceNumber represents a 10-digit German health insurance id.
type HealthInsuranceNumber string

var healthInsuranceIDRegexp = regexp.MustCompile(`^([A-Z])([0-9]{8})([0-9])$`)

func Parse(s string) (HealthInsuranceNumber, error) {
	// https://de.wikipedia.org/wiki/Krankenversichertennummer#Deutschland_seit_2012
	m := healthInsuranceIDRegexp.FindStringSubmatch(s)
	if len(m) == 0 {
		return "", ErrSyntax
	}

	letterToNum := m[1][0] - 'A' + 1

	var sum int

	sum += int(letterToNum / 10)

	digit2 := letterToNum % 10 * 2
	if digit2 > 9 {
		digit2 -= 9
	}
	sum += int(digit2)

	numStr := m[2]
	for i, r := range numStr {
		n := int(r - '0')
		if i%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	checkDigit := int(m[3][0] - '0')
	if checkDigit != sum%10 {
		return "", ErrCheckDigit
	}

	return HealthInsuranceNumber(s), nil
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}

func (id HealthInsuranceNumber) String() string {
	return string(id)
}

func (id HealthInsuranceNumber) Compact() string {
	return string(id)
}

var _ encoding.TextMarshaler = HealthInsuranceNumber("")

func (id HealthInsuranceNumber) MarshalText() ([]byte, error) {
	return []byte(id.Compact()), nil
}

var _ encoding.TextUnmarshaler = (*HealthInsuranceNumber)(nil)

func (id *HealthInsuranceNumber) UnmarshalText(text []byte) error {
	*id = HealthInsuranceNumber(text)
	return nil
}
