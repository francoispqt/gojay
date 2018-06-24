package main

import (
	"errors"
	"fmt"
	"go/ast"
	"log"
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
			return fmt.Errorf("Unknown type %s", n)
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
		g.mapMarshalString(ptr)
	case "bool":
		g.mapMarshalBool(ptr)
	case "int":
		g.mapMarshalInt("", ptr)
	case "int64":
		g.mapMarshalInt("64", ptr)
	case "int32":
		g.mapMarshalInt("32", ptr)
	case "int16":
		g.mapMarshalInt("16", ptr)
	case "int8":
		g.mapMarshalInt("8", ptr)
	case "uint64":
		g.mapMarshalUint("64", ptr)
	case "uint32":
		g.mapMarshalUint("32", ptr)
	case "uint16":
		g.mapMarshalUint("16", ptr)
	case "uint8":
		g.mapMarshalUint("8", ptr)
	case "float64":
		g.mapMarshalFloat("64", ptr)
	case "float32":
		g.mapMarshalFloat("32", ptr)
	default:
		// if ident is already in our spec list
		if sp, ok := g.genTypes[i.Name]; ok {
			g.mapMarshalNonPrim(sp, ptr)
		} else if i.Obj != nil {
			switch t := i.Obj.Decl.(type) {
			case *ast.TypeSpec:
				g.mapMarshalNonPrim(t, ptr)
			default:
				return errors.New("could not determine what to do with type " + i.String())
			}
		} else {
			return fmt.Errorf("Unknown type %s", i.Name)
		}
	}
	return nil
}

func (g *Gen) mapMarshalNonPrim(sp *ast.TypeSpec, ptr bool) {
	switch sp.Type.(type) {
	case *ast.StructType:
		g.mapMarshalStruct(sp, ptr)
	case *ast.ArrayType:
		g.mapMarshalArr(sp, ptr)
	}
}

func (g *Gen) mapMarshalString(ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["string"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalBool(ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["bool"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalInt(intLen string, ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["int"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalUint(intLen string, ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["uint"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalFloat(intLen string, ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	err := mapMarshalTpl["float"].tpl.Execute(g.b, struct {
		IntLen string
		Ptr    string
	}{intLen, ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalStruct(st *ast.TypeSpec, ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = mapMarshalTpl["struct"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) mapMarshalArr(st *ast.TypeSpec, ptr bool) {
	ptrStr := ""
	if ptr {
		ptrStr = "*"
	}
	var err = mapMarshalTpl["arr"].tpl.Execute(g.b, struct {
		Ptr string
	}{ptrStr})
	if err != nil {
		log.Fatal(err)
	}
}
