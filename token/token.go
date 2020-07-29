package token

import (
	"fmt"
	"os"
)

type TokenType int

const (
	TK_CHAR  TokenType = iota // 'a', 't', 'D',..
	TK_UNION                  // '|'
	TK_EOF                    // EOF
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

func Tokenize(regexp string) []Token {
	tokens := []Token{}
	var t Token
	for _, c := range regexp {
		if isChar(c) {
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
