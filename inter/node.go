/*
	node.go

	interface for
*/

package inter

import (
	"fmt"
)

var (
	lexerLine = 0
	labels    = 0
)

type Node interface {
	// returns term that can fit the right side of three-address instruction
	gen() Expr
}

func newLabel() {
	labels++
}

func emitLabel(i int) {
	fmt.Printf("L%d:", i)
}

func emit(s string) {
	fmt.Printf("\t%s\n", s)
}

func error(s string) {
	fmt.Printf("[ERROR] near line: %d: %s", lexerLine, s)
}
