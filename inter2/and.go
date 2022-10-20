package inter2

import "compiler-frontend/lexer"

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

}
