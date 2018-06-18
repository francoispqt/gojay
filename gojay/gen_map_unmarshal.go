package main

import (
	"errors"
	"go/ast"
)

func (g *Gen) mapGenNKeys(n string, count int) error {
	err := mapUnmarshalTpl["nKeys"].tpl.Execute(g.b, struct {
		NKeys      int
		StructName string
	}{
		NKeys:      count,
		StructName: n,
	})
	return err
}

func (g *Gen) mapGenUnmarshalObj(n string, s *ast.MapType) error {
	err := mapUnmarshalTpl["def"].tpl.Execute(g.b, struct {
		TypeName string
	}{
		TypeName: n,
	})
	if err != nil {
		return err
	}
	switch t := s.Value.(type) {
	case *ast.Ident:
		var err error
		err = g.mapGenUnmarshalIdent(t, false)
		if err != nil {
			return err
		}
	case *ast.StarExpr:
		switch ptrExp := t.X.(type) {
		case *ast.Ident:
			var err error
			err = g.mapGenUnmarshalIdent(ptrExp, true)
			if err != nil {
				return err
			}
		default:
			return errors.New("Unknown type")
		}
	}
	_, err = g.b.Write(structUnmarshalClose)
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapGenUnmarshalIdent(i *ast.Ident, ptr bool) error {
	switch i.String() {
	case "string":
		var err = g.mapUnmarshalString(ptr)
		if err != nil {
			return err
		}
	case "bool":
		var err = g.mapUnmarshalBool(ptr)
		if err != nil {
			return err
		}
	case "int":
		var err = g.mapUnmarshalInt("", ptr)
		if err != nil {
			return err
		}
	case "int64":
		var err = g.mapUnmarshalInt("64", ptr)
		if err != nil {
			return err
		}
	case "int32":
		var err = g.mapUnmarshalInt("32", ptr)
		if err != nil {
			return err
		}
	case "int16":
		var err = g.mapUnmarshalInt("16", ptr)
		if err != nil {
			return err
		}
	case "int8":
		var err = g.mapUnmarshalInt("8", ptr)
		if err != nil {
			return err
		}
	case "uint64":
		var err = g.mapUnmarshalUint("64", ptr)
		if err != nil {
			return err
		}
	case "uint32":
		var err = g.mapUnmarshalUint("32", ptr)
		if err != nil {
			return err
		}
	case "uint16":
		var err = g.mapUnmarshalUint("16", ptr)
		if err != nil {
			return err
		}
	case "uint8":
		var err = g.mapUnmarshalUint("8", ptr)
		if err != nil {
			return err
		}
	case "float64":
		var err = g.mapUnmarshalFloat("", ptr)
		if err != nil {
			return err
		}
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			err := g.mapUnmarshalNonPrim(sp, ptr)
			if err != nil {
				return err
			}
		} else if i.Obj != nil {
			// else check the obj infos
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				err := g.mapUnmarshalNonPrim(t, ptr)
				if err != nil {
					return err
				}
			default:
				return errors.New("could not determine what to do with type " + i.String())
			}
		} else {
			return errors.New("Unknown type")
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.mapUnmarshalStruct(sp, ptr)
	case *ast.ArrayType:
		return g.mapUnmarshalArr(sp, ptr)
	}
	return errors.New("Unknown type")
}

func (g *Gen) mapUnmarshalString(ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalBool(ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalInt(intLen string, ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalUint(intLen string, ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalFloat(intLen string, ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalStruct(st *ast.TypeSpec, ptr bool) error {
	if ptr {
		err := mapUnmarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err := mapUnmarshalTpl["struct"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) mapUnmarshalArr(st *ast.TypeSpec, ptr bool) error {
	err := mapUnmarshalTpl["arr"].tpl.Execute(g.b, struct {
		TypeName string
	}{st.Name.String()})
	if err != nil {
		return err
	}
	return nil
}
