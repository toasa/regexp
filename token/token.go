package token

import (
	"fmt"
	"os"
)

type TokenType int

const (
	TK_CHAR TokenType = iota // 'a', 't', 'D',..
	TK_EOF                   // EOF
)

type Token struct {
	Type  TokenType
	Value byte
}

func newToken(tt TokenType, value byte) Token {
	return Token{
		Type:  tt,
		Value: value,
	}
}

func isChar(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func Tokenize(regexp string) []Token {
	tokens := []Token{}
	var t Token
	for i := 0; i < len(regexp); i++ {
		c := regexp[i]
		if isChar(c) {
			t = newToken(TK_CHAR, c)
		} else {
			fmt.Printf("unexpected input: %c", c)
			os.Exit(1)
		}
		tokens = append(tokens, t)
	}
	tokens = append(tokens, newToken(TK_EOF, '\000'))
	return tokens
}
