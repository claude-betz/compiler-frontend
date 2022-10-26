package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type If struct {
	expr Expr
	stmt Stmt
}

func NewIf(expr Expr, stmt Stmt) If {
	return If{
		expr: expr,
		stmt: stmt,
	}
}

func (i If) stmtNode() {}

func (i If) Token() lexer.Token {
	return nil
}

func (i If) Gen() string {
	l := NewLabel()
	after := EmitLabel(l)
	expr := RValue(i.expr)
	fmt.Printf("\tifFalse %s goto %s\n", expr.Gen(), after)
	stmt := i.stmt.Gen()
	fmt.Printf("\t%s\n", stmt)
	fmt.Printf("%s:", after)
	return ""
}
