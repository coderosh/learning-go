package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gocalcrepl/interpreter"
	"gocalcrepl/parser"
)

func scan(s *bufio.Scanner) (string, error) {
	if s.Scan() {
		return s.Text(), nil
	}
	err := s.Err()
	if err == nil {
		err = io.EOF
	}
	return "", err
}

func run(value string) float64 {
	p := parser.New(value)
	ast := p.Parse()

	val := interpreter.Eval(ast)

	return val
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("> ")
		value, _ := scan(s)

		if value == "exit" {
			break
		}

		val := run(value)
		fmt.Println(val)
		fmt.Print("\n")
	}
}
