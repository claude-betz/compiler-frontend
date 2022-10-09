/*
	main.go

	main package for running compiler front-end
*/

package main

import (
	"bufio"
	"bytes"
	"compiler-frontend/lexer"
	"fmt"
)

const (
	program = " { for ( 4 + 5 ) { int x; } } "
)

func main() {
	buf := bytes.NewBufferString(program)
	reader := bufio.NewReader(buf)

	// Lexical Analyser
	lex, err := lexer.NewLexer(reader)
	if err != nil {
		fmt.Println("error initialising lexer")
	}

	for i := 1; i <= 13; i++ {
		token := lex.Scan()
		fmt.Printf("tag: %s, value: %s\n", token.Tag().String(), token.Value())
	}
}
