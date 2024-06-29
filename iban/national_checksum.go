package iban

import (
	"strconv"
	"strings"

	"github.com/mavolin/standards/iso3166"
)

var checksumFuncs = map[iso3166.Alpha2Code]func(IBAN) bool{
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	// 2023-01-22
	// https://bank-code.net/iban/country-list
	// 2021-01-22
	//
	// Quick note on the Wikipedia source:
	// What they are saying in the paragraph above the table is not quite right
	// (or at the very least incredibly misleading):
	// Wiki says that the default case, i.e. when there is no comment, is that
	// the algorithm is used on the account number.
	// That is a bit colloquial, since what they actually mean is the entire
	// BBAN, minus the national checksum.

	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	iso3166.AL: func(iban IBAN) bool {
		check := weighted(iban.BankCode+iban.BranchCode, 10, 9, 7, 3, 1)
		if check != 0 {
			check = 10 - check
		}
		return iban.NationalChecksum == strconv.Itoa(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	// only for fields, variant not further clarified
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 23)
	iso3166.BE: func(iban IBAN) bool {
		check, err := strconv.ParseUint(iban.BankCode+iban.AccountNumber, 10, 64)
		if err != nil {
			return false
		}
		check %= 97
		if check == 0 {
			check = 97
		}
		return iban.NationalChecksum == twoDigitStr(int(check))
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	iso3166.BA: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.BranchCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	iso3166.HR: func(iban IBAN) bool {
		return iso7064Mod11_10(iban.BankCode) == 9 && iso7064Mod11_10(iban.AccountNumber) == 9
	},
	iso3166.CZ: czech,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	// Only works w/o appending "00", wiki lists no source.
	iso3166.TL: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 41)
	// https://media.voog.com/0000/0042/1620/files/BBAN-IBAN_php.pdf
	// w/0 length check
	// All sources specify to run the weighting backwards.
	// I don't see any mathematical reason for that, but still using reverse
	// weighting in case of stupidity, especially considering that this is
	// more effort and why would you do that if it's not necessary?
	// Also note that Wikipedia uses different weights, namely the weights I
	// would use if running the weighting LTR (which btw completes all tests).
	// Still: 2 sources > 1 source
	iso3166.EE: func(iban IBAN) bool {
		check := weightedRTL(iban.BranchCode+iban.AccountNumber, 10, 7, 3, 1)
		if check != 0 {
			check = 10 - check
		}
		return iban.NationalChecksum == strconv.Itoa(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 47)
	iso3166.FI: func(iban IBAN) bool {
		check := luhn(iban.BankCode + iban.AccountNumber)
		return iban.NationalChecksum == string(rune('0'+check))
	},
	iso3166.FR: france,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// Hungary has checksums in two different places, which right now is not
	// handled, we only check the national checksum.
	iso3166.HU: func(iban IBAN) bool {
		check := weighted(iban.AccountNumber, 10, 9, 7, 3, 1)
		if check != 0 {
			check = 10 - check
		}
		return iban.NationalChecksum == string(rune('0'+check))
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 68)
	iso3166.IS: func(iban IBAN) bool {
		check := weighted(iban.OwnerIdentificationNumber[:8], 11, 3, 2, 7, 6, 5, 4, 3, 2)
		if check != 0 {
			check = 11 - check
		}
		return iban.OwnerIdentificationNumber[8] == byte('0'+check)
	},
	iso3166.IT: italy,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 94)
	iso3166.MK: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	iso3166.MC: france,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	iso3166.ME: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 105)
	// ECBS calls the concatenation of the bank code and account no. the
	// "account number", which is a bit confusing.
	// Otherwise, the same as wiki.
	iso3166.NO: func(iban IBAN) bool {
		var check int
		if strings.HasPrefix(iban.AccountNumber, "00") {
			check = weighted(iban.AccountNumber[2:], 11, 5, 4, 3, 2)
		} else {
			check = weighted(iban.BankCode+iban.AccountNumber, 11, 5, 4, 3, 2, 7, 6, 5, 4, 3, 2)
		}

		if check == 1 {
			return false
		}

		if check != 0 {
			check = 11 - check
		}
		return iban.NationalChecksum == strconv.Itoa(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// ECBS uses ISO13616, which is not publicly available.
	// Wikipedia works fine, so I think we're good.
	iso3166.PL: func(iban IBAN) bool {
		check := weighted(iban.BankCode+iban.BranchCode, 10, 3, 9, 7, 1)
		if check != 0 {
			check = 10 - check
		}
		return iban.NationalChecksum == strconv.Itoa(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 112)
	iso3166.PT: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.BranchCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	iso3166.SM: italy,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 120)
	// ECBS once again says account number, but means bank code + account no.
	iso3166.RS: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	iso3166.SK: czech,
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 126)
	iso3166.SI: func(iban IBAN) bool {
		check := iso7064Mod97_10(iban.BankCode + iban.BranchCode + iban.AccountNumber)
		return iban.NationalChecksum == twoDigitStr(check)
	},
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country
	// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 131)
	iso3166.ES: func(iban IBAN) bool {
		check1 := weighted(iban.BankCode+iban.BranchCode, 11, 4, 8, 5, 10, 9, 7, 3, 6)
		if check1 > 1 {
			check1 = 11 - check1
		}

		check2 := weighted(iban.AccountNumber, 11, 1, 2, 4, 8, 5, 10, 9, 7, 3, 6)
		if check2 > 1 {
			check2 = 11 - check2
		}

		return iban.NationalChecksum == string(rune('0'+check1))+string(rune('0'+check2))
	},
}

// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 34)
// Wiki has complement wrong, otherwise as above.
func czech(iban IBAN) bool {
	r := weighted(iban.BranchCode, 11, 10, 5, 8, 4, 2, 1)
	if r != 0 {
		return false
	}

	r = weighted(iban.AccountNumber, 11, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1)
	return r == 0
}

// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
// https://www.ecbs.org/Download/Tr201v3.9.pdf
// ECBS is honestly not very clear, but Wikipedia works, and we have sufficient
// test cases that I am confident in the implementation.
func france(iban IBAN) bool {
	s := franceTransliterate(iban.BankCode + iban.BranchCode + iban.AccountNumber + iban.NationalChecksum)
	return baseISO7064Mod97_10(s) == 0
}

func franceTransliterate(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		r := bs[i]
		switch {
		case r >= 'A' && r <= 'I':
			bs[i] = r - 'A' + '1'
		case r >= 'J' && r <= 'N':
			bs[i] = r - 'J' + '1'
		case r >= 'S' && r <= 'Z':
			bs[i] = r - 'S' + '2'
		}
	}

	return string(bs)
}

