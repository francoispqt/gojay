package main

import (
	"fmt"
	"go/ast"
	"log"
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
					return 0, fmt.Errorf("Unknown type %s", n)
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
		g.structUnmarshalString(field, keyV, ptr)
		keys++
	case "bool":
		g.structUnmarshalBool(field, keyV, ptr)
		keys++
	case "int":
		g.structUnmarshalInt(field, keyV, "", ptr)
		keys++
	case "int64":
		g.structUnmarshalInt(field, keyV, "64", ptr)
		keys++
	case "int32":
		g.structUnmarshalInt(field, keyV, "32", ptr)
		keys++
	case "int16":
		g.structUnmarshalInt(field, keyV, "16", ptr)
		keys++
	case "int8":
		g.structUnmarshalInt(field, keyV, "8", ptr)
		keys++
	case "uint64":
		g.structUnmarshalUint(field, keyV, "64", ptr)
		keys++
	case "uint32":
		g.structUnmarshalUint(field, keyV, "32", ptr)
		keys++
	case "uint16":
		g.structUnmarshalUint(field, keyV, "16", ptr)
		keys++
	case "uint8":
		g.structUnmarshalUint(field, keyV, "8", ptr)
		keys++
	case "float64":
		g.structUnmarshalFloat(field, keyV, "64", ptr)
		keys++
	case "float32":
		g.structUnmarshalFloat(field, keyV, "32", ptr)
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
				g.structUnmarshalAny(field, keyV, sp, ptr)
				keys++
			}
		} else {
			g.structUnmarshalAny(field, keyV, sp, ptr)
			keys++
		}
	}
	return keys, nil
}

func (g *Gen) structUnmarshalNonPrim(field *ast.Field, keyV string, sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		g.structUnmarshalStruct(field, keyV, sp, ptr)
		return nil
	case *ast.ArrayType:
		g.structUnmarshalArr(field, keyV, sp, ptr)
		return nil
	default:
		g.structUnmarshalAny(field, keyV, sp, ptr)
		return nil
	}
}

func (g *Gen) structUnmarshalString(field *ast.Field, keyV string, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalBool(field *ast.Field, keyV string, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Field string
			Ptr   string
		}{key, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalInt(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalUint(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalFloat(field *ast.Field, keyV string, intLen string, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			Field  string
			IntLen string
			Ptr    string
		}{key, intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalStruct(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["struct"].tpl.Execute(g.b, struct {
			Field      string
			StructName string
		}{key, st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalArr(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["arrPtr"].tpl.Execute(g.b, struct {
			Field    string
			TypeName string
		}{key, st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["arr"].tpl.Execute(g.b, struct {
			Field    string
			TypeName string
		}{key, st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) structUnmarshalAny(field *ast.Field, keyV string, st *ast.TypeSpec, ptr bool) {
	key := field.Names[0].String()
	err := structUnmarshalTpl["case"].tpl.Execute(g.b, struct {
		Key string
	}{keyV})
	if err != nil {
		log.Fatal(err)
	}
	if ptr {
		err = structUnmarshalTpl["anyPtr"].tpl.Execute(g.b, struct {
			Field string
		}{key})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = structUnmarshalTpl["any"].tpl.Execute(g.b, struct {
			Field string
		}{key})
		if err != nil {
			log.Fatal(err)
		}
	}
}
