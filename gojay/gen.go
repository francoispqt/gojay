package main

import (
	"go/ast"
	"html/template"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

const genFileSuffix = "_gojay.go"

var pkgTpl *template.Template
var gojayImport = []byte("import \"github.com/francoispqt/gojay\"\n")

func init() {
	t, err := template.New("pkgDef").
		Parse("package {{.PkgName}} \n\n")
	if err != nil {
		log.Fatal(err)
	}
	pkgTpl = t
}

func (v *vis) gen() error {
	// open the file
	f, err := os.OpenFile(v.genFileName(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	// write package
	err = v.writePkg(f)
	if err != nil {
		return err
	}
	// write import of gojay
	err = v.writeGojayImport(f)
	if err != nil {
		return err
	}
	// range over specs
	// generate interfaces implementations based on type
	for _, s := range v.specs {
		switch t := s.Type.(type) {
		case *ast.StructType:
			err = v.genStruct(f, s.Name.String(), t)
			if err != nil {
				return err
			}
		case *ast.ArrayType:
			spew.Println(t, "arr")
		}
	}
	return nil
}

func (v *vis) genFileName() string {
	return v.file[:len(v.file)-3] + genFileSuffix
}

func (v *vis) writePkg(f *os.File) error {
	err := pkgTpl.Execute(f, struct {
		PkgName string
	}{
		PkgName: v.pkg,
	})
	if err != nil {
		return err
	}
	return nil
}

func (v *vis) writeGojayImport(f *os.File) error {
	_, err := f.Write(gojayImport)
	if err != nil {
		return err
	}
	return nil
}
