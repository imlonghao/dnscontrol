package models

import (
	"strings"
	"testing"
)

func TestIsQuoted(t *testing.T) {
	tests := []struct {
		d1 string
		e1 bool
	}{
		{``, false},
		{`foo`, false},
		{`""`, true},
		{`"a"`, true},
		{`"bb"`, true},
		{`"ccc"`, true},
		{`"aaa" "bbb"`, true},
	}
	for i, test := range tests {
		r := isQuoted(test.d1)
		if r != test.e1 {
			t.Errorf("%v: expected (%v) got (%v)", i, test.e1, r)
		}
	}
}

func TestStripQuotes(t *testing.T) {
	tests := []struct {
		d1 string
		e1 string
	}{
		{``, ``},
		{`a`, `a`},
		{`bb`, `bb`},
		{`ccc`, `ccc`},
		{`dddd`, `dddd`},
		{`"A"`, `A`},
		{`"BB"`, `BB`},
		{`"CCC"`, `CCC`},
		{`"DDDD"`, `DDDD`},
		{`"EEEEE"`, `EEEEE`},
		{`"aaa" "bbb"`, `aaa" "bbb`},
	}
	for i, test := range tests {
		r := StripQuotes(test.d1)
		if r != test.e1 {
			t.Errorf("%v: expected (%v) got (%v)", i, test.e1, r)
		}
	}
}

func r(s string, c int) string { return strings.Repeat(s, c) }

func TestParseQuotedTxt(t *testing.T) {
	tests := []struct {
		d1 string
		e2 []string
	}{
		{`foo`, []string{`foo`}},
		{`"foo"`, []string{`foo`}},
		{`"foo bar"`, []string{`foo bar`}},
		{`foo bar`, []string{`foo bar`}},
		{`"aaa" "bbb"`, []string{`aaa`, `bbb`}},
		{`"a"a" "bbb"`, []string{`a"a`, `bbb`}},
		// Seen in live traffic:
		{
			"\"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB\"",
			[]string{r("B", 254)},
		},
		{
			"\"CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC\"",
			[]string{r("C", 255)},
		},
		{
			"\"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD\" \"D\"",
			[]string{r("D", 255), "D"},
		},
		{
			"\"EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE\" \"EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE\"",
			[]string{r("E", 255), r("E", 255)},
		},
		{
			"\"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF\" \"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF\" \"F\"",
			[]string{r("F", 255), r("F", 255), "F"},
		},
		{
			"\"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\" \"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\" \"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG\"",
			[]string{r("G", 255), r("G", 255), r("G", 255)},
		},
		{
			"\"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH\" \"H\"",
			[]string{r("H", 255), r("H", 255), r("H", 255), "H"},
		},
		{"\"quo'te\"", []string{`quo'te`}},
		{"\"blah`blah\"", []string{"blah`blah"}},
		// {"\"quo\\\"te\"", []string{`quo"te`}},
		//{"\"q\\\"uo\\\"te\"", []string{`q"uo"te`}},
		//{"\"backs\\\\lash\"", []string{`back\slash`}},
	}
	for i, test := range tests {
		ls := ParseQuotedTxt(test.d1)
		if len(ls) != len(test.e2) {
			t.Errorf("%v: expected LEN TxtStrings=q(%q) got q(%q)", i, test.e2, ls)
		}
		for i := range ls {
			if ls[i] != test.e2[i] {
				t.Errorf("%v: expected TxtStrings=q(%q) got q(%q)", i, test.e2, ls)
			}
		}
	}
}

func TestParseQuotedFields(t *testing.T) {
	tests := []struct {
		d1 string
		e1 []string
	}{
		{`1 2 3`, []string{`1`, `2`, `3`}},
		{`1 "2" 3`, []string{`1`, `2`, `3`}},
		{`1 2 "three 3"`, []string{`1`, `2`, `three 3`}},
		{`1 2 "qu\"te" "4"`, []string{`1`, `2`, `qu\"te`, `4`}},
		{`0 issue "letsencrypt.org; validationmethods=dns-01; accounturi=https://acme-v02.api.letsencrypt.org/acme/acct/1234"`, []string{`0`, `issue`, `letsencrypt.org; validationmethods=dns-01; accounturi=https://acme-v02.api.letsencrypt.org/acme/acct/1234`}},
	}
	for i, test := range tests {
		ls, _ := ParseQuotedFields(test.d1)
		// fmt.Printf("%v: expected TxtStrings:\nWANT: %v\n GOT: %v\n", i, test.e1, ls)
		if len(ls) != len(test.e1) {
			t.Errorf("%v: expected TxtStrings=(%v) got (%v)", i, test.e1, ls)
		}
		for i := range ls {
			if ls[i] != test.e1[i] {
				t.Errorf("%v: expected TxtStrings=(%v) got (%v)", i, test.e1, ls)
			}
		}
	}
}
