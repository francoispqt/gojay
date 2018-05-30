package main

import (
	"errors"
	"go/ast"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var structUnmarshalDefTpl *template.Template
var structUnmarshalCaseTpl *template.Template
var structUnmarshalStringTpl *template.Template
var structUnmarshalIntTpl *template.Template
var structUnmarshalUintTpl *template.Template
var structUnmarshalBoolTpl *template.Template
var structUnmarshalStructTpl *template.Template

var structNKeysTpl *template.Template

var nKeysMethod = `
// NKeys returns the number of keys to unmarshal
func (v *{{.StructName}}) NKeys() int { return {{.NKeys}} }
`

var structUnmarshalSwitchOpen = []byte("\tswitch k {\n")
var structUnmarshalClose = []byte("\treturn nil\n}\n")

func init() {
	t, err := template.New("structUnmarshalDef").
		Parse("\n// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject" +
			"\nfunc (v *{{.StructName}}) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {\n",
		)
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalDefTpl = t

	t, err = template.New("structUnmarshalCase").
		Parse("\tcase \"{{.Key}}\":\n")
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalCaseTpl = t

	t, err = template.New("structUnmarshalCaseString").
		Parse("\t\treturn dec.String(&v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalStringTpl = t

	t, err = template.New("structUnmarshalCaseInt").
		Parse("\t\treturn dec.Int{{.IntLen}}(&v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalIntTpl = t

	t, err = template.New("structUnmarshalCaseUint").
		Parse("\t\treturn dec.Uint{{.IntLen}}(&v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalUintTpl = t

	t, err = template.New("structUnmarshalCaseBool").
		Parse("\t\treturn dec.Bool(&v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalBoolTpl = t

	t, err = template.New("structUnmarshalCaseStruct").
		Parse(
			`		if v.{{.Field}} == nil {
			v.{{.Field}} = &{{.StructName}}{}
		}
		dec.Object(v.{{.Field}})
`)
	if err != nil {
		log.Fatal(err)
	}
	structUnmarshalStructTpl = t

	t, err = template.New("structUnmarshalNKeys").
		Parse(nKeysMethod)
	if err != nil {
		log.Fatal(err)
	}
	structNKeysTpl = t

}

func (v *vis) structGenNKeys(f *os.File, n string, count int) error {
	err := structNKeysTpl.Execute(f, struct {
		NKeys      int
		StructName string
	}{
		NKeys:      count,
		StructName: n,
	})
	return err
}

func (v *vis) structGenUnmarshalObj(f *os.File, n string, s *ast.StructType) (int, error) {
	err := structUnmarshalDefTpl.Execute(f, struct {
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
		f.Write(structUnmarshalSwitchOpen)

		// TODO:  check tags
		for _, field := range s.Fields.List {
			switch t := field.Type.(type) {
			case *ast.Ident:
				var err error
				keys, err = v.structGenUnmarshalIdent(f, field, t, keys)
				if err != nil {
					return 0, err
				}
			case *ast.StarExpr:
				switch ptrExp := t.X.(type) {
				case *ast.Ident:
					var err error
					keys, err = v.structGenUnmarshalIdent(f, field, ptrExp, keys)
					if err != nil {
						return 0, err
					}
				default:
					spew.Println(reflect.TypeOf(ptrExp))
					spew.Println(ptrExp)
				}
			default:
				spew.Println(t)
			}
		}
		// close  switch statement
		f.Write([]byte("\t}\n"))
	}
	_, err = f.Write(structUnmarshalClose)
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (v *vis) structGenUnmarshalIdent(f *os.File, field *ast.Field, i *ast.Ident, keys int) (int, error) {
	switch i.String() {
	case "string":
		var err = v.structUnmarshalString(f, field)
		if err != nil {
			return 0, err
		}
		keys++
	case "bool":
		var err = v.structUnmarshalBool(f, field)
		if err != nil {
			return 0, err
		}
		keys++
	case "int":
		var err = v.structUnmarshalInt(f, field, "")
		if err != nil {
			return 0, err
		}
		keys++
	case "int64":
		var err = v.structUnmarshalInt(f, field, "64")
		if err != nil {
			return 0, err
		}
		keys++
	case "int32":
		var err = v.structUnmarshalInt(f, field, "32")
		if err != nil {
			return 0, err
		}
		keys++
	case "int16":
		var err = v.structUnmarshalInt(f, field, "16")
		if err != nil {
			return 0, err
		}
		keys++
	case "int8":
		var err = v.structUnmarshalInt(f, field, "8")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint64":
		var err = v.structUnmarshalUint(f, field, "64")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint32":
		var err = v.structUnmarshalUint(f, field, "32")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint16":
		var err = v.structUnmarshalUint(f, field, "16")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint8":
		var err = v.structUnmarshalUint(f, field, "8")
		if err != nil {
			return 0, err
		}
		keys++
	default:
		switch t := i.Obj.Decl.(type) {
		case *ast.TypeSpec:
			var err = v.structUnmarshalStruct(f, field, t)
			if err != nil {
				return 0, err
			}
			keys++
		default:
			return 0, errors.New("could not determine what to do with type " + i.String())
		}
	}
	return keys, nil
}

func (v *vis) structUnmarshalString(f *os.File, field *ast.Field) error {
	key := field.Names[0].String()
	err := structUnmarshalCaseTpl.Execute(f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalStringTpl.Execute(f, struct {
		Field string
	}{key})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structUnmarshalBool(f *os.File, field *ast.Field) error {
	key := field.Names[0].String()
	err := structUnmarshalCaseTpl.Execute(f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalBoolTpl.Execute(f, struct {
		Field string
	}{key})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structUnmarshalInt(f *os.File, field *ast.Field, intLen string) error {
	key := field.Names[0].String()
	err := structUnmarshalCaseTpl.Execute(f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalIntTpl.Execute(f, struct {
		Field  string
		IntLen string
	}{key, intLen})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structUnmarshalUint(f *os.File, field *ast.Field, intLen string) error {
	key := field.Names[0].String()
	err := structUnmarshalCaseTpl.Execute(f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalUintTpl.Execute(f, struct {
		Field  string
		IntLen string
	}{key, intLen})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structUnmarshalStruct(f *os.File, field *ast.Field, st *ast.TypeSpec) error {
	key := field.Names[0].String()
	err := structUnmarshalCaseTpl.Execute(f, struct {
		Key string
	}{strings.ToLower(key)})
	if err != nil {
		return err
	}
	err = structUnmarshalStructTpl.Execute(f, struct {
		Field      string
		StructName string
	}{key, st.Name.String()})
	if err != nil {
		return err
	}
	return nil
}
