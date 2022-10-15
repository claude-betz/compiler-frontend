/*
	op.go

	class representing operators
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Op struct {
	Expr
}

func NewOp(tok lexer.Token, typ lexer.Type) Op {
	return Op{
		NewExpr(tok, typ),
	}
}

// provides implementation of op which is users by Arith, Unary and Access
func (o Op) reduce() Temp {
	// generate a term
	x := o.gen()

	// create new temp name
	t := NewTemp(o.typ)

	// emit an instruction to assign generated term to new temp name
	emit(fmt.Sprintf("%s = %s", t.toString(), x.toString()))

	// return the temp
	return t
}