// https://www.ecbs.org/Download/Tr201v3.9.pdf (page 77)
// Officially called odd mapping, but since they consider the first digit
// to be odd, and 0 is even, I'm swapping the names.
var italyEvenMapping = [...]int{
	1, 0, 5, 7, 9, 13, 15, 17, 19, 21,
	2, 4, 18, 20, 11, 3, 6, 8, 12, 14,
	16, 10, 22, 25, 24, 23,
}

// Odd mapping is just ascending numbers.

func italy(iban IBAN) bool {
	if len(iban.NationalChecksum) != 1 {
		return false
	}
	checksum := int(iban.NationalChecksum[0] - 'A')

	var sum int
	for i, r := range iban.BankCode + iban.BranchCode + iban.AccountNumber {
		var num int
		if r >= 'A' && r <= 'Z' {
			num = int(r - 'A')
		} else {
			num = int(r - '0')
		}
		if i%2 == 0 {
			num = italyEvenMapping[num]
		}
		sum = (sum + num) % 26
	}
	return sum%26 == checksum
}

// ============================================================================
// Algorithms
// ======================================================================================

// ISO really made sure to gatekeep.
// Only reliable source I could find is in German.
// https://de.wikipedia.org/wiki/ISO/IEC_7064#Algorithmus_f%C3%BCr_reine_Systeme_mit_zwei_Pr%C3%BCfzeichen
//
//goland:noinspection GoSnakeCaseUsage
func iso7064Mod97_10(s string) int {
	return 98 - baseISO7064Mod97_10(s)
}

//goland:noinspection GoSnakeCaseUsage
func baseISO7064Mod97_10(s string) int {
	var prod int
	for _, char := range s {
		sum := digit(char) + prod
		prod = (sum * 10) % 97
	}
	return (prod * 10) % 97
}

// also from German Wikipedia
// https://de.wikipedia.org/wiki/ISO/IEC_7064#Algorithmus_f%C3%BCr_hybride_Systeme
//
//goland:noinspection GoSnakeCaseUsage
func iso7064Mod11_10(s string) int {
	prod := 10
	for _, r := range s {
		sum := (digit(r) + prod) % 10
		if sum == 0 {
			sum = 10
		}
		prod = (sum * 2) % 11
	}

	if prod == 1 {
		return 0
	}
	return 11 - prod
}

func digit(r rune) int {
	return int(r - '0')
}

func weighted(s string, mod int, weights ...int) int {
	var sum int
	for i, r := range s {
		sum += (digit(r) * weights[i%len(weights)]) % mod
	}
	return sum % mod
}

func weightedRTL(s string, mod int, weights ...int) int {
	var sum int
	for i, r := 0, len(s)-1; r >= 0; i, r = i+1, r-1 {
		sum += (digit(rune(s[r])) * weights[i%len(weights)]) % mod
	}
	return sum % mod
}

func luhn(s string) int {
	var sum int

	for i := len(s) - 1; i >= 0; i-- {
		n := int(s[i] - '0')
		if i%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum = (sum + n) % 10
	}

	return 10 - sum%10
}
