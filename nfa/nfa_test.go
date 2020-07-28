package nfa

import (
	"testing"

	"regexp/parser"
	"regexp/token"
)

func genNFA(regexp string) NFA {
	tokens := token.Tokenize(regexp)
	ast := parser.Parse(tokens)
	return CreateNFA(ast)
}

func TestSymbol(t *testing.T) {
	regexp := "a"
	nfa := genNFA(regexp)

	testData := []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"b", false},
	}

	for _, data := range testData {
		if nfa.accept(data.input) != data.expected {
			t.Errorf("regexp is %s but NFA accepts %s\n", regexp, data.input)
		}
	}
}
