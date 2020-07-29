package token

import (
	"fmt"
	"os"
)

type TokenType int

const (
	TK_CHAR   TokenType = iota // 'a', 't', 'D',..
	TK_UNION                   // '|'
	TK_CONCAT                  // '・' (・ is usually omitted in regular expression)
	TK_EOF                     // EOF
)

type Token struct {
	Type  TokenType
	Value rune
}

func newToken(tt TokenType, value rune) Token {
	return Token{
		Type:  tt,
		Value: value,
	}
}

func isChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func lastTokenIsChar(tokens []Token) bool {
	if len(tokens) == 0 {
		return false
	}

	return tokens[len(tokens)-1].Type == TK_CHAR
}

func Tokenize(regexp string) []Token {
	tokens := []Token{}
	var t Token
	for _, c := range regexp {
		if isChar(c) {
			if lastTokenIsChar(tokens) {
				t = newToken(TK_CONCAT, '・')
				tokens = append(tokens, t)
			}
			t = newToken(TK_CHAR, c)
		} else if c == '|' {
			t = newToken(TK_UNION, c)
		} else {
			fmt.Printf("unexpected input: %c", c)
			os.Exit(1)
		}
		tokens = append(tokens, t)
	}
	tokens = append(tokens, newToken(TK_EOF, '\000'))
	return tokens
}
