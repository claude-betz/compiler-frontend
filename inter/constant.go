/*
	constant.go
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	True  = NewConstant(lexer.True, lexer.Bool)
	False = NewConstant(lexer.False, lexer.Bool)
)

type Constant struct {
	Expr
}

func NewConstant(tok lexer.Token, typ lexer.Type) Constant {
	return Constant{
		Expr: NewExpr(tok, typ),
	}
}

func NewIntConstant(i int) Constant {
	num := lexer.NewNum(lexer.NUM, i)
	return Constant{
		Expr: NewExpr(num, lexer.Int),
	}
}

func (c Constant) jumping(trueLine, falseLine int) {
	if c == True && trueLine != 0 {
		emit(fmt.Sprintf("goto L%d", trueLine))
	} else if c == False && falseLine != 0 {
		emit(fmt.Sprintf("goto L%d", falseLine))
	} else {
		// nothing
	}
}
