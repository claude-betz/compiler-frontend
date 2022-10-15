/*
	expr.go
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Expr struct {
	lexerLine int
	operator  lexer.Token
	typ       lexer.Type
}

func NewExpr(t lexer.Token, typ lexer.Type) Expr {
	return Expr{
		lexerLine: lexer.LexerLine, // current line of lexical analyser
		operator:  t,
		typ:       typ,
	}
}

func (e Expr) gen() Expr {
	return e
}

func (e Expr) reduce() Expr {
	return e
}

func (e Expr) jumping(trueLabel, falseLabel int) {
	e.emitjumps(e.toString(), trueLabel, falseLabel)
}

func (e Expr) emitjumps(testVar string, trueLabel, falseLabel int) {
	if trueLabel != 0 && falseLabel != 0 { // conditions for both ifTrue ifFalse
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
		ifFalse := fmt.Sprintf("goto L%d", falseLabel)

		emit(ifTrue)
		emit(ifFalse)
	} else if trueLabel != 0 { // if true
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLabel)
		emit(ifTrue)
	} else if falseLabel != 0 { // if false
		ifFalse := fmt.Sprintf("iffalse %s goto L%d", testVar, falseLabel)
		emit(ifFalse)
	} else {
		// nothing
	}
}

func (e Expr) toString() string {
	return (e.operator).String()
}
