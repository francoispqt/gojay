package main

import (
	"go/ast"
)

func (g *Gen) genArray(n string, s *ast.ArrayType) error {
	err := g.arrGenUnmarshal(n, s)
	if err != nil {
		return err
	}
	err = g.arrGenMarshal(n, s)
	if err != nil {
		return err
	}
	return g.arrGenIsNil(n)
}
