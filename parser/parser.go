package parser

import (
	"regexp/token"
)

type NodeType int

const (
	ND_SYMBOL NodeType = iota // 'a', 't', 'D',..
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
	Value byte
}

func newNode(t NodeType, v byte) *Node {
	return &Node{
		Type:  t,
		Value: v,
	}
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

func (p *Parser) parseSymbol() *Node {
	var node *Node
	if p.curTokenTypeIs(token.TK_CHAR) {
		node = newNode(ND_SYMBOL, p.getCurToken().Value)
		p.nextToken()
	}
	return node
}

func Parse(tokens []token.Token) *Node {
	p := NewParser(tokens)
	return p.parseSymbol()
}
