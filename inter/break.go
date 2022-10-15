/*
	break.go
*/

package inter

import "fmt"

type Break struct {
	stmt Stmt
}

func NewBreak() Break {
	if EnclosingStmt == NullStmt {
		// error unenclosed break
		fmt.Println("[error] break.go - unenclosed break")
	}
	return Break{
		stmt: EnclosingStmt,
	}
}

func (b Break) gen(before, after int) {
	emit(fmt.Sprintf("goto L%d", b.stmt.after))
}
