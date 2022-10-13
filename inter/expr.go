/*
	expr.go
*/

package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

type Expr struct {
	lexLine int
	op      lexer.Token
	typee   lexer.Type
}

func (e *Expr) gen() *Expr {
	return e
}

func (e *Expr) reduce() *Expr {
	return e
}

func (e *Expr) jumping(t, f int) {
	e.emitjumps(e.toString(), t, f)
}

func (e *Expr) emitjumps(testVar string, t, f int) {
	if t != 0 && f != 0 { // conditions for both ifTrue ifFalse
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, t)
		ifFalse := fmt.Sprintf("goto L%d", f)

		emit(ifTrue)
		emit(ifFalse)
	} else if t != 0 { // if true
		ifTrue := fmt.Sprintf("if %s goto L%d", testVar, t)
		emit(ifTrue)
	} else if f != 0 { // if false
		ifFalse := fmt.Sprintf("iffalse %s goto L%d", testVar, f)
		emit(ifFalse)
	} else {
		// nothing
	}
}

func (e *Expr) toString() string {
	return e.op.String()
}
