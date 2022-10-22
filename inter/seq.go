package inter

import "compiler-frontend/lexer"

type Seq struct {
	stmt  Stmt
	stmts Stmt
}

func NewSeq(stmt1 Stmt, stmt2 Stmt) Seq {
	return Seq{
		stmt:  stmt1,
		stmts: stmt2,
	}
}

func (s Seq) stmtNode() {}

func (s Seq) Token() lexer.Token { return nil }

func (s Seq) Gen() string {
	// for now assume we pass in the same in both i.e no sequences - I just wrote the parser with this in it
	return s.stmt.Gen()
}
