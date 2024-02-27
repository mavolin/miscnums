package iban

import "testing"

func TestParse(t *testing.T) {
	// https://ibanvalidieren.de/beispiele.html
	// 2023-01-22
	// https://www.iban.com/structure
	// 2024-02-27
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "DE02120300000000202051", Expect: "DE02 1203 0000 0000 2020 51"},
		{In: "DE02500105170137075030", Expect: "DE02 5001 0517 0137 0750 30"},
		{In: "DE88100900001234567892", Expect: "DE88 1009 0000 1234 5678 92"},
		{In: "AT026000000001349870", Expect: "AT02 6000 0000 0134 9870"},
		{In: "AT021420020010147558", Expect: "AT02 1420 0200 1014 7558"},
		{In: "AT023200000000641605", Expect: "AT02 3200 0000 0064 1605"},
		{In: "CH0209000000100013997", Expect: "CH02 0900 0000 1000 1399 7"},
		{In: "CH0204835000626882001", Expect: "CH02 0483 5000 6268 8200 1"},
		{In: "CH0200700110000387896", Expect: "CH02 0070 0110 0003 8789 6"},
		{In: "LI0208800000017197386", Expect: "LI02 0880 0000 0171 9738 6"},
		{In: "LI0508812105028570001", Expect: "LI05 0881 2105 0285 7000 1"},
		{In: "LI2608802001003488101", Expect: "LI26 0880 2001 0034 8810 1"},

		{In: "LI26 0 8802 0010 0348 8101", Expect: "LI26 0880 2001 0034 8810 1"},
		{In: "LI26 0  88 02  00 10  0348  8 101", Expect: "LI26 0880 2001 0034 8810 1"},

		// National Checksums

		{In: "AL47 2121 1009 0000 0002 3569 8741", Expect: "AL47 2121 1009 0000 0002 3569 8741"},
		{In: "AL35 2021 1109 0000 0000 0123 4567", Expect: "AL35 2021 1109 0000 0000 0123 4567"},
		{In: "BE68 5390 0754 7034", Expect: "BE68 5390 0754 7034"},
		{In: "BE71 0961 2345 6769", Expect: "BE71 0961 2345 6769"},
		{In: "BA39 1290 0794 0102 8494", Expect: "BA39 1290 0794 0102 8494"},
		{In: "HR12 1001 0051 8630 0016 0", Expect: "HR12 1001 0051 8630 0016 0"},
		{In: "HR17 2360 0001 1012 3456 5", Expect: "HR17 2360 0001 1012 3456 5"},
		{In: "CZ65 0800 0000 1920 0014 5399", Expect: "CZ65 0800 0000 1920 0014 5399"},
		{In: "CZ55 0800 0000 0012 3456 7899", Expect: "CZ55 0800 0000 0012 3456 7899"},
		{In: "TL38 0080 0123 4567 8910 157", Expect: "TL38 0080 0123 4567 8910 157"},
		{In: "TL38 0010 0123 4567 8910 106", Expect: "TL38 0010 0123 4567 8910 106"},
		{In: "EE38 2200 2210 2014 5685", Expect: "EE38 2200 2210 2014 5685"},
		{In: "EE32 1282 9733 6485 9699", Expect: "EE32 1282 9733 6485 9699"},
		{In: "EE47 1000 0010 2014 5685", Expect: "EE47 1000 0010 2014 5685"},
		{In: "EE90 1700 0170 0000 0006", Expect: "EE90 1700 0170 0000 0006"},
		{In: "EE97 5500 0005 5000 8329", Expect: "EE97 5500 0005 5000 8329"},
		{In: "EE44 3300 3384 0010 0007", Expect: "EE44 3300 3384 0010 0007"},
		{In: "FI21 1234 5600 0007 85", Expect: "FI21 1234 5600 0007 85"},
		{In: "FI14 1009 3000 1234 58", Expect: "FI14 1009 3000 1234 58"},
		{In: "FR14 2004 1010 0505 0001 3M02 606", Expect: "FR14 2004 1010 0505 0001 3M02 606"},
		{In: "FR76 3000 6000 0112 3456 7890 189", Expect: "FR76 3000 6000 0112 3456 7890 189"},
		{In: "HU42 1177 3016 1111 1018 0000 0000", Expect: "HU42 1177 3016 1111 1018 0000 0000"},
		{In: "HU93 1160 0006 0000 0000 1234 5676", Expect: "HU93 1160 0006 0000 0000 1234 5676"},
		{In: "IS14 0159 2600 7654 5510 7303 39", Expect: "IS14 0159 2600 7654 5510 7303 39"},
		{In: "IS75 0001 1212 3456 3108 9620 99", Expect: "IS75 0001 1212 3456 3108 9620 99"},
		{In: "IT60 X054 2811 1010 0000 0123 456", Expect: "IT60 X054 2811 1010 0000 0123 456"},
		{In: "IT21 Q054 2801 6000 0ABC D12Z E34", Expect: "IT21 Q054 2801 6000 0ABC D12Z E34"},
		{In: "IT30 C080 0001 0001 23VA LE45 6NA", Expect: "IT30 C080 0001 0001 23VA LE45 6NA"},
		{In: "IT11 V060 0003 2000 0001 1556 BFE", Expect: "IT11 V060 0003 2000 0001 1556 BFE"},
		{In: "IT21 J010 0516 0521 2005 0012 345", Expect: "IT21 J010 0516 0521 2005 0012 345"},
		{In: "MK07 2501 2000 0058 984", Expect: "MK07 2501 2000 0058 984"},
		{In: "MK07 2000 0278 5123 453", Expect: "MK07 2000 0278 5123 453"},
		{In: "ME25 5050 0001 2345 6789 51", Expect: "ME25 5050 0001 2345 6789 51"},
		{In: "NO93 8601 1117 947", Expect: "NO93 8601 1117 947"},
		{In: "NO83 3000 1234 567", Expect: "NO83 3000 1234 567"},
		{In: "PL61 1090 1014 0000 0712 1981 2874", Expect: "PL61 1090 1014 0000 0712 1981 2874"},
		{In: "PL10 1050 0099 7603 1234 5678 9123", Expect: "PL10 1050 0099 7603 1234 5678 9123"},
		{In: "PL27 1140 2004 0000 3002 0135 5387", Expect: "PL27 1140 2004 0000 3002 0135 5387"},
		{In: "PL25 1060 1028 2276 7272 1438 5741", Expect: "PL25 1060 1028 2276 7272 1438 5741"},
		{In: "MC58 1122 2000 0101 2345 6789 030", Expect: "MC58 1122 2000 0101 2345 6789 030"},
		{In: "MC58 1009 6180 7901 2345 6789 085", Expect: "MC58 1009 6180 7901 2345 6789 085"},
		{In: "PT50 0002 0123 1234 5678 9015 4", Expect: "PT50 0002 0123 1234 5678 9015 4"},
		{In: "PT50 0027 0000 0001 2345 6783 3", Expect: "PT50 0027 0000 0001 2345 6783 3"},
		{In: "PT50 0002 0023 0023 8430 0057 8", Expect: "PT50 0002 0023 0023 8430 0057 8"},
		{In: "PT50 0003 0109 0012 6570 1395 8", Expect: "PT50 0003 0109 0012 6570 1395 8"},
		{In: "PT50 0004 0501 0020 5001 0144 1", Expect: "PT50 0004 0501 0020 5001 0144 1"},
		{In: "PT50 0026 0000 0524 2186 0018 5", Expect: "PT50 0026 0000 0524 2186 0018 5"},
		{In: "RS35 2600 0560 1001 6113 79", Expect: "RS35 2600 0560 1001 6113 79"},
		{In: "RS35 1050 0812 3123 1231 73", Expect: "RS35 1050 0812 3123 1231 73"},
		{In: "SK31 1200 0000 1987 4263 7541", Expect: "SK31 1200 0000 1987 4263 7541"},
		{In: "SK89 7500 0000 0000 1234 5671", Expect: "SK89 7500 0000 0000 1234 5671"},
		{In: "ES91 2100 0418 4502 0005 1332", Expect: "ES91 2100 0418 4502 0005 1332"},
		{In: "ES79 2100 0813 6101 2345 6789", Expect: "ES79 2100 0813 6101 2345 6789"},
		{In: "SI56 2633 0001 2039 086", Expect: "SI56 2633 0001 2039 086"},
		{In: "SI56 1920 0123 4567 892", Expect: "SI56 1920 0123 4567 892"},
		{In: "SI56 1910 0000 0123 438", Expect: "SI56 1910 0000 0123 438"},
		{In: "SI56 0510 0800 0032 875", Expect: "SI56 0510 0800 0032 875"},
		{In: "SM76 P085 4009 8121 2345 6789 123", Expect: "SM76 P085 4009 8121 2345 6789 123"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			iban, err := Parse(c.In)
			if err != nil {
				t.Fatalf("Parse(%q): %s", c.In, err)
			}

			if iban.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, iban.String())
			}
		})
	}
}
