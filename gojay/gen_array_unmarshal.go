package main

import (
	"errors"
	"fmt"
	"go/ast"
	"log"
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
			return fmt.Errorf("Unknown type %s", n)
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
		g.arrUnmarshalString(ptr)
	case "bool":
		g.arrUnmarshalBool(ptr)
	case "int":
		g.arrUnmarshalInt("", ptr)
	case "int64":
		g.arrUnmarshalInt("64", ptr)
	case "int32":
		g.arrUnmarshalInt("32", ptr)
	case "int16":
		g.arrUnmarshalInt("16", ptr)
	case "int8":
		g.arrUnmarshalInt("8", ptr)
	case "uint64":
		g.arrUnmarshalUint("64", ptr)
	case "uint32":
		g.arrUnmarshalUint("32", ptr)
	case "uint16":
		g.arrUnmarshalUint("16", ptr)
	case "uint8":
		g.arrUnmarshalUint("8", ptr)
	case "float64":
		g.arrUnmarshalFloat("64", ptr)
	case "float32":
		g.arrUnmarshalFloat("32", ptr)
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
		return fmt.Errorf("Unknown type %s", i.Name)
	}
	return nil
}

func (g *Gen) arrUnmarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		g.arrUnmarshalStruct(sp, ptr)
	case *ast.ArrayType:
		g.arrUnmarshalArr(sp, ptr)
	}
	return nil
}

func (g *Gen) arrUnmarshalString(ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["stringPtr"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalBool(ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalInt(intLen string, ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalUint(intLen string, ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalFloat(intLen string, ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalStruct(st *ast.TypeSpec, ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["struct"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrUnmarshalArr(st *ast.TypeSpec, ptr bool) {
	if ptr {
		err := arrUnmarshalTpl["arrPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrUnmarshalTpl["arr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}
