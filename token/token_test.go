package token

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	inputs := []string{
		"a",
	}

	expecteds := [][]Token{
		[]Token{
			Token{TK_CHAR, 'a'},
			Token{TK_EOF, '\000'},
		},
	}

	for i, input := range inputs {
		actual := Tokenize(input)
		expected := expecteds[i]
		for j, actual_token := range actual {
			expected_token := expected[j]

			// compare token type
			if actual_token.Type != expected_token.Type {
				t.Errorf("token type wrong\n")
				t.Fatalf("expected %d, but got %d\n", expected_token.Type, actual_token.Type)
			}

			// compare token value
			if actual_token.Value != expected_token.Value {
				t.Errorf("token value wrong\n")
				t.Fatalf("expected %c, but got %c\n", expected_token.Value, actual_token.Value)
			}
		}
	}
}