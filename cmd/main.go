package main

import (
	"fmt"
	"os"

	lisp "github.com/inazo1115/go-lisp"
)

func main() {
	parser := lisp.NewParser(os.Args[1])
	ast := parser.Parse()
	fmt.Println(ast.String())
}
