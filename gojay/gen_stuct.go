package main

import (
	"go/ast"
	"os"
)

func (v *vis) genStruct(f *os.File, n string, s *ast.StructType) error {
	keys, err := v.structGenUnmarshalObj(f, n, s)
	if err != nil {
		return err
	}
	err = v.structGenNKeys(f, n, keys)

	keys, err = v.structGenMarshalObj(f, n, s)
	if err != nil {
		return err
	}
	err = v.structGenIsNil(f, n)
	return nil
}
