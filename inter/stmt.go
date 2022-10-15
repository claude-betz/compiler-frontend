/*
	stmt.go
*/

package inter

var (
	Null      = NewStmt() // represents an empty sequence of statements
	Enclosing = Null      // used for break statements
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
