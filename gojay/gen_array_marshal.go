package main

import (
	"errors"
	"fmt"
	"go/ast"
	"log"
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
			return fmt.Errorf("Unknown type %s", n)
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
		g.arrMarshalString(ptr)
	case "bool":
		g.arrMarshalBool(ptr)
	case "int":
		g.arrMarshalInt("", ptr)
	case "int64":
		g.arrMarshalInt("64", ptr)
	case "int32":
		g.arrMarshalInt("32", ptr)
	case "int16":
		g.arrMarshalInt("16", ptr)
	case "int8":
		g.arrMarshalInt("8", ptr)
	case "uint64":
		g.arrMarshalUint("64", ptr)
	case "uint32":
		g.arrMarshalUint("32", ptr)
	case "uint16":
		g.arrMarshalUint("16", ptr)
	case "uint8":
		g.arrMarshalUint("8", ptr)
	case "float64":
		g.arrMarshalFloat("64", ptr)
	case "float32":
		g.arrMarshalFloat("32", ptr)
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
		return fmt.Errorf("Unknown type %s", i.Name)
	}
	return nil
}

func (g *Gen) arrMarshalNonPrim(sp *ast.TypeSpec, ptr bool) error {
	switch sp.Type.(type) {
	case *ast.StructType:
		g.arrMarshalStruct(sp, ptr)
	case *ast.ArrayType:
		g.arrMarshalArr(sp, ptr)
	}
	return nil
}

func (g *Gen) arrMarshalString(ptr bool) {
	if ptr {
		err := arrMarshalTpl["stringPtr"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["string"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalBool(ptr bool) {
	if ptr {
		err := arrMarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["bool"].tpl.Execute(g.b, struct {
			Ptr string
		}{"&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalInt(intLen string, ptr bool) {
	if ptr {
		err := arrMarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["int"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalFloat(intLen string, ptr bool) {
	if ptr {
		err := arrMarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["float"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalUint(intLen string, ptr bool) {
	if ptr {
		err := arrMarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, ""})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["uint"].tpl.Execute(g.b, struct {
			IntLen string
			Ptr    string
		}{intLen, "&"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalStruct(st *ast.TypeSpec, ptr bool) {
	if ptr {
		err := arrMarshalTpl["structPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["struct"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Gen) arrMarshalArr(st *ast.TypeSpec, ptr bool) {
	if ptr {
		err := arrMarshalTpl["arrPtr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := arrMarshalTpl["arr"].tpl.Execute(g.b, struct {
			StructName string
		}{st.Name.String()})
		if err != nil {
			log.Fatal(err)
		}
	}
}
