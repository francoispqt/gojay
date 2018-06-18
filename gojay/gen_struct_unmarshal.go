package main

import (
	"errors"
	"go/ast"
)

var structUnmarshalSwitchOpen = []byte("\tswitch k {\n")
var structUnmarshalClose = []byte("\treturn nil\n}\n")

func (g *Gen) structGenNKeys(n string, count int) error {
	err := structUnmarshalTpl["nKeys"].tpl.Execute(g.b, struct {
		NKeys      int
		StructName string
	}{
		NKeys:      count,
		StructName: n,
	})
	return err
}

func (g *Gen) structGenUnmarshalObj(n string, s *ast.StructType) (int, error) {
	err := structUnmarshalTpl["def"].tpl.Execute(g.b, struct {
		StructName string
	}{
		StructName: n,
	})
	if err != nil {
		return 0, err
	}
	keys := 0
	if len(s.Fields.List) > 0 {
		// open  switch statement
		g.b.Write(structUnmarshalSwitchOpen)
		// TODO:  check tags
		// check type of field
		// add accordingly
		for _, field := range s.Fields.List {
			// check if has hide tag
			if field.Tag != nil && hasTagUnmarshalHide(field.Tag) {
				continue
			}
			switch t := field.Type.(type) {
			case *ast.Ident:
				var err error
				keys, err = g.structGenUnmarshalIdent(field, t, keys, false)
				if err != nil {
					return 0, err
				}
			case *ast.StarExpr:
				switch ptrExp := t.X.(type) {
				case *ast.Ident:
					var err error
					keys, err = g.structGenUnmarshalIdent(field, ptrExp, keys, true)
					if err != nil {
						return 0, err
					}
				default:
					return 0, errors.New("Unknown type")
				}
			}
		}
		// close  switch statement
		g.b.Write([]byte("\t}\n"))
	}
	_, err = g.b.Write(structUnmarshalClose)
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (g *Gen) structGenUnmarshalIdent(field *ast.Field, i *ast.Ident, keys int, ptr bool) (int, error) {
	var keyV = getStructFieldJSONKey(field)

	switch i.String() {
	case "string":
		var err = g.structUnmarshalString(field, keyV, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "bool":
		var err = g.structUnmarshalBool(field, keyV, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int":
		var err = g.structUnmarshalInt(field, keyV, "", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int64":
		var err = g.structUnmarshalInt(field, keyV, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int32":
		var err = g.structUnmarshalInt(field, keyV, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int16":
		var err = g.structUnmarshalInt(field, keyV, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int8":
		var err = g.structUnmarshalInt(field, keyV, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint64":
		var err = g.structUnmarshalUint(field, keyV, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint32":
		var err = g.structUnmarshalUint(field, keyV, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint16":
		var err = g.structUnmarshalUint(field, keyV, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint8":
		var err = g.structUnmarshalUint(field, keyV, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "float64":
		var err = g.structUnmarshalFloat(field, keyV, "", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			err := g.structUnmarshalNonPrim(field, keyV, sp, ptr)
			if err != nil {
				return 0, err
			}
			keys++
		} else if i.Obj != nil {
			// else check the obj infos
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				err := g.structUnmarshalNonPrim(field, keyV, t, ptr)
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

func (g *Gen) structUnmarshalNonPrim(field *ast.Field, keyV string, sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.structUnmarshalStruct(field, keyV, sp, ptr)
	case *ast.ArrayType:
		return g.structUnmarshalArr(field, keyV, sp, ptr)
	}
	return errors.New("Unknown type")
}

func (g *Gen) structUnmarshalString(field *ast.Field, keyV string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalBool(field *ast.Field, keyV string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalInt(field *ast.Field, keyV string, intLen string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalUint(field *ast.Field, keyV string, intLen string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalFloat(field *ast.Field, keyV string, intLen string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalStruct(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["struct"].tpl.Execute(g.b, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) structUnmarshalArr(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		return err
	}
	err = structUnmarshalTpl["arr"].tpl.Execute(g.b, struct {
		Field    string
		TypeName string
	}{key, st.Name.String()})
	if err != nil {
		return err
	}
	return nil
}
