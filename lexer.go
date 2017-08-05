package golisp

type lexer struct {
	expr   string
	carsor int
}

func NewLexer(expr string) *lexer {
	return &lexer{expr, 0}
}

func (lex *lexer) peek() string {
	return string(lex.expr[lex.carsor])
}

func (lex *lexer) next() string {
	lex.skip()
	if lex.peek() == "(" {
		ret := lex.peek()
		lex.carsor++
		return ret
	}
	if lex.peek() == ")" {
		ret := lex.peek()
		lex.carsor++
		return ret
	}
	from := lex.carsor
	for {
		lex.carsor++
		s := lex.peek()
		if s == "(" || s == ")" || s == " " || s == "\n" {
			break
		}
	}
	return string(lex.expr[from:lex.carsor])
}

func (lex *lexer) skip() {
	for {
		if lex.peek() != " " && lex.peek() != "\n" {
			break
		}
		lex.carsor++
	}
}
