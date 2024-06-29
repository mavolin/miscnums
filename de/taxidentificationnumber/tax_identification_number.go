// Package taxidentificationnumber provides parsing and validation for German
// Tax Identification Numbers (Steuerliche Identifikationsnummer).
package taxidentificationnumber

import (
	"encoding"
	"strconv"
)

// TaxIdentificationNumber is a German Tax Identification Number (Steuerliche
// Identifikationsnummer).
//
// It does not contain any spaces.
type TaxIdentificationNumber uint64

// String returns a human-readable representation of the tax ID, by separating
// the tax ID into its parts.
//
// The returned string is in the format "XX XXX XXX XXX".
//
// If the TaxIdentificationNumber is invalid, it is returned as is.
//
// It is guaranteed that for any valid TaxIdentificationNumber, the returned string can be parsed
// back into the same TaxIdentificationNumber.
func (id TaxIdentificationNumber) String() string {
	if id < 10_000_000_000 || id > 99_999_999_999 {
		return strconv.FormatUint(uint64(id), 10)
	}

	s := strconv.FormatUint(uint64(id), 10)
	return s[:2] + " " + s[2:5] + " " + s[5:8] + " " + s[8:]
}

var _ encoding.TextMarshaler = TaxIdentificationNumber(0)

func (id TaxIdentificationNumber) Compact() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id TaxIdentificationNumber) MarshalText() ([]byte, error) {
	return []byte(id.Compact()), nil
}

var _ encoding.TextUnmarshaler = (*TaxIdentificationNumber)(nil)

func (id *TaxIdentificationNumber) UnmarshalText(text []byte) error {
	parsedID, err := Parse(string(text))
	if err != nil {
		return err
	}

	*id = parsedID
	return nil
}
