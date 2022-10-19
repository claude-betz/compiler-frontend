/*
	id.go

	Id structs are leaf nodes in the Abstract Syntax Tree
*/

package inter2

import (
	"compiler-frontend/lexer"
)

type Id struct {
	offset int         // holds relative address of the identifier
	token  lexer.Token // holds token
}

func (i Id) exprNode() {}

func NewId(w lexer.Token, t lexer.Type, o int) Id {
	return Id{
		offset: o,
	}
}

func (i Id) TokenLiteral() lexer.Tag {
	return i.token.Tag()
}

func (i Id) Gen() string {
	return i.token.Value()
}
