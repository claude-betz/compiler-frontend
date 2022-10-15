/*
	set.go

	implements assignments with an identifier on the left side and an expression
	on the right side
*/

package inter

import (
	"compiler-frontend/lexer"
	"errors"
	"fmt"
)

type Set struct {
	id   Id
	expr Expr
}

func NewSet(id *Id, expr *Expr) Set {
	// check types
	_, err := checkTypes(*id.typ, *expr.typ)
	if err != nil {
		fmt.Println("[error] set.go - type error")
	}

	return Set{
		id:   *id,
		expr: *expr,
	}
}

func (s Set) gen(before, after int) {
	emit(fmt.Sprintf("%s = %s", s.id.toString(), s.expr.gen().toString()))
}

func checkTypes(t1, t2 lexer.Type) (lexer.Type, error) {
	if lexer.Numeric(t1) && lexer.Numeric(t2) {
		return t2, nil
	} else if t1 == lexer.Bool && t2 == lexer.Bool {
		return t2, nil
	} else {
		return t2, errors.New("type error")
	}
}
