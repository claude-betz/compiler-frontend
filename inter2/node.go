package inter2

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	labels = 0 // global variable for labels used in jumps
)

type Node interface {
	TokenLiteral() lexer.Token // token
	Gen()                      // generate three address code
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

func NewLabel() int {
	// increment global labels
	labels++

	return labels
}

func EmitLabel(i int) {
	fmt.Printf("L%d:", i)
}

func Emit(s string) string {
	return fmt.Sprintf("\t%s\n", s)
}

func printError(s string, line int) {
	fmt.Printf("[ERROR] near line: %d: %s\n", line, s)
}

// generate jumping code for boolean expressions
func Emitjumps(testVar string, trueLabel, falseLabel int) string {
	if trueLabel != 0 && falseLabel != 0 { // conditions for both ifTrue ifFalse
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
		ifFalse := fmt.Sprintf("goto L%d", falseLabel)

		Emit(ifTrue)
		Emit(ifFalse)
	} else if trueLabel != 0 { // if true
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
		Emit(ifTrue)
	} else if falseLabel != 0 { // if false
		ifFalse := fmt.Sprintf("iffalse %s goto L%d", testVar, falseLabel)
		Emit(ifFalse)
	} else {
		// nothing
	}
}

func LValue(expr Expr) Expr {
	tokenLiteral := expr.TokenLiteral().Tag()
	if tokenLiteral == lexer.ID {
		return expr
	} else if tokenLiteral == lexer.ASSIGN {
		assign := expr.(Access)
		return NewAccess(assign.id, RValue(assign.expr))
	} else {
		// error
	}
}

func RValue(expr Expr) Expr {

}
