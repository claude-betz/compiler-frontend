/*
	env.go
*/

package symbols

import "compiler-frontend/lexer"

type Env struct {
	table map[string]lexer.Token
	prev  *Env
}

func NewEnv(env *Env) *Env {
	t := make(map[string]lexer.Token)

	return &Env{
		table: t,
		prev:  env,
	}
}

func (e *Env) put(t lexer.Token) {
	e.table[t.Value()] = t
}

func (e *Env) get(key string) lexer.Token {
	for i := e; i != nil; i = i.prev { // loop through all Env tables (scopes)
		val := e.table[key]
		if val != nil { // if we get a non nil value
			return val
		}
	}
	// not found
	return nil
}
