package inter

import "compiler-frontend/lexer"

type Seq struct {
	stmt1 Stmt
	stmt2 Stmt
}

func NewSeq(stmt1 Stmt, stmt2 Stmt) Seq {
	return Seq{
		stmt1: stmt1,
		stmt2: stmt2,
	}
}

func (s Seq) stmtNode() {}

func (s Seq) Token() lexer.Token { return nil }

func (s Seq) Gen() string {
	val := s.stmt1.Gen()

	// stmts can be epsilon according to grammar
	if s.stmt2 != nil {
		val = s.stmt2.Gen()
	}
	return val
}
