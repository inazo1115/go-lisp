package golisp

import (
	"regexp"
)

var (
	RString  = regexp.MustCompile(`^"[_a-zA-Z][_a-zA-Z0-9]*"$`)
	RInteger = regexp.MustCompile(`^-?[0-9]+$`)
	RFloat   = regexp.MustCompile(`^-?[0-9]+\.[0-9]+$`)
)

type Parser struct {
	lex *lexer
}

func NewParser(expr string) *Parser {
	return &Parser{NewLexer(expr)}
}

func (parser *Parser) Parse() Sexprs {
	token := parser.lex.next()
	if token == "(" {
		return parser.parseList()
	}
	if RString.MatchString(token) {
		return NewAtom(TString, token)
	}
	if RInteger.MatchString(token) {
		return NewAtom(TInteger, token)
	}
	if RFloat.MatchString(token) {
		return NewAtom(TFloat, token)
	}
	return NewAtom(TSymbol, token)
}

func (parser *Parser) parseList() Sexprs {
	if parser.lex.peek() == ")" {
		return NewNil()
	}
	return NewCell(parser.Parse(), parser.parseList())
}
