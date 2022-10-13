/*
	id.go
*/

package inter

import "compiler-frontend/lexer"

type Id struct {
	offset int
	node   Node
}

func NewId(w lexer.Word, t lexer.Type, o int) Id {
	return Id{
		offset: o,
		node:   NewExpr(w, t),
	}
}

func (i Id) gen() Expr {
	return i.node.gen()
}
