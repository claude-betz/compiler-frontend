/*
	logical.go

	common functionality for Or, And, Not
*/

package inter

import "compiler-frontend/lexer"

type Logical struct {
	expr1 Expr
	expr2 Expr
}

func NewLogical(tok lexer.Token, expr1, expr2 Expr) Logical {
	typ := check(expr1.typ, expr2.typ)
	if typ == lexer.NullType {
		// error
	}

	return Logical{
		expr1: expr1,
		expr2: expr2,
	}
}

func (l Logical) gen() {
	// TODO
}

func check(t1, t2 lexer.Type) lexer.Type {
	if t1 == lexer.Bool && t2 == lexer.Bool {
		return lexer.Bool
	} else {
		return lexer.NullType
	}
}
