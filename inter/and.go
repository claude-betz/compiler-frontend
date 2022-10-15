/*
	and.go

	generates jumping code for expression E = E1 && E2
*/

package inter

import "compiler-frontend/lexer"

type And struct {
	tok   lexer.Token
	expr1 Expr
	expr2 Expr
	Logical
}

func NewAnd(tok lexer.Token, expr1, expr2 Expr) And {
	return And{
		Logical: NewLogical(tok, expr1, expr2),
	}
}

func (a And) jumping(trueLabel, falseLabel int) {
	label := 0
	if falseLabel != 0 {
		label = falseLabel
	} else {
		label = newLabel()
	}

	a.expr1.jumping(0, label)
	a.expr2.jumping(trueLabel, falseLabel)

	if falseLabel == 0 {
		emitLabel(label)
	}
}
