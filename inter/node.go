/*
	node.go

	interface for
*/

package inter

import (
	"fmt"
)

var (
	labels = 0 // global variable for labels used in jumps
)

type Node interface {
	// returns term that can fit the right side of three-address instruction
	gen() Expr
}

func newLabel() int {
	labels++
	return labels
}

func emitLabel(i int) {
	fmt.Printf("L%d:", i)
}

func emit(s string) {
	fmt.Printf("\t%s\n", s)
}

func error(s string, line int) {
	fmt.Printf("[ERROR] near line: %d: %s", line, s)
}
