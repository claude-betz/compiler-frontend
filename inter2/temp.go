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
	value  string
}

func NewTemp(s string) Temp {
	// increase global count
	count++

	return Temp{
		number: count,
		value:  s,
	}
}

func (t Temp) exprNode() {}

func (t Temp) Token() lexer.Token {
	return lexer.Temp
}

func (t Temp) Gen() string {
	return t.toString()
}

func (t Temp) toString() string {
	return fmt.Sprintf("t%d = %s", t.number, t.value)
}
