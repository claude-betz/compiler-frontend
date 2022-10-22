/*
	parser.go

	Class to implement parser to construct
*/

package parser

import (
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

func (p *Parser) matchTokenTag(tag lexer.Tag) {
	if p.lookahead.Tag() == tag { // matched
		// advance lookahead
		fmt.Print(p.lookahead.Value() + " ")
		p.advanceLookahead()
		return
	} else {
		// error
	}
}

func (p *Parser) Program() {
	p.block()

	fmt.Println("success!")
}

func (p *Parser) block() {
	p.matchCharacter("{")

	// save symbol table from previous scope
	s := p.top

	// new symbol table for new scope
	p.top = symbol.NewEnv(s)

	p.decls()
	p.stmts()
	p.matchCharacter("}")

	// assign saved symbol table
	p.top = s
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

func (p *Parser) stmts() {
	if p.lookahead.Value() == "}" {
		return
	}
	p.stmt()
	p.stmts()
}

func (p *Parser) stmt() {
	switch p.lookahead.Tag() {
	case (lexer.ID):
		p.loc()
		p.matchCharacter("=")
		p.bool()
		p.matchCharacter(";")
	case (lexer.IF):
		p.matchTokenTag(lexer.IF)
		p.matchCharacter("(")
		p.bool()
		p.matchCharacter(")")
		p.stmt()
		return
	case (lexer.WHILE):
	case (lexer.CHARACTER):
		if p.lookahead.Value() == "{" {
			p.block()
		}
	}
}

func (p *Parser) loc() {
	p.matchTokenTag(lexer.ID) // match ID
	p.restLoc()
}

func (p *Parser) restLoc() {
	if p.lookahead.Value() == "[" { // match "["
		p.matchCharacter("[")
		p.bool()
		p.matchCharacter("]")
		p.restLoc()
	}
}

func (p *Parser) bool() {
	p.join()
	p.restBool()
}

func (p *Parser) restBool() {
	lookaheadTag := p.lookahead.Tag()
	if lookaheadTag == lexer.OR { // match "||"
		p.matchTokenTag(lexer.OR)
		p.join()
		p.restBool()
	}
}

func (p *Parser) join() {
	p.equality()
	p.restJoin()
}

func (p *Parser) restJoin() {
	lookaheadTag := p.lookahead.Tag()
	if lookaheadTag == lexer.AND { // match "&&"
		p.matchTokenTag(lexer.AND)
		p.equality()
		p.restJoin()
	}
}

func (p *Parser) equality() {
	p.rel()
	p.restEquality()
}

func (p *Parser) restEquality() {
	lookaheadTag := p.lookahead.Tag()

	// match "==" or "!="
	if lookaheadTag == lexer.EQUAL_TO || lookaheadTag == lexer.NOT_EQUAL_TO {
		p.matchTokenTag(lookaheadTag)
		p.rel()
		p.restEquality()
	}
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

func (p *Parser) factor() {
	l := p.lookahead

	if l.Tag().String() == "(" {
		p.matchCharacter("(")
		p.expr()
		p.matchCharacter(")")
		return
	}

	if l.Tag() == lexer.NUM || l.Tag() == lexer.ID {
		p.matchTokenTag(l.Tag())
		return
	}
}
