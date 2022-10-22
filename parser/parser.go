/*
	parser.go

	Class to implement parser to construct
*/

package parser

import (
	"compiler-frontend/inter"
	"compiler-frontend/lexer"
	"compiler-frontend/symbol"
	"fmt"
)

type Parser struct {
	lexer     *lexer.Lexer
	lookahead lexer.Token
	top       *symbol.Env
}

func NewParser(lexer *lexer.Lexer) *Parser {
	lookahead := lexer.Scan()

	return &Parser{
		lexer:     lexer,
		lookahead: lookahead,
		top:       symbol.NewEnv(nil),
	}
}

func (p *Parser) advanceLookahead() {
	p.lookahead = p.lexer.Scan()
}

func (p *Parser) matchCharacter(r string) {
	if p.lookahead.Value() == r { // matched
		// advance lookahead
		fmt.Print(r + " ")
		p.advanceLookahead()
	} else {
		// error
	}
}

func (p *Parser) matchTokenTag(tag lexer.Tag) inter.Expr {
	lookahead := p.lookahead
	if lookahead.Tag() == tag { // matched
		// advance lookahead
		fmt.Print(lookahead.Value() + " ")
		p.advanceLookahead()

		id := inter.NewId(lookahead)
		return id
	} else {
		// error
		return nil
	}
}

func (p *Parser) Program() {
	fmt.Println("parsing source...")

	stmtNode := p.block()
	blockNode := inter.NewBlock(stmtNode)

	fmt.Println("\nsuccessfully parsed source.")

	// call generate at root
	fmt.Println("generating target code...")

	gen := blockNode.Gen()
	fmt.Println(gen)

	fmt.Println("successfully generated target code.")
}

func (p *Parser) block() inter.Stmt {
	p.matchCharacter("{")

	// save symbol table from previous scope
	s := p.top

	// new symbol table for new scope
	p.top = symbol.NewEnv(s)

	p.decls()
	stmtsNode := p.stmts()
	p.matchCharacter("}")

	// assign saved symbol table
	p.top = s

	// return statement node
	return stmtsNode
}

func (p *Parser) decls() {
	if p.lookahead.Tag() == lexer.PRIMITIVE {
		p.decl()
		p.decls()
	}
}

func (p *Parser) decl() {
	if p.lookahead.Tag() == lexer.PRIMITIVE {
		p.matchTokenTag(lexer.PRIMITIVE) // match TYPE
		p.matchTokenTag(lexer.ID)        // match ID
		p.matchCharacter(";")
	}
}

func (p *Parser) stmts() inter.Stmt {
	if p.lookahead.Value() == "}" {
		return nil
	}
	stmt := p.stmt()
	stmts := p.stmts()

	return inter.NewSeq(stmt, stmts)
}

func (p *Parser) stmt() inter.Stmt {
	switch p.lookahead.Tag() {
	case (lexer.ID):
		id := p.loc()
		p.matchCharacter("=")
		expr := p.bool()
		p.matchCharacter(";")
		return inter.NewAssign(id, expr)
	case (lexer.IF):
		p.matchTokenTag(lexer.IF)
		p.matchCharacter("(")
		p.bool()
		p.matchCharacter(")")
		p.stmt()
	case (lexer.WHILE):
	case (lexer.CHARACTER):
		if p.lookahead.Value() == "{" {
			p.block()
		}
	}
	return nil
}

func (p *Parser) loc() inter.Expr {
	id := p.matchTokenTag(lexer.ID) // match ID
	b := p.restLoc()

	// if access return new access
	if b != nil {
		return inter.NewAccess(id.(inter.Id), b)
	}

	// return id
	return id
}

func (p *Parser) restLoc() inter.Expr {
	if p.lookahead.Value() == "[" { // match "["
		p.matchCharacter("[")

		b := p.bool()

		p.matchCharacter("]")
		p.restLoc() // only allow 1D arrays for now
		return b
	}
	return nil
}

func (p *Parser) bool() inter.Expr {
	expr1 := p.join()

	for {
		if p.lookahead.Tag() == lexer.OR { // match "||"
			p.matchTokenTag(lexer.OR)
			expr2 := p.join()

			expr1 = inter.NewOr(lexer.Or, expr1, expr2)
		}
		break
	}

	// return expr
	return expr1
}

func (p *Parser) join() inter.Expr {
	expr1 := p.equality()

	for {
		if p.lookahead.Tag() == lexer.AND { // match "&&"
			p.matchTokenTag(lexer.AND)
			expr2 := p.equality()

			expr1 = inter.NewAnd(lexer.And, expr1, expr2)
		}
		break
	}

	// return expr1
	return expr1
}

func (p *Parser) equality() inter.Expr {
	// should be rel
	expr1 := p.factor()

	for {
		lookahead := p.lookahead

		// match "==" and "!="
		if lexer.EqMap[lookahead.Tag().String()] == true {
			token := p.matchTokenTag(lookahead.Tag())
			// should be rel
			expr2 := p.factor()

			expr1 = inter.NewEquality(token.Token(), expr1, expr2)
		}

		break
	}

	return expr1
}

func (p *Parser) rel() {
	p.expr()

	switch p.lookahead.Tag() {
	case lexer.LESS_THAN:
		p.matchTokenTag(lexer.LESS_THAN)
		p.expr()
	case lexer.LESS_THAN_EQUAL_TO:
		p.matchTokenTag(lexer.LESS_THAN_EQUAL_TO)
		p.expr()
	case lexer.GREATER_THAN_EQUAL_TO:
		p.matchTokenTag(lexer.GREATER_THAN_EQUAL_TO)
		p.expr()
	case lexer.GREATER_THAN:
		p.matchTokenTag(lexer.GREATER_THAN)
		p.expr()
	}
}

func (p *Parser) expr() {
	p.term()
	p.restExpr()
}

func (p *Parser) restExpr() {
	lookupTag := p.lookahead.Tag()

	if lookupTag.String() == lexer.ADD.String() || lookupTag.String() == lexer.SUBTRACT.String() {
		p.matchTokenTag(lookupTag)
		p.term()
		p.restExpr()
	}
}

func (p *Parser) term() {
	p.unary()
	p.restTerm()
}

func (p *Parser) restTerm() {
	lVal := p.lookahead.Value()

	// match "*" or "/"
	if lVal == lexer.MULTIPLY.String() || lVal == lexer.DIVIDE.String() {
		p.unary()
		p.restTerm()
	}
}

func (p *Parser) unary() {
	lVal := p.lookahead.Value()

	if lVal == lexer.NOT.String() || lVal == lexer.MINUS.String() {
		p.unary()
	} else {
		p.factor()
	}
}

func (p *Parser) factor() inter.Expr {
	l := p.lookahead

	// ( expr)
	if l.Tag().String() == "(" {
		p.matchCharacter("(")
		p.expr()
		p.matchCharacter(")")
	}

	// id or num
	if l.Tag() == lexer.NUM || l.Tag() == lexer.ID {
		p.matchTokenTag(l.Tag())
		return inter.NewId(l)
	}

	return nil
}
