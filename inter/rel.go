/*
	rel.go

	implements <, <=, ==, !=, =>, >
	without coercions for simplicity
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Rel struct {
	Logical
}

func NewRel(tok lexer.Token, expr1, expr2 Expr) Rel {
	return Rel{
		Logical: NewLogical(tok, expr1, expr2),
	}
}

func (r Rel) check(t1, t2 lexer.Type) lexer.Type {
	if t1 == t2 {
		return lexer.Bool
	} else {
		return lexer.NullType
	}
}

func (r Rel) jumping(trueLabel, falseLabel int) {
	a := r.expr1.reduce()
	b := r.expr2.reduce()

	test := fmt.Sprintf("%s %s %s", a.toString(), r.operator.String(), b.toString())
	r.emitjumps(test, trueLabel, falseLabel)
}
