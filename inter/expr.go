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
	operator  *lexer.Token
	typ       *lexer.Type
}

func NewExpr(t *lexer.Token, typ *lexer.Type) Expr {
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

func (e Expr) jumping(trueLine, falseLine int) {
	e.emitjumps(e.toString(), trueLine, falseLine)
}

func (e Expr) emitjumps(testVar string, trueLine, falseLine int) {
	if trueLine != 0 && falseLine != 0 { // conditions for both ifTrue ifFalse
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLine)
		ifFalse := fmt.Sprintf("goto L%d", falseLine)

		emit(ifTrue)
		emit(ifFalse)
	} else if trueLine != 0 { // if true
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, trueLine)
		emit(ifTrue)
	} else if falseLine != 0 { // if false
		ifFalse := fmt.Sprintf("iffalse %s goto L%d", testVar, falseLine)
		emit(ifFalse)
	} else {
		// nothing
	}
}

func (e Expr) toString() string {
	return (*e.operator).String()
}
