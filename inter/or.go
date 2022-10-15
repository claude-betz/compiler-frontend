/*
	or.go

	generates jumping code for expression E = E1 || E2
*/

package inter

import "compiler-frontend/lexer"

type Or struct {
	tok   lexer.Token
	expr1 Expr
	expr2 Expr
	Logical
}

func NewOr(tok lexer.Token, expr1, expr2 Expr) Or {
	return Or{
		Logical: NewLogical(tok, expr1, expr2),
	}
}

func (o Or) jumping(trueLabel, falseLabel int) {
	label := 0
	if trueLabel != 0 {
		label = 0
	} else {
		label = newLabel()
	}

	o.expr1.jumping(label, 0)
	o.expr2.jumping(trueLabel, falseLabel)

	if trueLabel == 0 {
		emitLabel(label)
	}
}
