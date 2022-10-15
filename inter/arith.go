/*
	arith.go

	arithmetic operators
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Arith struct {
	expr1, expr2 Expr
	Op
}

func NewArith(tok lexer.Token, expr1, expr2 Expr) Arith {
	// do type coercions
	typ := lexer.Max(expr1.typ, expr2.typ)
	if typ != lexer.NullType {
		// error
	}

	// create arith object
	return Arith{
		expr1: expr1,
		expr2: expr2,
		Op:    NewOp(tok, typ),
	}
}

func (a *Arith) gen() Arith {
	return NewArith(a.operator, a.expr1.reduce(), a.expr2.reduce())
}

func (a *Arith) toString() string {
	return fmt.Sprintf("%s %s %s", a.expr1.toString(), a.operator.String(), a.expr2.toString())
}
