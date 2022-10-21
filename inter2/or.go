package inter2

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Or struct {
	token lexer.Token
	expr1 Expr
	expr2 Expr
}

func NewOr(token lexer.Token, expr1 Expr, expr2 Expr) Or {
	return Or{
		token: token,
		expr1: expr1,
		expr2: expr2,
	}
}

func (o Or) exprNode() {}

func (o Or) Token() lexer.Token {
	return o.token
}

func (o Or) Gen() string {
	rVal1 := RValue(o.expr1)
	rVal2 := RValue(o.expr2)

	return fmt.Sprintf("%s or %s", rVal1.Gen(), rVal2.Gen())
}
