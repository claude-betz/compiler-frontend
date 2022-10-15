/*
	seq.go

	implements a sequence of statements
*/

package inter

type Seq struct {
	stmt1 Stmt
	stmt2 Stmt
}

func NewSeq(stmt1, stmt2 Stmt) Seq {
	return Seq{
		stmt1: stmt1,
		stmt2: stmt2,
	}
}

func (s Seq) gen(before, after int) {
	if s.stmt1 == NullStmt {
		s.stmt2.gen(before, after)
	} else if s.stmt2 == NullStmt {
		s.stmt1.gen(before, after)
	} else {
		label := newLabel()
		s.stmt1.gen(before, label)
		emitLabel(label)
		s.stmt2.gen(label, after)
	}
}
