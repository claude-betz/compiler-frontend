/*
	logical.go

	common functionality for Or, And, Not
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Logical struct {
	typ   lexer.Type
	expr1 Expr
	expr2 Expr
	Expr
}

func NewLogical(tok lexer.Token, expr1, expr2 Expr) Logical {
	typ := check(expr1.typ, expr2.typ)
	if typ == lexer.NullType {
		// error
	}

	return Logical{
		typ:   typ,
		expr1: expr1,
		expr2: expr2,
	}
}

func (l Logical) gen() Temp {
	falseLine := newLabel()
	after := newLabel()
	temp := NewTemp(l.typ)

	l.jumping(0, falseLine)
	emit(fmt.Sprintf("%s = true", temp.toString()))
	emit(fmt.Sprintf("goto L%d", after))
	emitLabel(falseLine)
	emit(fmt.Sprintf("%s = false", temp.toString()))
	emitLabel(after)
	return temp
}

func check(t1, t2 lexer.Type) lexer.Type {
	if t1 == lexer.Bool && t2 == lexer.Bool {
		return lexer.Bool
	} else {
		return lexer.NullType
	}
}
