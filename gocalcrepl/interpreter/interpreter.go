package interpreter

import (
	"strconv"

	"gocalcrepl/parser"
)

func Eval(ast parser.Node) float64 {
	aType := ast["type"]

	switch aType {
	case "Expression":
		return expression(ast)
	case "Number":
		value := ast["value"].(string)
		num, _ := strconv.ParseFloat(value, 64)
		return num
	}

	return 0
}

func expression(node parser.Node) float64 {
	lhs := Eval(node["left"].(parser.Node))
	rhs := Eval(node["right"].(parser.Node))
	operator := node["operator"].(string)

	switch operator {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return 0
}
