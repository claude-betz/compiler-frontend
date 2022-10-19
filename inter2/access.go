package inter2

import "compiler-frontend/lexer"

type Access struct {
	id   Expr
	expr Expr
}

func NewAccess(id, expr Expr) *Access {
	return &Access{
		id:   id,
		expr: expr,
	}
}

func (a Access) exprNode() {}

func (a Access) TokenLiteral() lexer.Token {
	return lexer.Access
}

func (a Access) Gen() {

}
