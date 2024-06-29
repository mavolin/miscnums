package taxidentificationnumber

import "testing"

func TestParse(t *testing.T) {
	// https://www.zfa.deutsche-rentenversicherung-bund.de/de/Inhalt/public/4_ID/47_Pruefziffernberechnung/001_Pruefziffernberechnung.pdf?__blob=publicationFile&v=2
	// p.4, section 2
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "86095742719", Expect: "86 095 742 719"},
		{In: "47036892816", Expect: "47 036 892 816"},
		{In: "65929970489", Expect: "65 929 970 489"},
		{In: "65929970489", Expect: "65 929 970 489"},
		{In: "57549285017", Expect: "57 549 285 017"},
		{In: "25768131411", Expect: "25 768 131 411"},

		{In: "86 095 742 719", Expect: "86 095 742 719"},
		{In: "86/095/742/719", Expect: "86 095 742 719"},
		{In: "86 / 095 / 742 / 719", Expect: "86 095 742 719"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			tin, err := Parse(c.In)
			if err != nil {
				t.Errorf("Parse(%q): %s", c.In, err)
			}

			if tin.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, tin.String())
			}
		})
	}
}
