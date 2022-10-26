package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Assign struct {
	id   Expr
	expr Expr
}

func NewAssign(id Expr, expr Expr) Assign {
	return Assign{
		id:   id,
		expr: expr,
	}
}

func (a Assign) exprNode() {}

func (a Assign) stmtNode() {}

func (a Assign) Token() lexer.Token {
	return lexer.Assign
}

func (a Assign) Gen() string {
	lVal := LValue(a.id)
	rVal := RValue(a.expr)

	return fmt.Sprintf("%s = %s", lVal.Gen(), rVal.Gen())
}
