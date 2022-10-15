/*
	if.go
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type If struct {
	expr Expr
	stmt Stmt
	Stmt
}

func NewIf(expr Expr, stmt Stmt) If {
	// type check
	if *expr.typ != lexer.Bool {
		// error
		fmt.Println("If statement requires boolean")
	}

	return If{
		expr: expr,
		stmt: stmt,
	}
}

func (i If) gen(before, after int) {
	label := newLabel()      // label for code of stmt
	i.expr.jumping(0, after) // fall through on true, goto after on false

	emitLabel(label)         // becomes the before for the generation of stmt
	i.stmt.gen(label, after) // in the if fall through
}
