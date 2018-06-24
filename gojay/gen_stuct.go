package main

import (
	"go/ast"
	"strings"
)

func getStructFieldJSONKey(field *ast.Field) string {
	var keyV string
	if field.Tag != nil {
		keyV = tagKeyName(field.Tag)
	}
	if keyV == "" {
		keyV = strings.ToLower(field.Names[0].String()[:1]) + field.Names[0].String()[1:]
	}
	return keyV
}

func (g *Gen) genStruct(n string, s *ast.StructType) error {
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
