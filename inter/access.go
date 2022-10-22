package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Access struct {
	id   Id
	expr Expr
}

func NewAccess(id Id, expr Expr) Access {
	return Access{
		id:   id,
		expr: expr,
	}
}

func (a Access) exprNode() {}

func (a Access) Token() lexer.Token {
	return lexer.Access
}

func (a Access) Gen() string {
	lVal := LValue(a.id)
	rVal := RValue(a.expr)

	return fmt.Sprintf("%s[%s]", lVal.Gen(), rVal.Gen())
}
