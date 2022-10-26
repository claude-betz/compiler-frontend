/*
	main.go

	main package for running compiler front-end
*/

package main

import (
	"bufio"
	"bytes"
	"compiler-frontend/lexer"
	"compiler-frontend/parser"
	"fmt"
)

const (
	// program = "{ if (a + b > 2) a = b[2] + 2; if (2 + 5 > b) b = a + 2; }"
	program = "{ while (a > b[2]) a = a - b[2]; if (a < 2) b = a; }"
)

func main() {
	buf := bytes.NewBufferString(program)
	reader := bufio.NewReader(buf)

	// Lexical Analys
	lex, err := lexer.NewLexer(reader)
	if err != nil {
		fmt.Println("error initialising lexer")
	}

	// for i := 1; i <= 13; i++ {
	// 	token := lex.Scan()
	// 	fmt.Printf("tag: %s, value: %s\n", token.Tag().String(), token.Value())
	// }

	// Parsing
	parser := parser.NewParser(lex)
	parser.Program()
}
