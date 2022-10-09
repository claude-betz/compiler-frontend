/*
	token.go

	implementation of tokens for use by lexer.go.
*/

package lexer

import (
	"strconv"
)

// Token interface to be implemented by different types of tokens
type Token interface {
	Tag() Tag
	Value() string
	String() string
}

type Word struct {
	tag    Tag
	lexeme string
}

func NewWord(tag Tag, lexeme string) Word {
	return Word{
		tag:    tag,
		lexeme: lexeme,
	}
}

func (w Word) Tag() Tag {
	return w.tag
}

func (w Word) Value() string {
	return w.lexeme
}

func (w Word) String() string {
	return "{" + w.Tag().String() + ":" + w.Value() + "}"
}

type Num struct {
	tag   Tag
	value int
}

func NewNum(tag Tag, value int) Num {
	return Num{
		tag:   tag,
		value: value,
	}
}

func (n Num) Tag() Tag {
	return n.tag
}

func (n Num) Value() string {
	return strconv.Itoa(n.value)
}

func (n Num) String() string {
	return "{" + n.Tag().String() + ":" + n.Value() + "}"
}

type Type struct {
	tag   Tag
	value string
}

func NewType(tag Tag, value string) Type {
	return Type{
		tag:   tag,
		value: value,
	}
}

func (t Type) Tag() Tag {
	return t.tag
}

func (t Type) Value() string {
	return t.value
}

func (t Type) String() string {
	return "{" + t.Tag().String() + ":" + t.Value() + "}"
}

type Char struct {
	tag   Tag
	value rune
}

func NewChar(tag Tag, value rune) Char {
	return Char{
		tag:   tag,
		value: value,
	}
}

func (c Char) Tag() Tag {
	return c.tag
}

func (c Char) Value() string {
	return string(c.value)
}

func (c Char) String() string {
	return "{" + c.Tag().String() + ":" + c.Value() + "}"
}
