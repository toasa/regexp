package nfa

import (
	"testing"

	"regexp/parser"
	"regexp/token"
)

func genNFA(regexp string) *NFA {
	tokens := token.Tokenize(regexp)
	ast := parser.Parse(tokens)
	return CreateNFA(ast)
}

type testUnit struct {
	input    string
	expected bool
}

func testRegExp(t *testing.T, regexp string, tests []testUnit) {
	nfa := genNFA(regexp)
	for _, test := range tests {
		if nfa.accept(test.input) != test.expected {
			if test.expected {
				t.Errorf("regexp is %s but NFA doesn't accept %s.\n", regexp, test.input)
			} else {
				t.Errorf("regexp is %s but NFA accepts %s.\n", regexp, test.input)
			}
		}
	}
}

func TestSymbol(t *testing.T) {
	regexp1 := "a"
	units1 := []testUnit{
		{"a", true},
		{"b", false},
	}
	testRegExp(t, regexp1, units1)
}

func TestUnion(t *testing.T) {
	regexp1 := "a|b"
	units1 := []testUnit{
		{"a", true},
		{"b", true},
		{"c", false},
		{"ab", false},
	}
	testRegExp(t, regexp1, units1)

	regexp2 := "a|b|c"
	units2 := []testUnit{
		{"a", true},
		{"b", true},
		{"c", true},
		{"ab", false},
		{"bc", false},
		{"ca", false},
		{"abc", false},
	}
	testRegExp(t, regexp2, units2)
}

func TestConcate(t *testing.T) {
	regexp := "ab"
	units := []testUnit{
		{"ab", true},
		{"a", false},
		{"b", false},
	}
	testRegExp(t, regexp, units)
}

func TestStar(t *testing.T) {
	regexp := "a*"
	units := []testUnit{
		{"", true},
		{"a", true},
		{"aaaaaa", true},
		{"b", false},
	}
	testRegExp(t, regexp, units)
}

func TestParent(t *testing.T) {
	regexp := "(ab)*"
	units := []testUnit{
		{"", true},
		{"ab", true},
		{"abab", true},
		{"aab", false},
		{"abb", false},
	}
	testRegExp(t, regexp, units)
}

func TestComprehensive(t *testing.T) {
	regexp1 := "ab|c"
	units1 := []testUnit{
		{"ab", true},
		{"c", true},
		{"a", false},
		{"ac", false},
	}
	testRegExp(t, regexp1, units1)

	regexp2 := "a|bc"
	units2 := []testUnit{
		{"a", true},
		{"bc", true},
		{"b", false},
		{"ab", false},
	}
	testRegExp(t, regexp2, units2)

	regexp3 := "a*|b"
	units3 := []testUnit{
		{"", true},
		{"a", true},
		{"aaa", true},
		{"b", true},
		{"ab", false},
		{"aaaab", false},
		{"bb", false},
	}
	testRegExp(t, regexp3, units3)

	regexp4 := "ab*"
	units4 := []testUnit{
		{"a", true},
		{"ab", true},
		{"abbbbbbb", true},
		{"", false},
		{"b", false},
	}
	testRegExp(t, regexp4, units4)

	regexp5 := "a|(bc)*"
	units5 := []testUnit{
		{"", true},
		{"a", true},
		{"bc", true},
		{"bcbcbc", true},
		{"ab", false},
		{"bca", false},
		{"ac", false},
	}
	testRegExp(t, regexp5, units5)

	regexp6 := "(a|bc)*"
	units6 := []testUnit{
		{"", true},
		{"a", true},
		{"aaa", true},
		{"bc", true},
		{"bcbcbc", true},
		{"bca", true},
		{"abc", true},
		{"abcbca", true},
		{"bcaabc", true},
		{"bac", false},
		{"aab", false},
		{"bcbcabcc", false},
	}
	testRegExp(t, regexp6, units6)
}
