package taxidentificationnumber

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrLength = errors.New("tax identification number: must be 11 digits long")
	ErrSyntax = errors.New("tax identification number: must only contain digits, and may not start with 0")
	// ErrRepetition signifies that the parser found a number that appears
	// more than three times, or that it found two numbers that appear
	// more than once.
	ErrRepetition = errors.New("tax identification number: mustn't contain the same digit more than three times")
	ErrCheckDigit = errors.New("tax identification number: invalid check digit")
)

// tinRegexp is the regular expression used to validate string-formatted tax
// ids.
var tinRegexp = regexp.MustCompile(`^[1-9]\d{10}$`)

// Parse parses the passed German tax id.
//
// Spaces and '/' are ignored.
//
// If Parse returns without an error, the tax id is considered syntactically
// valid.
func Parse(s string) (TaxIdentificationNumber, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "/", "")

	// https://www.zfa.deutsche-rentenversicherung-bund.de/de/Inhalt/public/4_ID/47_Pruefziffernberechnung/001_Pruefziffernberechnung.pdf?__blob=publicationFile&v=2

	// besides the no-0-check, this is also necessary, because ParseUint
	// also accepts things like hex or octal numbers, or '_' as separators
	if !tinRegexp.MatchString(s) {
		return 0, ErrSyntax
	}

	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return ParseNum(n)
}

// ParseNum is the same as [Parse], but takes a number instead of a string.
func ParseNum(n uint64) (TaxIdentificationNumber, error) {
	// a tax id is elven digits long
	if n < 10_000_000_000 || n > 99_999_999_999 {
		return 0, ErrLength
	}

	if !isValidDigitRepetition(n) {
		return 0, ErrRepetition
	}

	if calcCheckDigit(n) != nthDigit(n, 1) {
		return 0, ErrCheckDigit
	}

	return TaxIdentificationNumber(n), nil
}

// isValidDigitRepetition checks that the passed tax identification number does
// not contain the same digit more than three times and that no digits other
// than one appear more than once.
func isValidDigitRepetition(n uint64) bool {
	var repCounts [10]uint8

	for i := 2; i <= 11; i++ {
		repCounts[nthDigit(n, i)]++
	}

	// var foundDigitThatAppearsMultipleTimes bool
	var foundMulti bool

	for i, count := range repCounts {
		if count > 3 { // three is the maximum number of repetitions allowed
			return false
		}

		if count > 1 { // we found the digit that appears multiple times
			if foundMulti {
				return false
			}

			foundMulti = true
		}

		// if a digit repeats three times, then at least one occurrence must
		// not be adjacent to the other two
		if count == 3 {
			numRep := 0

			for j := 2; j <= 11; j++ {
				// we found one of the three occurrences
				if nthDigit(n, j) == uint8(i) {
					numRep++ // increase the number of occurrences found

					// check if the next digit is another of the three
					// occurrences
					if nthDigit(n, j+1) != uint8(i) {
						// it's not; that means it's valid!
						break
					}

					if numRep >= 2 {
						// the first digit was adjacent to the second, and so
						// was the second to the third
						return false
					}
				}
			}
		}
	}

	return true
}

func calcCheckDigit(n uint64) uint8 {
	var product uint8 = 10
	for i := 11; i >= 2; i-- {
		sum := (nthDigit(n, i) + product) % 10
		if sum == 0 {
			sum = 10
		}

		product = (2 * sum) % 11
	}

	checkDigit := 11 - product
	if checkDigit == 10 {
		checkDigit = 0
	}

	return checkDigit
}

// nthDigit returns the nth digit of n, where n=1 would return the rightmost
// digit.
func nthDigit(num uint64, n int) uint8 {
	return uint8((num / (uint64)(math.Pow10(n-1))) % 10)
}

// IsValid validates that s represents a syntactically valid German tax id.
//
// Spaces are ignored.
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}

// IsValidNum is the same as [IsValid], but takes a number instead of a
// string.
func IsValidNum(n uint64) bool {
	_, err := ParseNum(n)
	return err == nil
}
