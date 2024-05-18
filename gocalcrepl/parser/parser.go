package parser

import (
	"fmt"
	"log"

	t "gocalcrepl/tokenizer"
)

type Parser struct {
	tokenizer t.Tokenizer
	current   t.Token
}

type Node map[string]any

func New(value string) Parser {
	tokenizer := t.New(value)
	parser := Parser{tokenizer: tokenizer}
	parser.current = parser.tokenizer.Next()

	return parser
}

func (p *Parser) eat(expectType t.TokenType) t.Token {
	cur := p.current

	if p.current.Type != expectType {
		log.Fatal(fmt.Sprintf("Expected type %v, got %v", expectType, cur.Type))
	}

	p.current = p.tokenizer.Next()

	return cur
}

func (p *Parser) Parse() Node {
	return p.parseAdditive()
}

func (p *Parser) parseAdditive() Node {
	left := p.parseMultiplicative()

	for p.current.Value == "-" || p.current.Value == "+" {
		operator := p.eat(t.Operand)

		left = Node{
			"type":     "Expression",
			"left":     left,
			"operator": operator.Value,
			"right":    p.parseMultiplicative(),
		}
	}

	return left
}

func (p *Parser) parseMultiplicative() Node {
	left := p.parseNumber()

	for p.current.Value == "*" || p.current.Value == "/" {
		operator := p.eat(t.Operand)

		left = Node{
			"type":     "Expression",
			"left":     left,
			"operator": operator.Value,
			"right":    p.parseNumber(),
		}
	}

	return left
}

func (p *Parser) parseNumber() Node {
	token := p.eat(t.Number)

	node := Node{
		"type":  "Number",
		"value": token.Value,
	}

	return node
}
