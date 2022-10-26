package inter

import (
	"compiler-frontend/lexer"
	"fmt"
)

var (
	labels = 0 // global variable for labels used in jumps
)

type Node interface {
	Token() lexer.Token // token
	Gen() string        // generate three address code
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

func NewLabel() int {
	// increment global labels
	labels++

	return labels
}

func EmitLabel(i int) string {
	return fmt.Sprintf("L%d", i)
}

func Emit(s string) string {
	return fmt.Sprintf("\t%s\n", s)
}

func printError(s string, line int) {
	fmt.Printf("[ERROR] near line: %d: %s\n", line, s)
}

func LValue(expr Expr) Expr {
	tag := expr.Token().Tag()
	if tag == lexer.ID {
		return expr
	} else if tag == lexer.ACCESS {
		access := expr.(Access)
		return NewAccess(access.id, access.expr)
	} else {
		// error
		return nil
	}
}

func RValue(expr Expr) Expr {
	tag := expr.Token().Tag()
	if tag == lexer.ID || tag == lexer.NUM {
		return expr
	} else if lexer.BoolMap[tag.String()] == true {
		switch tag {
		case lexer.OR:
			or := expr.(Or)

			t := NewTemp()
			s := or.Gen()
			fmt.Printf("\t%s = %s\n", t.toString(), s)

			return t
		case lexer.AND:
			and := expr.(And)

			t := NewTemp()
			s := and.Gen()
			fmt.Printf("\t%s = %s\n", t.toString(), s)

			return t
		}
	} else if lexer.EqMap[tag.String()] == true {
		equality := expr.(Equality)

		t := NewTemp()
		s := equality.Gen()
		fmt.Printf("\t%s = %s\n", t.toString(), s)

		return t
	} else if lexer.RelMap[tag.String()] == true {
		relation := expr.(Rel)

		t := NewTemp()
		s := relation.Gen()
		fmt.Printf("\t%s = %s\n", t.toString(), s)

		return t
	} else if lexer.ExprMap[tag.String()] == true {
		expression := expr.(Expression)

		t := NewTemp()
		s := expression.Gen()
		fmt.Printf("\t%s = %s\n", t.toString(), s)

		return t
	} else if tag == lexer.ACCESS {
		access := expr.(Access)

		t := NewTemp()
		s := access.Gen()
		fmt.Printf("\t%s = %s\n", t.toString(), s)

		return t
	} else if tag == lexer.ASSIGN {
		assign := expr.(Assign)

		s := assign.Gen()
		fmt.Printf("%s\n", s)
		return RValue(assign.expr)
	}
	return nil
}
