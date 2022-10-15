/*
	not.go

*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Not struct {
	Logical
}

func NewNot(tok lexer.Token, expr Expr) Not {
	return Not{
		Logical: NewLogical(tok, expr, expr),
	}
}

func (n *Not) jumping(trueLabel, falseLabel int) {
	// call expr.jumping with the true and false exit labels reversed
	n.expr1.jumping(falseLabel, trueLabel)
}

func (n *Not) toString() string {
	return fmt.Sprintf("%s %s", n.operator.String(), n.expr2.toString())
}
