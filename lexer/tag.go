/*
	tag.go

	file implementing enum type for tags
*/

package lexer

type Tag int64

const (
	Undefined Tag = iota

	// terminals
	NUM
	ID

	// reserved words
	// statements
	FOR
	WHILE
	IF
	ELSE
	DO
	BREAK

	// expressions
	// assignment
	ASSIGN

	// boolean ops
	OR
	AND

	// equality
	EQUAL_TO
	NOT_EQUAL_TO

	// relational
	LESS_THAN
	LESS_THAN_EQUAL_TO
	GREATER_THAN
	GREATER_THAN_EQUAL_TO

	// exprs
	ADD
	SUBTRACT

	// terms
	MULTIPLY
	DIVIDE

	// unary
	NOT
	MINUS

	// basic types
	PRIMITIVE
	BOOL

	// basic values
	TRUE
	FALSE

	// character
	CHARACTER
)

func (t Tag) String() string {
	switch t {
	case NUM:
		return "num"
	case ID:
		return "id"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case DO:
		return "do"
	case BREAK:
		return "break"
	case ASSIGN:
		return "="
	case OR:
		return "||"
	case AND:
		return "&&"
	case EQUAL_TO:
		return "=="
	case NOT_EQUAL_TO:
		return "!="
	case LESS_THAN:
		return "<"
	case LESS_THAN_EQUAL_TO:
		return "<="
	case GREATER_THAN:
		return ">"
	case GREATER_THAN_EQUAL_TO:
		return ">="
	case ADD:
		return "+"
	case SUBTRACT:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case NOT:
		return "!"
	case MINUS:
		return "-"
	case PRIMITIVE:
		return "primitive"
	case CHARACTER:
		return "character"
	}

	return "undefined"
}
