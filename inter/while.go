package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type While struct {
	expr Expr
	stmt Stmt
}

func NewWhile(expr Expr, stmt Stmt) While {
	return While{
		expr: expr,
		stmt: stmt,
	}
}

func (w While) stmtNode() {}

func (w While) Token() lexer.Token {
	return nil
}

func (w While) Gen() string {
	l := NewLabel()
	before := EmitLabel(l)
	fmt.Printf("%s\n", before)

	l2 := NewLabel()
	after := EmitLabel(l2)
	expr := w.expr.Gen()
	fmt.Printf("IfFalse %s goto %s\n", expr, after)
	stmt := w.stmt.Gen()
	fmt.Printf("%s\n", stmt)
	fmt.Printf("goto: %s\n", before)
	fmt.Printf("%s\n", after)
	return ""
}
