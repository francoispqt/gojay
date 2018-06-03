package main

import (
	"go/ast"
	"log"
	"os"
	"text/template"
)

const genFileSuffix = "_gojay.go"

var pkgTpl *template.Template
var gojayImport = []byte("import \"github.com/francoispqt/gojay\"\n")

type gen struct {
	f        *os.File
	genTypes map[string]*ast.TypeSpec
	vis      *vis
}

type genTpl struct {
	strTpl string
	tpl    *template.Template
}

type templateList map[string]*genTpl

func init() {
	t, err := template.New("pkgDef").
		Parse("package {{.PkgName}} \n\n")
	if err != nil {
		log.Fatal(err)
	}
	pkgTpl = t
}

func parseTemplates(tpls templateList, pfx string) {
	for k, t := range tpls {
		tpl, err := template.New(pfx + k).Parse(t.strTpl)
		if err != nil {
			log.Fatal(err)
		}
		t.tpl = tpl
	}
}

func (g *gen) writePkg(pkg string) error {
	err := pkgTpl.Execute(g.f, struct {
		PkgName string
	}{
		PkgName: pkg,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) writeGojayImport() error {
	_, err := g.f.Write(gojayImport)
	if err != nil {
		return err
	}
	return nil
}

func (g *gen) gen(pkg string) error {
	// write package
	err := g.writePkg(pkg)
	if err != nil {
		return err
	}
	// write import of gojay
	err = g.writeGojayImport()
	if err != nil {
		return err
	}
	// range over specs
	// generate interfaces implementations based on type
	for _, s := range g.genTypes {
		switch t := s.Type.(type) {
		case *ast.StructType:
			err = g.genStruct(s.Name.String(), t)
			if err != nil {
				return err
			}
		case *ast.ArrayType:
			err = g.genArray(s.Name.String(), t)
			if err != nil {
				return err
			}
		case *ast.MapType:
			// TODO: generate for map type
		}
	}
	return nil
}
