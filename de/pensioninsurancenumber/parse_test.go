package pensioninsurancenumber

import "testing"

func TestParse(t *testing.T) {
	// https://de.wikipedia.org/wiki/Versicherungsnummer#Aufbau_der_Ziffern_von_der_Bereichsnummer_bis_zur_Seriennummer
	// 2023-01-22
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "15070649C103", Expect: "15 070649 C 103"},
		{In: "15 070649 C 103", Expect: "15 070649 C 103"},
		{In: "15/070649/C/103", Expect: "15 070649 C 103"},
		{In: "15 / 070649 / C / 103", Expect: "15 070649 C 103"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			bic, err := Parse(c.In)
			if err != nil {
				t.Errorf("Parse(%q): %s", c.In, err)
			}

			if bic.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, bic.String())
			}
		})
	}
}
