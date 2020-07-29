package parser

import (
	"fmt"
	"regexp/token"
)

type NodeType int

const (
	ND_SYMBOL NodeType = iota // 'a', 't', 'D',..
	ND_UNION                  // '|'
	ND_CONCAT                 // '・'
	ND_STAR                   // '*'
)

type Parser struct {
	tokens      []token.Token
	curTokenNum int
	nextStateID int
}

type Node struct {
	Type  NodeType
	Lhs   *Node
	Rhs   *Node
	Value rune
}

func newNode(t NodeType, v rune) *Node {
	return &Node{
		Type:  t,
		Value: v,
	}
}

func newNodeWithLR(t NodeType, v rune, lhs, rhs *Node) *Node {
	n := newNode(t, v)
	n.Lhs = lhs
	n.Rhs = rhs
	return n
}

func NewParser(tokens []token.Token) Parser {
	return Parser{
		tokens:      tokens,
		curTokenNum: 0,
		nextStateID: 0,
	}
}

func (p *Parser) nextToken() {
	p.curTokenNum++
}

func (p Parser) getCurToken() token.Token {
	return p.tokens[p.curTokenNum]
}

func (p Parser) curTokenTypeIs(tt token.TokenType) bool {
	return (p.getCurToken().Type) == tt
}

func nodeToStr(n *Node) string {
	switch n.Type {
	case ND_UNION:
		return "Union"
	case ND_CONCAT:
		return "Concat"
	case ND_STAR:
		return "Star"
	default:
		return string(n.Value)
	}
}

func dumpDotForEachNode(n *Node) {
	if n.Lhs != nil {
		fmt.Printf("    %s -> %s\n", nodeToStr(n), nodeToStr(n.Lhs))
		dumpDotForEachNode(n.Lhs)
	}
	if n.Rhs != nil {
		fmt.Printf("    %s -> %s\n", nodeToStr(n), nodeToStr(n.Rhs))
		dumpDotForEachNode(n.Rhs)
	}
}

// for debug
func DumpDOT(n *Node) {
	fmt.Printf("digraph AST {\n")
	dumpDotForEachNode(n)
	fmt.Printf("}\n")
}

func (p *Parser) parseSymbol() *Node {
	var node *Node
	if p.curTokenTypeIs(token.TK_SYMBOL) {
		node = newNode(ND_SYMBOL, p.getCurToken().Value)
		p.nextToken()
	}
	return node
}

func (p *Parser) parseStar() *Node {
	n := p.parseSymbol()
	// TODO? need to handle multiple stars?
	if p.curTokenTypeIs(token.TK_STAR) {
		v := p.getCurToken().Value
		p.nextToken()
		n = newNodeWithLR(ND_STAR, v, n, nil)
	}
	return n
}

func (p *Parser) parseConcate() *Node {
	lhs := p.parseStar()
	for p.curTokenTypeIs(token.TK_CONCAT) {
		v := p.getCurToken().Value
		p.nextToken()
		lhs = newNodeWithLR(ND_CONCAT, v, lhs, p.parseStar())
	}
	return lhs
}

func (p *Parser) parseUnion() *Node {
	lhs := p.parseConcate()
	for p.curTokenTypeIs(token.TK_UNION) {
		v := p.getCurToken().Value
		p.nextToken()
		lhs = newNodeWithLR(ND_UNION, v, lhs, p.parseConcate())
	}
	return lhs
}

func Parse(tokens []token.Token) *Node {
	p := NewParser(tokens)
	return p.parseUnion()
}
