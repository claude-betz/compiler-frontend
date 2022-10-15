/*
	type.go

	handles Type, and Type conversions
*/

package symbols

import "compiler-frontend/lexer"

type Type struct {
	tag   lexer.Tag
	value string
}

func NewType(tag lexer.Tag, value string) Type {
	return Type{
		tag:   tag,
		value: value,
	}
}

func (t Type) Tag() lexer.Tag {
	return t.tag
}

func (t Type) Value() string {
	return t.value
}

func (t Type) String() string {
	return "{" + t.Tag().String() + ":" + t.Value() + "}"
}
