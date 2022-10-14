/*
	id.go
*/

package inter

import (
	"compiler-frontend/lexer"
)

type Id struct {
	offset int
	node   Node
}

func NewId(w *lexer.Token, t *lexer.Type, o int) Id {
	return Id{
		offset: o,
		node:   NewExpr(w, t),
	}
}
