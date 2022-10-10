/*
	lexer.go

	implementation of a lexical analyser to tokenize the source code (subset of java) and return a
	stream of tokens for consumption by the subsequent stages of the compiler front-end.
*/

package lexer

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
)

const (
	// valid characters
	newline    rune = 0x0A
	tab        rune = 0x09
	whitespace rune = 0x20
)

type Lexer struct {
	source io.RuneScanner
	words  map[string]Token
	peek   rune
	line   int
}

func NewLexer(source io.RuneScanner) (*Lexer, error) {
	words := make(map[string]Token)

	l := &Lexer{
		source: source,
		words:  words,
		peek:   whitespace,
		line:   0,
	}

	// reserve statements
	l.put(NewWord(FOR, FOR.String()))
	l.put(NewWord(WHILE, WHILE.String()))
	l.put(NewWord(IF, IF.String()))
	l.put(NewWord(ELSE, ELSE.String()))
	l.put(NewWord(DO, DO.String()))
	l.put(NewWord(BREAK, BREAK.String()))

	// expressions
	// assignment
	l.put(NewWord(ASSIGN, ASSIGN.String()))

	// boolean ops
	l.put(NewWord(OR, OR.String()))
	l.put(NewWord(AND, AND.String()))

	// equality
	l.put(NewWord(EQUAL_TO, EQUAL_TO.String()))
	l.put(NewWord(NOT_EQUAL_TO, NOT_EQUAL_TO.String()))

	// relational
	l.put(NewWord(LESS_THAN, LESS_THAN.String()))
	l.put(NewWord(LESS_THAN_EQUAL_TO, LESS_THAN_EQUAL_TO.String()))
	l.put(NewWord(GREATER_THAN, GREATER_THAN.String()))
	l.put(NewWord(GREATER_THAN_EQUAL_TO, GREATER_THAN_EQUAL_TO.String()))

	// expr
	l.put(NewWord(ADD, ADD.String()))
	l.put(NewWord(SUBTRACT, ADD.String()))

	// terms
	l.put(NewWord(MULTIPLY, MULTIPLY.String()))
	l.put(NewWord(DIVIDE, DIVIDE.String()))

	// unary
	l.put(NewWord(NOT, NOT.String()))
	l.put(NewWord(MINUS, MINUS.String()))

	// types
	l.put(NewType(PRIMITIVE, "int"))
	l.put(NewType(PRIMITIVE, "bool"))
	l.put(NewType(PRIMITIVE, "char"))
	l.put(NewType(PRIMITIVE, "float"))

	// values
	l.put(NewType(TRUE, "true"))
	l.put(NewType(FALSE, "false"))

	return l, nil
}

func (l *Lexer) put(t Token) {
	l.words[t.Value()] = t
}

func (l *Lexer) get(value string) Token {
	return l.words[value]
}

func (l *Lexer) advancePeek() error {
	r, _, err := l.source.ReadRune()
	l.peek = r

	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("[lexer] reached end of source: %v", err)
		}
		return fmt.Errorf("[lexer] failed to read character: %v", err)
	}
	return nil
}

func (l *Lexer) Scan() Token {
	l.skipWhitespace()

	token := l.readCompositeOperators()
	if token != nil {
		return token
	}

	token = l.readNumber()
	if token != nil {
		return token
	}

	token = l.readWord()
	if token != nil {
		return token
	}

	token = l.readCharacters()

	return token
}

func (l *Lexer) skipWhitespace() {
	for {
		p := l.peek

		// skip spaces
		if unicode.IsSpace(p) {
			// count new lines
			if newline == p {
				l.line++
			}

			// read next character
			l.advancePeek()
		} else {
			// when we have a non whitespace character break out of loop
			break
		}
	}
}

func (l *Lexer) readNumber() Token {
	p := l.peek

	// hold number
	v := 0

	// first must have digit
	if unicode.IsDigit(p) {
		for {
			// if digit
			if unicode.IsDigit(p) {
				v = v*10 + runeToInt(p)

				// read next character
				l.advancePeek()
				p = l.peek
			} else {
				break
			}
		}
	} else {
		return nil
	}

	// return Int Token
	return Num{
		tag:   NUM,
		value: v,
	}
}

func (l *Lexer) readWord() Token {
	p := l.peek

	// buffer to read word
	buf := bytes.Buffer{}
	buf.Grow(20)

	// first character needs to be letter
	if unicode.IsLetter(p) {
		for {
			// subsequent characters can be letters or numbers
			if unicode.IsLetter(p) || unicode.IsDigit(p) {
				// add to buffer
				buf.WriteRune(p)

				// read next character
				l.advancePeek()
				p = l.peek
			} else {
				// when we are no longer reading character
				break
			}
		}

		// check for lexeme in list of words
		lexeme := buf.String()
		token := l.words[lexeme]

		if token != nil {
			// token already exists
			return token
		} else {
			// write token to words
			word := Word{
				tag:    ID,
				lexeme: lexeme,
			}
			l.words[lexeme] = word

			// return
			return word
		}
	} else {
		// return if we are not dealing with a letter
		return nil
	}
}

func (l *Lexer) readCompositeOperators() Token {

	switch l.peek {
	case '|':
		l.advancePeek()
		if l.peek == '|' {
			l.advancePeek()
			return l.get(OR.String())
		} else {
			return NewChar(OR, '|')
		}
	case '&':
		l.advancePeek()
		if l.peek == '&' {
			l.advancePeek()
			return l.get("&&")
		} else {
			return NewChar(AND, '&')
		}
	case '=':
		l.advancePeek()
		if l.peek == '=' {
			l.advancePeek()
			return l.get(EQUAL_TO.String())
		} else {
			return l.get(ASSIGN.String())
		}
	case '!':
		l.advancePeek()
		if l.peek == '=' {
			l.advancePeek()
			return l.get(NOT_EQUAL_TO.String())
		} else {
			return l.get(NOT.String())
		}
	case '<':
		l.advancePeek()
		if l.peek == '=' {
			l.advancePeek()
			return l.get(LESS_THAN_EQUAL_TO.String())
		} else {
			return l.get(LESS_THAN.String())
		}
	case '>':
		l.advancePeek()
		if l.peek == '=' {
			l.advancePeek()
			return l.get(GREATER_THAN_EQUAL_TO.String())
		} else {
			return l.get(GREATER_THAN.String())
		}
	default:
		return nil
	}
}

func (l *Lexer) readCharacters() Token {
	char := NewChar(ID, l.peek)
	l.peek = whitespace
	return char
}

func runeToInt(r rune) int {
	return int(r - '0')
}
