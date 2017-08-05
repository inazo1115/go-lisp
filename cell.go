package golisp

import (
	"fmt"
)

type Type int

const (
	TSymbol Type = iota
	TString
	TInteger
	TFloat
)

type Sexprs interface {
	fmt.Stringer
	Eval() Sexprs
}

type Atom struct {
	type_ Type
	val   string
}

type Cell struct {
	left  Sexprs
	right Sexprs
}

type Nil struct {
}

func NewAtom(type_ Type, val string) *Atom {
	return &Atom{type_, val}
}

func NewCell(left, right Sexprs) *Cell {
	return &Cell{left, right}
}

func NewNil() *Nil {
	return &Nil{}
}

func (atom *Atom) Eval() Sexprs {
	panic("not implemented")
}

func (cell *Cell) Eval() Sexprs {
	panic("not implemented")
}

func (nil_ *Nil) Eval() Sexprs {
	panic("not implemented")
}

func (atom *Atom) String() string {
	return atom.val
}

func (cell *Cell) String() string {
	return fmt.Sprintf("(%s . %s)", cell.left.String(), cell.right.String())
}

func (nil_ *Nil) String() string {
	return "NIL"
}
