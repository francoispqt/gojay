package main

import (
	"errors"
	"go/ast"
)

func (g *Gen) arrGenUnmarshal(n string, s *ast.ArrayType) error {
	err := arrUnmarshalTpl["def"].tpl.Execute(g.b, struct {
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
		err := g.arrGenUnmarshalIdent(t, false)
		if err != nil {
			return err
		}
	case *ast.StarExpr:
		switch ptrExp := t.X.(type) {
		case *ast.Ident:
			err := g.arrGenUnmarshalIdent(ptrExp, true)
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
	return err
}

func (g *Gen) arrGenUnmarshalIdent(i *ast.Ident, ptr bool) error {
	switch i.String() {
	case "string":
		return g.arrUnmarshalString(ptr)
	case "bool":
		return g.arrUnmarshalBool(ptr)
	case "int":
		return g.arrUnmarshalInt("", ptr)
	case "int64":
		return g.arrUnmarshalInt("64", ptr)
	case "int32":
		return g.arrUnmarshalInt("32", ptr)
	case "int16":
		return g.arrUnmarshalInt("16", ptr)
	case "int8":
		return g.arrUnmarshalInt("8", ptr)
	case "uint64":
		return g.arrUnmarshalUint("64", ptr)
	case "uint32":
		return g.arrUnmarshalUint("32", ptr)
	case "uint16":
		return g.arrUnmarshalUint("16", ptr)
	case "uint8":
		return g.arrUnmarshalUint("8", ptr)
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			return g.arrUnmarshalNonPrim(sp, ptr)
		} else if i.Obj != nil {
			// else check the obj infos
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				return g.arrUnmarshalNonPrim(t, ptr)
			default:
				return errors.New("could not determine what to do with type " + i.String())
			}
		}
		return errors.New("Unknown type")
	}
}

func (g *Gen) arrUnmarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		return g.arrUnmarshalStruct(sp, ptr)
	case *ast.ArrayType:
		return g.arrUnmarshalArr(sp, ptr)
	}
	return nil
}

func (g *Gen) arrUnmarshalString(ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["stringPtr"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrUnmarshalBool(ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrUnmarshalInt(intLen string, ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrUnmarshalUint(intLen string, ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrUnmarshalStruct(st *ast.TypeSpec, ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["struct"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) arrUnmarshalArr(st *ast.TypeSpec, ptr bool) error {
	if ptr {
		err := arrUnmarshalTpl["arrPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	} else {
		err := arrUnmarshalTpl["arr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			return err
		}
	}
	return nil
}
