package main

import (
	"go/ast"
	"html/template"
	"log"
	"os"
	"strings"
)

var structMarshalDefTpl *template.Template
var structMarshalStringTpl *template.Template
var structMarshalIntTpl *template.Template
var structMarshalUintTpl *template.Template
var structMarshalBoolTpl *template.Template

var structIsNilTpl *template.Template

var isNilMethod = `
// IsNil returns wether the structure is nil value or not
func (v *{{.StructName}}) IsNil() bool { return v == nil }
`

func init() {
	t, err := template.New("structUnmarshalDef").
		Parse("\n// MarshalJSONObject implements gojay's MarshalerJSONObject" +
			"\nfunc (v *{{.StructName}}) MarshalJSONOject(enc *gojay.Encoder) {\n",
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
				switch t.String() {
				case "string":
					err = v.structMarshalString(f, field)
					if err != nil {
						return 0, err
					}
					keys++
				case "bool":
					err = v.structMarshalBool(f, field)
					if err != nil {
						return 0, err
					}
					keys++
				case "int":
					err = v.structMarshalInt(f, field, "")
					if err != nil {
						return 0, err
					}
					keys++
				case "int64":
					err = v.structMarshalInt(f, field, "64")
					if err != nil {
						return 0, err
					}
					keys++
				case "int32":
					err = v.structMarshalInt(f, field, "32")
					if err != nil {
						return 0, err
					}
					keys++
				case "int16":
					err = v.structMarshalInt(f, field, "16")
					if err != nil {
						return 0, err
					}
					keys++
				case "int8":
					err = v.structMarshalInt(f, field, "8")
					if err != nil {
						return 0, err
					}
					keys++
				case "uint64":
					err = v.structMarshalUint(f, field, "64")
					if err != nil {
						return 0, err
					}
					keys++
				case "uint32":
					err = v.structMarshalUint(f, field, "32")
					if err != nil {
						return 0, err
					}
					keys++
				case "uint16":
					err = v.structMarshalUint(f, field, "16")
					if err != nil {
						return 0, err
					}
					keys++
				case "uint8":
					err = v.structMarshalUint(f, field, "8")
					if err != nil {
						return 0, err
					}
					keys++
				}
			}
		}
	}
	_, err = f.Write([]byte("}\n"))
	if err != nil {
		return 0, err
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
