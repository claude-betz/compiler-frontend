/*
	temp.go
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	count = 0
)

type Temp struct {
	number int
	Expr
}

func NewTemp(t *lexer.Type) Temp {
	// increase global count
	count++
	number := count

	return Temp{
		number: number,
		Expr:   NewExpr(nil, t),
	}
}

func (t Temp) toString() string {
	return fmt.Sprintf("t%d", t.number)
}
