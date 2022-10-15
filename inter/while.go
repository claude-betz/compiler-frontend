/*
	while.go
*/

package inter

import "fmt"

type While struct {
	expr Expr
	stmt Stmt
	Stmt
}

func NewWhile(expr Expr, stmt Stmt) While {
	return While{
		expr: expr,
		stmt: stmt,
	}
}

func (w While) gen(before, after int) {
	w.after = after // save after label

	w.expr.jumping(0, after)

	label := newLabel() // label for stmt
	emitLabel(label)
	w.stmt.gen(label, before)

	emit(fmt.Sprintf("goto L%d", before))
}
