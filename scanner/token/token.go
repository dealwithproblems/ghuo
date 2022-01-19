package token

import "fmt"

type token uint

type tokenType string

type Token struct {
	Type    tokenType
	lexeme  string
	literal interface{}
	line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.Type, t.lexeme, t.literal)
}

const (
	_ token = iota
	_LBRACE
	_RBRACE
	_LPAREN
	RPAREN
)
