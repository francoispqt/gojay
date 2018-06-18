package main

import (
	"errors"
	"go/ast"
)

func (g *Gen) mapGenIsNil(n string) error {
	err := mapMarshalTpl["isNil"].tpl.Execute(g.b, struct {
		StructName string
	}{
		StructName: n,
	})
	return err
}

func (g *Gen) mapGenMarshalObj(n string, s *ast.MapType) error {
	err := mapMarshalTpl["def"].tpl.Execute(g.b, struct {
		StructName string
	}{
		StructName: n,
	})
	if err != nil {
		return err
	}
	switch t := s.Value.(type) {
	case *ast.Ident:
		var err error
		err = g.mapGenMarshalIdent(t, false)
		if err != nil {
			return err
		}
	case *ast.StarExpr:
		switch ptrExp := t.X.(type) {
		case *ast.Ident:
			var err error
			err = g.mapGenMarshalIdent(ptrExp, true)
			if err != nil {
				return err
			}
		default:
			return errors.New("Unknown type")
		}
	}
	_, err = g.b.Write([]byte("}\n"))
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapGenMarshalIdent(i *ast.Ident, ptr bool) error {
	switch i.String() {
	case "string":
		var err = g.mapMarshalString(ptr)
		if err != nil {
			return err
		}
	case "bool":
		var err = g.mapMarshalBool(ptr)
		if err != nil {
			return err
		}
	case "int":
		var err = g.mapMarshalInt("", ptr)
		if err != nil {
			return err
		}
	case "int64":
		var err = g.mapMarshalInt("64", ptr)
		if err != nil {
			return err
		}
	case "int32":
		var err = g.mapMarshalInt("32", ptr)
		if err != nil {
			return err
		}
	case "int16":
		var err = g.mapMarshalInt("16", ptr)
		if err != nil {
			return err
		}
	case "int8":
		var err = g.mapMarshalInt("8", ptr)
		if err != nil {
			return err
		}
	case "uint64":
		var err = g.mapMarshalUint("64", ptr)
		if err != nil {
			return err
		}
	case "uint32":
		var err = g.mapMarshalUint("32", ptr)
		if err != nil {
			return err
		}
	case "uint16":
		var err = g.mapMarshalUint("16", ptr)
		if err != nil {
			return err
		}
	case "uint8":
		var err = g.mapMarshalUint("8", ptr)
		if err != nil {
			return err
		}
	case "float64":
		var err = g.mapMarshalFloat("", ptr)
		if err != nil {
			return err
		}
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			err := g.mapMarshalNonPrim(sp, ptr)
			if err != nil {
				return err
			}

		} else if i.Obj != nil {
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				var err = g.mapMarshalNonPrim(t, ptr)
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

func (g *Gen) mapMarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.mapMarshalStruct(sp, ptr)
	case *ast.ArrayType:
		return g.mapMarshalArr(sp, ptr)
	}
	return nil
}

func (g *Gen) mapMarshalString(ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["string"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalBool(ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["bool"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalInt(intLen string, ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["int"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalUint(intLen string, ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["uint"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalFloat(intLen string, ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["float"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalStruct(st *ast.TypeSpec, ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = mapMarshalTpl["struct"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) mapMarshalArr(st *ast.TypeSpec, ptr bool) error {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = mapMarshalTpl["arr"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		return err
	}
	return nil
}
