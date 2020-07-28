package main

import (
	"regexp/nfa"
	"regexp/parser"
	"regexp/token"
)

func main() {
	regexp := "a"
	tokens := token.Tokenize(regexp)
	ast := parser.Parse(tokens)
	nfa := nfa.CreateNFA(ast)
	nfa.DumpDOT()
}
