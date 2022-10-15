/*
	token.go

	implementation of tokens for use by lexer.go.
*/

package lexer

import (
	"strconv"
)

// convenience
var (
	// expressions
	// assign
	assign = NewWord(ASSIGN, ASSIGN.String())

	// boolean ops
	or  = NewWord(OR, OR.String())
	and = NewWord(AND, AND.String())

	// equality
	eq = NewWord(EQUAL_TO, EQUAL_TO.String())
	ne = NewWord(NOT_EQUAL_TO, NOT_EQUAL_TO.String())

	// relational
	lt = NewWord(LESS_THAN, LESS_THAN.String())
	le = NewWord(LESS_THAN_EQUAL_TO, LESS_THAN_EQUAL_TO.String())
	gt = NewWord(GREATER_THAN, GREATER_THAN.String())
	ge = NewWord(GREATER_THAN_EQUAL_TO, GREATER_THAN_EQUAL_TO.String())

	// expr
	add  = NewWord(ADD, ADD.String())
	diff = NewWord(SUBTRACT, ADD.String())

	// terms
	mul = NewWord(MULTIPLY, MULTIPLY.String())
	div = NewWord(DIVIDE, DIVIDE.String())

	// unary
	not   = NewWord(NOT, NOT.String())
	minus = NewWord(MINUS, MINUS.String())
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

var (
	True  = NewWord(TRUE, "true")
	False = NewWord(FALSE, "false")
)

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
	width int // width in words
	Word
}

var (
	Int      = NewType(PRIMITIVE, "int", 4)
	Float    = NewType(PRIMITIVE, "float", 8)
	Char     = NewType(PRIMITIVE, "char", 1)
	Bool     = NewType(PRIMITIVE, "bool", 1)
	NullType = NewType(Undefined, "undefined", 0)
)

func NewType(tag Tag, value string, width int) Type {
	return Type{
		width: width,
		Word:  NewWord(tag, value),
	}
}

func (t Type) Tag() Tag {
	return t.tag
}

func (t Type) Value() string {
	return t.lexeme
}

func (t Type) String() string {
	return "{" + t.Tag().String() + ":" + t.Value() + "}"
}

func Numeric(t Type) bool {
	if t == Char || t == Int || t == Float {
		return true
	} else {
		return false
	}
}

func Max(t1, t2 Type) Type {
	if !Numeric(t1) || !Numeric(t2) {
		return NullType
	} else if t1 == Float || t2 == Float {
		return Float
	} else if t1 == Int || t2 == Int {
		return Int
	} else {
		return Char
	}
}

type Character struct {
	tag   Tag
	value rune
}

func NewChar(tag Tag, value rune) Character {
	return Character{
		tag:   tag,
		value: value,
	}
}

func (c Character) Tag() Tag {
	return c.tag
}

func (c Character) Value() string {
	return string(c.value)
}

func (c Character) String() string {
	return "{" + c.Tag().String() + ":" + c.Value() + "}"
}
