package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Equality struct {
	token lexer.Token
	expr1 Expr
	expr2 Expr
}

func NewEquality(token lexer.Token, expr1, expr2 Expr) Equality {
	return Equality{
		token: token,
		expr1: expr1,
		expr2: expr2,
	}
}

func (e Equality) exprNode() {}

func (e Equality) Token() lexer.Token {
	return e.token
}

func (e Equality) Gen() string {
	expr1 := RValue(e.expr1)
	expr2 := RValue(e.expr2)

	return fmt.Sprintf("%s %s %s", expr1.Gen(), e.token.Value(), expr2.Gen())
}
