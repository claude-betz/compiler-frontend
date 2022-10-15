/*
	stmt.go
*/

package inter

var (
	NullStmt  = NewStmt() // represents an empty sequence of statements
	Enclosing = NullStmt  // used for break statements
)

type Stmt struct {
	after int
	Node
}

func NewStmt() Stmt {
	return Stmt{
		after: 0,
	}
}

func (s Stmt) gen(before, after int) {

}
