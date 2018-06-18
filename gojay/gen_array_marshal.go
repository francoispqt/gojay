package main

import (
	"errors"
	"go/ast"
)

func init() {}

func (g *Gen) arrGenIsNil(n string) error {
	err := arrMarshalTpl["isNil"].tpl.Execute(g.b, struct {
		TypeName string
	}{
		TypeName: n,
	})
	return err
}

func (g *Gen) arrGenMarshal(n string, s *ast.ArrayType) error {
	err := arrMarshalTpl["def"].tpl.Execute(g.b, struct {
		TypeName string
	}{
		TypeName: n,
	})
	if err != nil {
		return err
	}
	// determine type of element in array
	switch t := s.Elt.(type) {
	case *ast.Ident:
		err := g.arrGenMarshalIdent(t, false)
		if err != nil {
			return err
		}
	case *ast.StarExpr:
		switch ptrExp := t.X.(type) {
		case *ast.Ident:
			err := g.arrGenMarshalIdent(ptrExp, true)
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
	return err
}

func (g *Gen) arrGenMarshalIdent(i *ast.Ident, ptr bool) error {
	switch i.String() {
	case "string":
		return g.arrMarshalString(ptr)
	case "bool":
		return g.arrMarshalBool(ptr)
	case "int":
		return g.arrMarshalInt("", ptr)
	case "int64":
		return g.arrMarshalInt("64", ptr)
	case "int32":
		return g.arrMarshalInt("32", ptr)
	case "int16":
		return g.arrMarshalInt("16", ptr)
	case "int8":
		return g.arrMarshalInt("8", ptr)
	case "uint64":
		return g.arrMarshalUint("64", ptr)
	case "uint32":
		return g.arrMarshalUint("32", ptr)
	case "uint16":
		return g.arrMarshalUint("16", ptr)
	case "uint8":
		return g.arrMarshalUint("8", ptr)
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			return g.arrMarshalNonPrim(sp, ptr)
		} else if i.Obj != nil {
			// else check the obj infos
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				return g.arrMarshalNonPrim(t, ptr)
			default:
				return errors.New("could not determine what to do with type " + i.String())
			}
		}
		return errors.New("Unknown type")
	}
}

func (g *Gen) arrMarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.arrMarshalStruct(sp, ptr)
	case *ast.ArrayType:
		return g.arrMarshalArr(sp, ptr)
	}
	return nil
}

func (g *Gen) arrMarshalString(ptr bool) error {
	if ptr {
		err := arrMarshalTpl["stringPtr"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrMarshalBool(ptr bool) error {
	if ptr {
		err := arrMarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrMarshalInt(intLen string, ptr bool) error {
	if ptr {
		err := arrMarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrMarshalUint(intLen string, ptr bool) error {
	if ptr {
		err := arrMarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrMarshalStruct(st *ast.TypeSpec, ptr bool) error {
	if ptr {
		err := arrMarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["struct"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrMarshalArr(st *ast.TypeSpec, ptr bool) error {
	if ptr {
		err := arrMarshalTpl["arrPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err := arrMarshalTpl["arr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}
