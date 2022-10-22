package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Expression struct {
	token lexer.Token
	expr1 Expr
	expr2 Expr
}

func NewExpression(token lexer.Token, expr1, expr2 Expr) Expression {
	return Expression{
		token: token,
		expr1: expr1,
		expr2: expr2,
	}
}

func (e Expression) exprNode() {}

func (e Expression) Token() lexer.Token {
	return e.token
}

func (e Expression) Gen() string {
	expr1 := RValue(e.expr1)
	expr2 := RValue(e.expr2)

	return fmt.Sprintf("%s %s %s", expr1.Gen(), e.token.Value(), expr2.Gen())
}
