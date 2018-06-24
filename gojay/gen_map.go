package main

import "go/ast"

func (g *Gen) genMap(n string, s *ast.MapType) error {
	err := g.mapGenUnmarshalObj(n, s)
	if err != nil {
		return err
	}
	err = g.mapGenNKeys(n, 0)
	if err != nil {
		return err
	}
	err = g.mapGenMarshalObj(n, s)
	if err != nil {
		return err
	}
	return g.mapGenIsNil(n)
}
