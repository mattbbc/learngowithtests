package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Name   string
	Arabic uint16
	Roman  string
}{
	{"1 becomes I", 1, "I"},
	{"2 becomes II", 2, "II"},
	{"3 becomes III", 3, "III"},
	{"4 becomes IV (can't repeat more than 3 times)", 4, "IV"},
	{"5 becomes V", 5, "V"},
	{"6 becomes VI", 6, "VI"},
	{"7 becomes VII", 7, "VII"},
	{"8 becomes VIII", 8, "VIII"},
	{"9 becomes IX", 9, "IX"},
	{"10 becomes X", 10, "X"},
	{"14 becomes XIV", 14, "XIV"},
	{"18 becomes XVIII", 18, "XVIII"},
	{"20 becomes XX", 20, "XX"},
	{"39 becomes XXXIX", 39, "XXXIX"},
	{"40 becomes XL", 40, "XL"},
	{"47 becomes XLVII", 47, "XLVII"},
	{"49 becomes XLIX", 49, "XLIX"},
	{"50 becomes L", 50, "L"},
	{"100 becomes C", 100, "C"},
	{"90 becomes XC", 90, "XC"},
	{"400 becomes CD", 400, "CD"},
	{"500 becomes D", 500, "D"},
	{"900 becomes CM", 900, "CM"},
	{"1000 becomes M", 1000, "M"},
	{"1984 becomes MCMLXXXIV", 1984, "MCMLXXXIV"},
	{"3999 becomes MMMCMXCIX", 3999, "MMMCMXCIX"},
	{"2014 becomes MMXIV", 2014, "MMXIV"},
	{"1006 becomes MVI", 1006, "MVI"},
	{"798 becomes DCCXCVIII", 798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, wanted %q", got, test.Roman)
			}
		})
	}
}

func TestConvertRomanToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, wanted %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log(arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
