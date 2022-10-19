package inter2

import "compiler-frontend/lexer"

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

func (a Or) Gen() {

}
