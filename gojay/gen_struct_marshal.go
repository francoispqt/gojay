package main

import (
	"errors"
	"go/ast"
	"strings"
)

func (g *gen) structGenIsNil(n string) error {
	err := structMarshalTpl["isNil"].tpl.Execute(g.f, struct {
		StructName string
	}{
		StructName: n,
	})
	return err
}

func (g *gen) structGenMarshalObj(n string, s *ast.StructType) (int, error) {
	err := structMarshalTpl["def"].tpl.Execute(g.f, struct {
		StructName string
	}{
		StructName: n,
	})
	if err != nil {
		return 0, err
	}
	keys := 0
	if len(s.Fields.List) > 0 {
		// TODO:  check tags
		for _, field := range s.Fields.List {
			switch t := field.Type.(type) {
			case *ast.Ident:
				var err error
				keys, err = g.structGenMarshalIdent(field, t, keys, false)
				if err != nil {
					return 0, err
				}
			case *ast.StarExpr:
				switch ptrExp := t.X.(type) {
				case *ast.Ident:
					var err error
					keys, err = g.structGenMarshalIdent(field, ptrExp, keys, true)
					if err != nil {
						return 0, err
					}
				default:
					return 0, errors.New("Unknown type")
				}
			}
		}
	}
	_, err = g.f.Write([]byte("}\n"))
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (g *gen) structGenMarshalIdent(field *ast.Field, i *ast.Ident, keys int, ptr bool) (int, error) {
	switch i.String() {
	case "string":
		var err = g.structMarshalString(field, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "bool":
		var err = g.structMarshalBool(field, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int":
		var err = g.structMarshalInt(field, "", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int64":
		var err = g.structMarshalInt(field, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int32":
		var err = g.structMarshalInt(field, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int16":
		var err = g.structMarshalInt(field, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int8":
		var err = g.structMarshalInt(field, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint64":
		var err = g.structMarshalUint(field, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint32":
		var err = g.structMarshalUint(field, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint16":
		var err = g.structMarshalUint(field, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint8":
		var err = g.structMarshalUint(field, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	default:
		// if ident is already in our spec list
		if sp, ok := g.vis.specs[i.Name]; ok {
			err := g.structMarshalNonPrim(field, sp, ptr)
			if err != nil {
				return 0, err
			}
			keys++
		} else if i.Obj != nil {
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				var err = g.structMarshalNonPrim(field, t, ptr)
				if err != nil {
					return 0, err
				}
				keys++
			default:
				return 0, errors.New("could not determine what to do with type " + i.String())
			}
		} else {
			return 0, errors.New("Unknown type")
		}
	}
	return keys, nil
}

func (g *gen) structMarshalNonPrim(field *ast.Field, sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.structMarshalStruct(field, sp, ptr)
	case *ast.ArrayType:
		return g.structMarshalArr(field, sp, ptr)
	}
	return nil
}

func (g *gen) structMarshalString(field *ast.Field, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["string"].tpl.Execute(g.f, struct {
		Field string
		Key   string
		Ptr   string
	}{key, strings.ToLower(key), ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) structMarshalBool(field *ast.Field, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["bool"].tpl.Execute(g.f, struct {
		Field string
		Key   string
		Ptr   string
	}{key, strings.ToLower(key), ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) structMarshalInt(field *ast.Field, intLen string, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["int"].tpl.Execute(g.f, struct {
		Field  string
		IntLen string
		Key    string
		Ptr    string
	}{key, intLen, strings.ToLower(key), ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) structMarshalUint(field *ast.Field, intLen string, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["uint"].tpl.Execute(g.f, struct {
		Field  string
		IntLen string
		Key    string
		Ptr    string
	}{key, intLen, strings.ToLower(key), ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) structMarshalStruct(field *ast.Field, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = structMarshalTpl["struct"].tpl.Execute(g.f, struct {
		Key   string
		Field string
		Ptr   string
	}{strings.ToLower(key), key, ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) structMarshalArr(field *ast.Field, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = structMarshalTpl["arr"].tpl.Execute(g.f, struct {
		Key   string
		Field string
		Ptr   string
	}{strings.ToLower(key), key, ptrStr})
	if err != nil {
		return err
	}
	return nil
}
