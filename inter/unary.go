/*
	unary.go

	dealing with operators that have one operand
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Unary struct {
	e Expr
	Op
}

func NewUnary(tok *lexer.Token, expr *Expr) Unary {
	// do type coersions
	typ := lexer.Max(lexer.Int, *expr.typ)
	if typ == nil {
		// error
	}

	return Unary{
		e:  *expr,
		Op: NewOp(tok, typ),
	}
}

func (u Unary) gen() Unary {
	reduced := u.e.reduce()
	return NewUnary(u.operator, &reduced)
}

func (u Unary) toString() string {
	op := (*u.operator).String()
	return fmt.Sprintf("%s %s", op, u.e.toString())
}
