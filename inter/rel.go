package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Rel struct {
	token lexer.Token
	expr1 Expr
	expr2 Expr
}

func NewRel(token lexer.Token, expr1, expr2 Expr) Rel {
	return Rel{
		token: token,
		expr1: expr1,
		expr2: expr2,
	}
}

func (r Rel) exprNode() {}

func (r Rel) Token() lexer.Token {
	return r.token
}

func (r Rel) Gen() string {
	expr1 := RValue(r.expr1)
	expr2 := RValue(r.expr2)

	return fmt.Sprintf("%s %s %s", expr1.Gen(), r.token.Value(), expr2.Gen())
}
