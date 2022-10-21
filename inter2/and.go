package inter2

import (
	"compiler-frontend/lexer"
	"fmt"
)

type And struct {
	token lexer.Token
	expr1 Expr
	expr2 Expr
}

func NewAnd(token lexer.Token, expr1, expr2 Expr) And {
	return And{
		token: token,
		expr1: expr1,
		expr2: expr2,
	}
}

func (a And) exprNode() {}

func (a And) Token() lexer.Token {
	return a.token
}

func (a And) Gen() string {
	rVal1 := RValue(a.expr1)
	rVal2 := RValue(a.expr2)

	return fmt.Sprintf("%s and %s", rVal1.Gen(), rVal2.Gen())
}
