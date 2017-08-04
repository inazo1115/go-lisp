package golisp

import (
	"testing"
)

func TestPeek(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{"01234", "0"},
	}

	for _, test := range tests {
		lex := NewLexer(test.in)
		if lex.peek() != test.out {
			t.Errorf("error")
		}
	}
}

func TestNext(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{"(* 1 2)", "("},
		{"* 1 2)", "*"},
		{" 1 2)", "1"},
		{"1 2)", "1"},
		{"2)", "2"},
		{")", ")"},
	}

	for _, test := range tests {
		lex := NewLexer(test.in)
		if lex.next() != test.out {
			t.Errorf("error")
		}
	}
}
