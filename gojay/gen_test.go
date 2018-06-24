package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func MakeGenFromReader(input io.Reader) (*Gen, error) {
	g := NewGen("", []string{})
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	v := newVisitor(g, g.pkg)
	ast.Walk(v, f)
	if err != nil {
		return nil, err
	}
	g.vis = v
	return g, nil
}
