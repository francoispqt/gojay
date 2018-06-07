package main

import (
	"go/ast"
)

func (g *gen) genStruct(n string, s *ast.StructType) error {
	keys, err := g.structGenUnmarshalObj(n, s)
	if err != nil {
		return err
	}
	err = g.structGenNKeys(n, keys)
	keys, err = g.structGenMarshalObj(n, s)
	if err != nil {
		return err
	}
	return g.structGenIsNil(n)
}
