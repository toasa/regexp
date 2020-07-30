package main

import (
	"regexp/nfa"
	"regexp/parser"
	"regexp/token"
)

func main() {
	regexp := "(a|b)*aba"
	tokens := token.Tokenize(regexp)
	ast := parser.Parse(tokens)
	ast.DumpDOT()
	nfa := nfa.CreateNFA(ast)
	nfa.DumpDOT()
}
