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

func NewOp(t *lexer.Token, typ *lexer.Type) Op {
	return Op{
		NewExpr(t, typ),
	}
}

func (o Op) reduce() Temp {
	x := o.gen()
	t := NewTemp(o.typ)
	emit(fmt.Sprintf("%s = %s", t.toString(), x.toString()))

	return t
}
