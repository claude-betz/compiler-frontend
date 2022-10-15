/*
	while.go
*/

package inter

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

func (w While) gen() {

}
