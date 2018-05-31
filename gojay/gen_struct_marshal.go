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

var structMarshalDefTpl *template.Template
var structMarshalStringTpl *template.Template
var structMarshalIntTpl *template.Template
var structMarshalUintTpl *template.Template
var structMarshalBoolTpl *template.Template
var structMarshalStructTpl *template.Template

var structIsNilTpl *template.Template

var isNilMethod = `
// IsNil returns wether the structure is nil value or not
func (v *{{.StructName}}) IsNil() bool { return v == nil }
`

func init() {
	t, err := template.New("structUnmarshalDef").
		Parse("\n// MarshalJSONObject implements gojay's MarshalerJSONObject" +
			"\nfunc (v *{{.StructName}}) MarshalJSONObject(enc *gojay.Encoder) {\n",
		)
	if err != nil {
		log.Fatal(err)
	}
	structMarshalDefTpl = t

	t, err = template.New("structMarshalCaseString").
		Parse("\tenc.StringKey(\"{{.Key}}\", v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structMarshalStringTpl = t

	t, err = template.New("structMarshalCaseInt").
		Parse("\tenc.Int{{.IntLen}}Key(\"{{.Key}}\", v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structMarshalIntTpl = t

	t, err = template.New("structMarshalCaseUint").
		Parse("\tenc.Uint{{.IntLen}}Key(\"{{.Key}}\", v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structMarshalUintTpl = t

	t, err = template.New("structMarshalCaseBool").
		Parse("\tenc.BoolKey(\"{{.Key}}\", v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structMarshalBoolTpl = t

	t, err = template.New("structMarshalCaseStruct").
		Parse("\tenc.ObjectKey(\"{{.Key}}\", v.{{.Field}})\n")
	if err != nil {
		log.Fatal(err)
	}
	structMarshalStructTpl = t

	t, err = template.New("structMarhalIsNil").
		Parse(isNilMethod)
	if err != nil {
		log.Fatal(err)
	}
	structIsNilTpl = t
}

func (v *vis) structGenIsNil(f *os.File, n string) error {
	err := structIsNilTpl.Execute(f, struct {
		StructName string
	}{
		StructName: n,
	})
	return err
}

func (v *vis) structGenMarshalObj(f *os.File, n string, s *ast.StructType) (int, error) {
	err := structMarshalDefTpl.Execute(f, struct {
		StructName string
	}{
		StructName: n,
	})
	if err != nil {
		return 0, err
	}
	keys := 0
	if len(s.Fields.List) > 0 {
		// TODO:  check tags
		for _, field := range s.Fields.List {
			switch t := field.Type.(type) {
			case *ast.Ident:
				var err error
				keys, err = v.structGenMarshalIdent(f, field, t, keys)
				if err != nil {
					return 0, err
				}
			case *ast.StarExpr:
				switch ptrExp := t.X.(type) {
				case *ast.Ident:
					var err error
					keys, err = v.structGenMarshalIdent(f, field, ptrExp, keys)
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
	}
	_, err = f.Write([]byte("}\n"))
	if err != nil {
		return 0, err
	}
	return keys, nil
}

func (v *vis) structGenMarshalIdent(f *os.File, field *ast.Field, i *ast.Ident, keys int) (int, error) {
	switch i.String() {
	case "string":
		var err = v.structMarshalString(f, field)
		if err != nil {
			return 0, err
		}
		keys++
	case "bool":
		var err = v.structMarshalBool(f, field)
		if err != nil {
			return 0, err
		}
		keys++
	case "int":
		var err = v.structMarshalInt(f, field, "")
		if err != nil {
			return 0, err
		}
		keys++
	case "int64":
		var err = v.structMarshalInt(f, field, "64")
		if err != nil {
			return 0, err
		}
		keys++
	case "int32":
		var err = v.structMarshalInt(f, field, "32")
		if err != nil {
			return 0, err
		}
		keys++
	case "int16":
		var err = v.structMarshalInt(f, field, "16")
		if err != nil {
			return 0, err
		}
		keys++
	case "int8":
		var err = v.structMarshalInt(f, field, "8")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint64":
		var err = v.structMarshalUint(f, field, "64")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint32":
		var err = v.structMarshalUint(f, field, "32")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint16":
		var err = v.structMarshalUint(f, field, "16")
		if err != nil {
			return 0, err
		}
		keys++
	case "uint8":
		var err = v.structMarshalUint(f, field, "8")
		if err != nil {
			return 0, err
		}
		keys++
	default:
		switch t := i.Obj.Decl.(type) {
		case *ast.TypeSpec:
			var err = v.structMarshalStruct(f, field, t)
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

func (v *vis) structMarshalString(f *os.File, field *ast.Field) error {
	key := field.Names[0].String()
	err := structMarshalStringTpl.Execute(f, struct {
		Field string
		Key   string
	}{key, strings.ToLower(key)})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structMarshalBool(f *os.File, field *ast.Field) error {
	key := field.Names[0].String()
	err := structMarshalBoolTpl.Execute(f, struct {
		Field string
		Key   string
	}{key, strings.ToLower(key)})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structMarshalInt(f *os.File, field *ast.Field, intLen string) error {
	key := field.Names[0].String()
	err := structMarshalIntTpl.Execute(f, struct {
		Field  string
		IntLen string
		Key    string
	}{key, intLen, strings.ToLower(key)})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structMarshalUint(f *os.File, field *ast.Field, intLen string) error {
	key := field.Names[0].String()
	err := structMarshalUintTpl.Execute(f, struct {
		Field  string
		IntLen string
		Key    string
	}{key, intLen, strings.ToLower(key)})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) structMarshalStruct(f *os.File, field *ast.Field, st *ast.TypeSpec) error {
	key := field.Names[0].String()
	var err = structMarshalStructTpl.Execute(f, struct {
		Key   string
		Field string
	}{strings.ToLower(key), key})
	if err != nil {
		return err
	}
	return nil
}
