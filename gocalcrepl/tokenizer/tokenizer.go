package tokenizer

type TokenType string

const (
	EOF     = "EOF"
	Number  = "Number"
	Operand = "Operand"
	Invalid = "Invalid"
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	Value  string
	cursor int
}

func New(value string) Tokenizer {
	return Tokenizer{Value: value, cursor: 0}
}

func (p *Tokenizer) Next() Token {
	if p.isEOF() {
		return Token{Type: EOF, Value: ""}
	}

	if p.isEmpty() {
		return p.Next()
	}

	num := p.getFloatNumber()
	if num != "" {
		return Token{Type: Number, Value: num}
	}

	operand := p.getOperand()
	if operand != "" {
		return Token{Type: Operand, Value: operand}
	}

	p.cursor++
	return Token{Type: Invalid, Value: string(p.Value[p.cursor])}
}

func (p *Tokenizer) getFloatNumber() string {
	num := p.getNumber()
	if num == "" || p.isEOF() {
		return num
	}

	cur := p.Value[p.cursor]

	if string(cur) == "." {
		p.cursor++
		num += "."
	}

	num += p.getNumber()

	return num
}

func (p *Tokenizer) getNumber() string {
	num := ""

	cur := p.Value[p.cursor]
	for isNumber(cur) {

		num = num + string(cur)

		p.cursor++

		if p.isEOF() {
			break
		}

		cur = p.Value[p.cursor]
	}

	return num
}

func (p *Tokenizer) getOperand() string {
	cur := string(p.Value[p.cursor])
	p.cursor++

	if cur == "+" || cur == "-" || cur == "*" || cur == "/" {
		return cur
	}

	return ""
}

func (p *Tokenizer) isEmpty() bool {
	cur := string(p.Value[p.cursor])

	isEmpty := cur == " " || cur == "\n" || cur == "\t"

	if isEmpty {
		p.cursor++
	}

	return isEmpty
}

func (p *Tokenizer) isEOF() bool {
	return p.cursor >= len(p.Value)
}

func isNumber(char byte) bool {
	return char >= 48 && char <= 57
}
