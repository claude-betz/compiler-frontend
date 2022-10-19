/*
	temp.go
*/

package inter2

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	count = 0
)

type Temp struct {
	number int
	typ    lexer.Type
}

func NewTemp(t lexer.Type) *Temp {
	// increase global count
	count++

	return &Temp{
		number: count,
		typ:    t,
	}
}

func (t *Temp) exprNode() {}

func (t *Temp) TokenLiteral() string {
	return t.toString()
}

func (t *Temp) Gen() string {
	return t.toString()
}

func (t *Temp) toString() string {
	return fmt.Sprintf("t%d", t.number)
}
