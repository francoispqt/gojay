package main

import (
	"errors"
	"go/ast"
	"log"
)

func (g *Gen) structGenIsNil(n string) error {
	err := structMarshalTpl["isNil"].tpl.Execute(g.b, struct {
		StructName string
	}{
		StructName: n,
	})
	return err
}

func (g *Gen) structGenMarshalObj(n string, s *ast.StructType) (int, error) {
	err := structMarshalTpl["def"].tpl.Execute(g.b, struct {
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
			// check if has hide tag
			if field.Tag != nil && hasTagMarshalHide(field.Tag) {
				continue
			}
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
	_, err = g.b.Write([]byte("}\n"))
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (g *Gen) structGenMarshalIdent(field *ast.Field, i *ast.Ident, keys int, ptr bool) (int, error) {
	var keyV = getStructFieldJSONKey(field)

	switch i.String() {
	case "string":
		g.structMarshalString(field, keyV, ptr)
		keys++
	case "bool":
		g.structMarshalBool(field, keyV, ptr)
		keys++
	case "int":
		g.structMarshalInt(field, keyV, "", ptr)
		keys++
	case "int64":
		g.structMarshalInt(field, keyV, "64", ptr)
		keys++
	case "int32":
		g.structMarshalInt(field, keyV, "32", ptr)
		keys++
	case "int16":
		g.structMarshalInt(field, keyV, "16", ptr)
		keys++
	case "int8":
		g.structMarshalInt(field, keyV, "8", ptr)
		keys++
	case "uint64":
		g.structMarshalUint(field, keyV, "64", ptr)
		keys++
	case "uint32":
		g.structMarshalUint(field, keyV, "32", ptr)
		keys++
	case "uint16":
		g.structMarshalUint(field, keyV, "16", ptr)
		keys++
	case "uint8":
		g.structMarshalUint(field, keyV, "8", ptr)
		keys++
	case "float64":
		g.structMarshalFloat(field, keyV, "", ptr)
		keys++
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			err := g.structMarshalNonPrim(field, keyV, sp, ptr)
			if err != nil {
				return 0, err
			}
			keys++
		} else if i.Obj != nil {
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				var err = g.structMarshalNonPrim(field, keyV, t, ptr)
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

func (g *Gen) structMarshalNonPrim(field *ast.Field, keyV string, sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		g.structMarshalStruct(field, keyV, sp, ptr)
		return nil
	case *ast.ArrayType:
		g.structMarshalArr(field, keyV, sp, ptr)
		return nil
	}
	return nil
}

func (g *Gen) structMarshalString(field *ast.Field, keyV string, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["string"].tpl.Execute(g.b, struct {
		Field string
		Key   string
		Ptr   string
	}{key, keyV, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalBool(field *ast.Field, keyV string, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["bool"].tpl.Execute(g.b, struct {
		Field string
		Key   string
		Ptr   string
	}{key, keyV, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalInt(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["int"].tpl.Execute(g.b, struct {
		Field  string
		IntLen string
		Key    string
		Ptr    string
	}{key, intLen, keyV, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalUint(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["uint"].tpl.Execute(g.b, struct {
		Field  string
		IntLen string
		Key    string
		Ptr    string
	}{key, intLen, keyV, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalFloat(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := structMarshalTpl["float"].tpl.Execute(g.b, struct {
		Field  string
		IntLen string
		Key    string
		Ptr    string
	}{key, intLen, keyV, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalStruct(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = structMarshalTpl["struct"].tpl.Execute(g.b, struct {
		Key   string
		Field string
		Ptr   string
	}{keyV, key, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) structMarshalArr(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) {
	key := field.Names[0].String()
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = structMarshalTpl["arr"].tpl.Execute(g.b, struct {
		Key   string
		Field string
		Ptr   string
	}{keyV, key, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}
