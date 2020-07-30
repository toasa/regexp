package main

import (
	"flag"
	"fmt"
	"os"
	"regexp/nfa"
	"regexp/parser"
	"regexp/token"
)

func main() {
	var (
		regexp   = flag.String("regexp", "", "regular expression (Must be specified)")
		input    = flag.String("input", "", "input string")
		dump_ast = flag.Bool("ast", false, "dump AST")
		dump_sd  = flag.Bool("state-diagram", false, "dump state diagram")
	)
	flag.Parse()

	if *regexp == "" {
		fmt.Printf("$ go run -regexp <Regular expression>\n")
		os.Exit(1)
	}

	tokens := token.Tokenize(*regexp)
	ast := parser.Parse(tokens)
	if *dump_ast {
		ast.DumpDOT()
	}
	nfa := nfa.CreateNFA(ast)
	if *dump_sd {
		nfa.DumpDOT()
	}

	if *input != "" {
		result := nfa.Accept(*input)
		if result {
			fmt.Printf("%s accepts %s.\n", *regexp, *input)
		} else {
			fmt.Printf("%s doesn't accept %s.\n", *regexp, *input)
		}
	}
}
