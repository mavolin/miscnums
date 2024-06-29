package pensioninsurancenumber

import (
	"errors"
	"math"
	"strings"
)

var (
	ErrLength = errors.New("pension insurance number: must be 12 digits long")

	ErrAreaCode       = errors.New("pension insurance number: invalid area code")
	ErrBirthDay       = errors.New("pension insurance number: birthday must only contain digits")
	ErrBirthMonth     = errors.New("pension insurance number: birth month must be a number between '01' and '12'")
	ErrBirthYear      = errors.New("pension insurance number: birth year must contain only digits")
	ErrLastNameLetter = errors.New("pension insurance number: last name letter must be an ascii letter")
	ErrSerialNumber   = errors.New("pension insurance number: serial number must only contain digits")
	ErrCheckDigit     = errors.New("pension insurance number: check digit must be a digit")
)

// Parse parses the passed pension insurance number.
//
// It ignores spaces and '/'.
//
// Parse is case-insensitive, however, the case of letters in the returned
// PensionInsuranceNumber will always be uppercase.
//
// If Parse returns without an error, the pension insurance number is
// considered syntactically valid.
func Parse(s string) (PensionInsuranceNumber, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "/", "")

	if len(s) != 12 {
		return PensionInsuranceNumber{}, ErrLength
	}

	var pin PensionInsuranceNumber

	untypAreaCode, ok := parseTwoDigits(s[:2])
	if !ok {
		return PensionInsuranceNumber{}, ErrAreaCode
	}

	pin.AreaCode = AreaCode(untypAreaCode)
	if !pin.AreaCode.IsValid() {
		return PensionInsuranceNumber{}, ErrAreaCode
	}

	// can be 0-99, see doc of BirthDay
	pin.BirthDay, ok = parseTwoDigits(s[2:4])
	if !ok {
		return PensionInsuranceNumber{}, ErrBirthDay
	}

	pin.BirthMonth, ok = parseTwoDigits(s[4:6])
	if !ok || pin.BirthMonth < 1 || pin.BirthMonth > 12 {
		return PensionInsuranceNumber{}, ErrBirthMonth
	}

	pin.BirthYear, ok = parseTwoDigits(s[6:8])
	if !ok {
		return PensionInsuranceNumber{}, ErrBirthYear
	}

	pin.LastNameLetter = rune(s[8])
	if pin.LastNameLetter >= 'a' && pin.LastNameLetter <= 'z' {
		pin.LastNameLetter -= 'a' - 'A' // turn into uppercase
	} else if pin.LastNameLetter < 'A' || pin.LastNameLetter > 'Z' {
		return PensionInsuranceNumber{}, ErrLastNameLetter
	}

	pin.SerialNumber, ok = parseTwoDigits(s[9:11])
	if !ok {
		return PensionInsuranceNumber{}, ErrSerialNumber
	}

	pin.CheckDigit, ok = parseOneDigit(s[11:])
	if !ok {
		return PensionInsuranceNumber{}, ErrCheckDigit
	}

	if pin.CheckDigit != calcCheckDigit(pin) {
		return PensionInsuranceNumber{}, ErrCheckDigit
	}

	return pin, nil
}

// save ourselves pre-checks for '_', etc. when using strconv.ParseUint
func parseOneDigit(s string) (uint8, bool) {
	return parseDigit(s[0])
}

func parseTwoDigits(s string) (uint8, bool) {
	digit1, ok := parseDigit(s[0])
	if !ok {
		return 0, false
	}

	digit2, ok := parseDigit(s[1])
	if !ok {
		return 0, false
	}

	return digit1*10 + digit2, true
}

func parseDigit(digit byte) (uint8, bool) {
	if digit < '0' || digit > '9' {
		return 0, false
	}

	return digit - '0', true
}

func calcCheckDigit(pin PensionInsuranceNumber) uint8 {
	// https://de.wikipedia.org/wiki/Versicherungsnummer#Berechnung_der_Pr%C3%BCfziffer
	// 2023-01-22

	var sum int

	sum += int(twoDigitDigitSum(2 * nthDigit(uint8(pin.AreaCode), 2)))
	sum += int(nthDigit(uint8(pin.AreaCode), 1))
	sum += int(twoDigitDigitSum(2 * nthDigit(pin.BirthDay, 2)))
	sum += int(twoDigitDigitSum(5 * nthDigit(pin.BirthDay, 1)))
	sum += int(twoDigitDigitSum(7 * nthDigit(pin.BirthMonth, 2)))
	sum += int(nthDigit(pin.BirthMonth, 1))
	sum += int(twoDigitDigitSum(2 * nthDigit(pin.BirthYear, 2)))
	sum += int(nthDigit(pin.BirthYear, 1))

	numericLetter := uint8(pin.LastNameLetter - 'A' + 1)
	sum += int(twoDigitDigitSum(2 * nthDigit(numericLetter, 2)))
	sum += int(nthDigit(numericLetter, 1))

	sum += int(twoDigitDigitSum(2 * nthDigit(pin.SerialNumber, 2)))
	sum += int(nthDigit(pin.SerialNumber, 1))

	return uint8(sum % 10)
}

// nthDigit returns the nth digit of n, where n=1 would return the rightmost
// digit.
func nthDigit(num uint8, n int) uint8 {
	return (num / (uint8)(math.Pow10(n-1))) % 10
}

func twoDigitDigitSum(n uint8) uint8 {
	return n/10 + n%10
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
