package inter

import "compiler-frontend/lexer"

type Block struct {
	stmts Stmt
}

func NewBlock(stmts Stmt) Block {
	return Block{
		stmts: stmts,
	}
}

func (b Block) Token() lexer.Token { return nil }

func (b Block) Gen() string {
	return b.stmts.Gen()
}
