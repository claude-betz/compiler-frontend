package inter2

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	labels = 0 // global variable for labels used in jumps
)

type Node interface {
	Token() lexer.Token // token
	Gen() string        // generate three address code
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

// // generate jumping code for boolean expressions
// func Emitjumps(testVar string, trueLabel, falseLabel int) string {
// 	if trueLabel != 0 && falseLabel != 0 { // conditions for both ifTrue ifFalse
// 		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
// 		ifFalse := fmt.Sprintf("goto L%d", falseLabel)

// 		Emit(ifTrue)
// 		Emit(ifFalse)
// 	} else if trueLabel != 0 { // if true
// 		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
// 		Emit(ifTrue)
// 	} else if falseLabel != 0 { // if false
// 		ifFalse := fmt.Sprintf("iffalse %s goto L%d", testVar, falseLabel)
// 		Emit(ifFalse)
// 	} else {
// 		// nothing
// 	}
// }

func LValue(expr Expr) Expr {
	tag := expr.Token().Tag()
	if tag == lexer.ID {
		return expr
	} else if tag == lexer.ASSIGN {
		assign := expr.(Access)
		return NewAccess(assign.id, RValue(assign.expr))
	} else {
		// error
		return nil
	}
}

func RValue(expr Expr) Expr {
	tag := expr.Token().Tag()
	if tag == lexer.ID || tag == lexer.NUM {
		return expr
	} else if lexer.BoolMap[tag.String()] == true {
		switch tag.String() {
		case lexer.OR.String():
			or := expr.(Or)
			rValExpr1 := RValue(or.expr1)
			rValExpr2 := RValue(or.expr2)

			s := fmt.Sprintf("%s or %s", rValExpr1.Gen(), rValExpr2.Gen())
			return NewTemp(s)
		case lexer.AND.String():
			and := expr.(And)
			rValExpr1 := RValue(and.expr1)
			rValExpr2 := RValue(and.expr2)

			s := fmt.Sprintf("%s or %s", rValExpr1.Gen(), rValExpr2.Gen())
			return NewTemp(s)
		}
	} else if tag == lexer.ACCESS {

	} else if tag == lexer.ASSIGN {

	}
}
