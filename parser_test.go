package golisp

import (
	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{"(x y z)", "(x . (y . (z . NIL)))"},
		{"(* 1 2 (+ 3 4))", "(* . (1 . (2 . ((+ . (3 . (4 . NIL))) . NIL))))"},
	}

	for _, test := range tests {
		parser := NewParser(test.in)
		if parser.Parse().String() != test.out {
			t.Errorf("error")
		}
	}
}
