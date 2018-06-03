package main

import (
	"errors"
	"go/ast"
	"strings"
)

var structUnmarshalSwitchOpen = []byte("\tswitch k {\n")
var structUnmarshalClose = []byte("\treturn nil\n}\n")

func (g *gen) structGenNKeys(n string, count int) error {
	err := structUnmarshalTpl["nKeys"].tpl.Execute(g.f, struct {
		NKeys      int
		StructName string
	}{
		NKeys:      count,
		StructName: n,
	})
	return err
}

func (g *gen) structGenUnmarshalObj(n string, s *ast.StructType) (int, error) {
	err := structUnmarshalTpl["def"].tpl.Execute(g.f, struct {
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
		g.f.Write(structUnmarshalSwitchOpen)

		// TODO:  check tags
		for _, field := range s.Fields.List {
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
					return 0, errors.New("Unknown type1")
				}
			}
		}
		// close  switch statement
		g.f.Write([]byte("\t}\n"))
	}
	_, err = g.f.Write(structUnmarshalClose)
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (g *gen) structGenUnmarshalIdent(field *ast.Field, i *ast.Ident, keys int, ptr bool) (int, error) {
	switch i.String() {
	case "string":
		var err = g.structUnmarshalString(field, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "bool":
		var err = g.structUnmarshalBool(field, ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int":
		var err = g.structUnmarshalInt(field, "", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int64":
		var err = g.structUnmarshalInt(field, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int32":
		var err = g.structUnmarshalInt(field, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int16":
		var err = g.structUnmarshalInt(field, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "int8":
		var err = g.structUnmarshalInt(field, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint64":
		var err = g.structUnmarshalUint(field, "64", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint32":
		var err = g.structUnmarshalUint(field, "32", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint16":
		var err = g.structUnmarshalUint(field, "16", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	case "uint8":
		var err = g.structUnmarshalUint(field, "8", ptr)
		if err != nil {
			return 0, err
		}
		keys++
	default:
		// if ident is already in our spec list
		if sp, ok := g.vis.specs[i.Name]; ok {
			err := g.structUnmarshalNonPrim(field, sp, ptr)
			if err != nil {
				return 0, err
			}
			keys++
		} else if i.Obj != nil {
			// else check the obj infos
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				err := g.structUnmarshalNonPrim(field, t, ptr)
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

func (g *gen) structUnmarshalNonPrim(field *ast.Field, sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.structUnmarshalStruct(field, sp, ptr)
	case *ast.ArrayType:
		return g.structUnmarshalArr(field, sp, ptr)
	}
	return errors.New("Unknown type")
}

func (g *gen) structUnmarshalString(field *ast.Field, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["string"].tpl.Execute(g.f, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["string"].tpl.Execute(g.f, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *gen) structUnmarshalBool(field *ast.Field, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.f, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.f, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *gen) structUnmarshalInt(field *ast.Field, intLen string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["int"].tpl.Execute(g.f, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["int"].tpl.Execute(g.f, struct {
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

func (g *gen) structUnmarshalUint(field *ast.Field, intLen string, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.f, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.f, struct {
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

func (g *gen) structUnmarshalStruct(field *ast.Field, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	if ptr {
		err = structUnmarshalTpl["struct"].tpl.Execute(g.f, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err = structUnmarshalTpl["struct"].tpl.Execute(g.f, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *gen) structUnmarshalArr(field *ast.Field, st *ast.TypeSpec, ptr bool) error {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalTpl["arr"].tpl.Execute(g.f, struct {
		Field    string
		TypeName string
	}{key, st.Name.String()})
	if err != nil {
		return err
	}
	return nil
}
