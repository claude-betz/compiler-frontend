/*
	arith.go

	arithmetic operators
*/

package inter

import "compiler-frontend/lexer"

type Arith struct {
	expr1, expr2 Expr
	Op
}

func NewArith(tok *lexer.Token, e1, e2 Expr) Arith {
	// create arith object
	arith := Arith{
		expr1: e1,
		expr2: e2,
		Op:    NewOp(tok, nil),
	}

	// do type coercions

	// set type on arith

	// return
	return arith
}
