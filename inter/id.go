/*
	id.go

	Id structs are leaf nodes in the Abstract Syntax Tree
*/

package inter

import (
	"compiler-frontend/lexer"
)

type Id struct {
	offset int // holds relative address of the identifier
	Expr       // inherit default implementations of gen() and reduce()
}

func NewId(w lexer.Token, t lexer.Type, o int) Id {
	return Id{
		offset: o,
		Expr:   NewExpr(w, t),
	}
}
