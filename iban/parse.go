package iban

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mavolin/standards/iso3166"
)

var (
	ErrLength           = errors.New("iban: an iban must be at least 5 and at most 34 characters long")
	ErrCountryCode      = errors.New("iban: invalid country code")
	ErrChecksum         = errors.New("iban: checksum does not match")
	ErrCountrySpecific  = errors.New("iban: bban does not match country-specific rules")
	ErrNationalChecksum = errors.New("iban: national checksum does not match")
	ErrBBAN             = errors.New("iban: bban must consist of only letters and digits")
)

//go:generate go run github.com/mavolin/standards/tools/codegen/bban_regexp

// Parse parses the passed IBAN.
//
// Spaces are ignored and input is treated as case-insensitive, however, the
// returned IBAN will always be uppercase.
//
// # Validation
//
// If Parse returns without an error, the IBAN is considered syntactically
// valid.
//
// This means that the country code is valid, the checksum is correct, and the
// IBAN does not exceed the 34-character limit.
//
// If country-specific rules are available, Parse will also validate the
// country-specific length of the IBAN, the correct format (i.e. ensure that
// numeric parts only contain digits, and alphabetic parts only contain
// letters), and the national checksum, if there is one.
// See [IBAN formats by country] for the rules employed by Parse.
//
// [IBAN formats by country]: https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
func Parse(s string) (IBAN, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToUpper(s)

	if len(s) < 5 || len(s) > 34 {
		return IBAN{}, ErrLength
	}

	var iban IBAN

	var err error
	iban.CountryCode, err = iso3166.ParseAlpha2(s[:2])
	if err != nil {
		return IBAN{}, fmt.Errorf("iban: invalid country code: %w", err)
	}
	if !iban.CountryCode.Status.IsAssigned() {
		return IBAN{}, ErrCountryCode
	}

	var ok bool
	iban.Checksum, ok = parseTwoDigits(s[2:4])
	if !ok {
		return IBAN{}, ErrChecksum
	}

	if !isValidIBAN(s) {
		return IBAN{}, ErrChecksum
	}

	iban.BBAN = s[4:]

	bbanRegexp := bbanRegexps[iban.CountryCode]
	if bbanRegexp != nil {
		matches := bbanRegexp.FindStringSubmatch(iban.BBAN)
		if matches == nil {
			return IBAN{}, ErrCountrySpecific
		}

		for i, name := range bbanRegexp.SubexpNames() {
			if name == "" {
				continue
			}

			switch name {
			case "balanceAccountNumber":
				iban.BalanceAccountNumber = matches[i]
			case "bankCode":
				iban.BankCode += matches[i]
			case "bicBankCode":
				iban.BankCode += matches[i]
			case "accountNumber":
				iban.AccountNumber = matches[i]
			case "ownerIdentificationNumber":
				iban.OwnerIdentificationNumber = matches[i]
			case "currencyCode":
				iban.CurrencyCode = matches[i]
			case "ownerAccountNumber":
				iban.OwnerAccountNumber = matches[i]
			case "accountNumberPrefix":
				iban.AccountNumberPrefix = matches[i]
			case "branchCode":
				iban.BranchCode = matches[i]
			case "accountType":
				iban.AccountType = matches[i]
			case "checksum":
				iban.NationalChecksum = matches[i]
			}
		}
	} else {
		for _, r := range iban.BBAN {
			if (r < '0' || r > '9') && (r < 'A' || r > 'Z') {
				return IBAN{}, ErrBBAN
			}
		}
	}

	// check if we have a national checksum f
	if f, ok := checksumFuncs[iban.CountryCode]; ok {
		if !f(iban) {
			return IBAN{}, ErrNationalChecksum
		}
	}

	return iban, nil
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

func isValidIBAN(s string) bool {
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#Validating_the_IBAN
	// https://en.wikipedia.org/wiki/Modulo_operation

	// 1. Move the four initial characters to the end of the string.
	s = s[4:] + s[:4]

	// 2. Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35.
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			toDigit := c - 'A' + 10
			s = s[:i] + twoDigitToString(toDigit) + s[i+1:]
			i++ // we replaced a single character with two, so we need to skip
		}
	}

	// 3. Interpret the string as a decimal integer and compute the remainder of that number on division by 97.
	var remainder uint16
	for _, r := range s {
		remainder = remainder*10 + uint16(r-'0')
		remainder %= 97
	}

	return remainder == 1
}

func twoDigitToString(i uint8) string {
	//goland:noinspection GoVetIntToStringConversion
	return string(i/10+'0') + string(i%10+'0')
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
