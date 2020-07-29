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
}
