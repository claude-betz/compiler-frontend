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
	program = "{ a[2] = b[2<1] > c || d  }"
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
