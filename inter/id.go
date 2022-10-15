/*
	id.go
*/

package inter

import (
	"compiler-frontend/lexer"
)

type Id struct {
	offset int
	Expr
}

func NewId(w *lexer.Token, t *lexer.Type, o int) Id {
	return Id{
		offset: o,
		Expr:   NewExpr(w, t),
	}
}
