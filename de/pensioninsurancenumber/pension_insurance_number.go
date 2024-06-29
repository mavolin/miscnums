// Package pensioninsurancenumber provides parsing and validation for German
// Pension Insurance Numbers (Rentenversicherungsnummer).
package pensioninsurancenumber

import (
	"encoding"
	"fmt"
	"time"
)

type PensionInsuranceNumber struct {
	// AreaCode is the area code of the pension insurance number.
	AreaCode AreaCode
	// BirthDay is the birthday of the holder of the pension insurance number.
	//
	// It may not necessarily reflect the actual birthday of the holder, if
	// there are more than 50 people with the same date of birth in the same
	// area with the same first letter of the last name, and the same serial
	// number group.
	// Should that be the case, then BirthDay is the BirthDay will be a number
	// greater than 31.
	BirthDay uint8
	// BirthMonth is the birth month of the holder of the pension insurance
	// number.
	BirthMonth uint8
	// BirthYear are the last two digits of the birth year of the holder of the
	// pension insurance number.
	BirthYear uint8
	// LastNameLetter is the first letter of the last name of the holder of the
	// pension insurance number.
	LastNameLetter rune
	// SerialNumber is the serial number of the pension insurance number.
	//
	// A number between 0 and 49 indicates the holder identified as male when
	// applying for the pension insurance number.
	// A number between 50 and 99 indicates the holder identified as female
	// or non-binary ('divers') when applying for the pension insurance number.
	SerialNumber uint8
	CheckDigit   uint8
}

// Birthdate returns the assumed birthday of the holder of the pension
// insurance number.
//
// The returned time uses UTC as location.
//
// # Inaccuracies
//
// Birthdate may fail to return an accurate birthday, in any of these cases:
//
//  1. The pension insurance number is invalid.
//     In this case, Birthdate returns the zero value of time.Time.
//     You can use [time.Time.IsZero] to check if that is the case.
//
//     Note that Birthdate will not perform a full validity check.
//     Currently, the only case where Birthdate will return a zero time, is
//     when BirthDay is zero, or BirthMonth is outside the valid 1-12 range.
//
//  2. The pension insurance number is valid, but the birthday cannot be
//     determined, because of the special case described in
//     doc of the BirthDay field.
//     In this case, Birthdate returns a valid time that uses the first of the
//     month as day.
//     Check if BirthDay > 31 to determine if this is the case.
//
//  3. The pension insurance number holder is 100 years or older:
//     Since the BirthYear is specified using only two digits, it is not
//     possible to always determine the correct century.
//     Instead, Birthdate assumes that the holder is younger than 100 years,
//     and picks a year within the last 99 years.
func (pin PensionInsuranceNumber) Birthdate() time.Time {
	if pin.BirthMonth < 1 || pin.BirthMonth > 12 {
		return time.Time{}
	} else if pin.BirthDay == 0 {
		return time.Time{}
	}

	// Determine the Correct Year to Use

	now := time.Now().In(time.UTC)

	now4Year := now.Year()     // four-digit year
	now2Year := now4Year % 100 // two-digit year

	century := now4Year - now2Year // current century, e.g 2000

	// Assume we're dealing with a person younger than 100 years, then this
	// person was either born in this or the previous century.
	// If the last two digits of their birth year are greater than the last
	// two digits of the current year, then they must be born in the previous
	// century, as otherwise they would be born in the future.
	// In any other case, they must be born in the current century.
	if pin.BirthYear > uint8(now2Year) {
		century -= 100
	}

	birthYear := century + int(pin.BirthYear)

	// Determine the Correct Day to Use

	birthDay := int(pin.BirthDay)
	if birthDay > 31 {
		birthDay = 1
	}

	return time.Date(birthYear, time.Month(pin.BirthMonth), birthDay, 0, 0, 0, 0, time.UTC)
}

// String pretty-prints the pension insurance number, adding spaces between
// area code, birthday, first letter of last name, and serial number.
//
// The returned string will have the following format:
//
//	AA DDMMYY L SSC
func (pin PensionInsuranceNumber) String() string {
	return fmt.Sprintf("%02d %02d%02d%02d %c %02d%d",
		pin.AreaCode, pin.BirthDay, pin.BirthMonth, pin.BirthYear, pin.LastNameLetter, pin.SerialNumber, pin.CheckDigit)
}

var _ encoding.TextMarshaler = PensionInsuranceNumber{}

// Compact renders a compact representation of the pension insurance number.
//
// In contrast to String(), MarshalText does not add spaces between the
// different parts of the pension insurance number.
//
// The returned string will have the following format:
//
//	AADDMMYYLSSC
func (pin PensionInsuranceNumber) Compact() string {
	return fmt.Sprintf("%02d%02d%02d%02d%c%02d%d",
		pin.AreaCode, pin.BirthDay, pin.BirthMonth, pin.BirthYear, pin.LastNameLetter, pin.SerialNumber, pin.CheckDigit)
}

func (pin PensionInsuranceNumber) MarshalText() ([]byte, error) {
	s := fmt.Sprintf("%02d%02d%02d%02d%c%02d%d",
		pin.AreaCode, pin.BirthDay, pin.BirthMonth, pin.BirthYear, pin.LastNameLetter, pin.SerialNumber, pin.CheckDigit)
	return []byte(s), nil
}

var _ encoding.TextUnmarshaler = (*PensionInsuranceNumber)(nil)

func (pin *PensionInsuranceNumber) UnmarshalText(text []byte) error {
	parsed, err := Parse(string(text))
	if err != nil {
		return err
	}

	*pin = parsed
	return nil
}
