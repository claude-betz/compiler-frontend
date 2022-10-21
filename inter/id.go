/*
	id.go

	Id structs are leaf nodes in the Abstract Syntax Tree
*/

package inter

import (
	"compiler-frontend/lexer"
)

type Id struct {
	token lexer.Token // holds token
}

func (i Id) exprNode() {}

func NewId(w lexer.Token) Id {
	return Id{
		token: w,
	}
}

func (i Id) Token() lexer.Token {
	return i.token
}

func (i Id) Gen() string {
	return i.token.Value()
}
