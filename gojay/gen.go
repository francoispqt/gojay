package main

import (
	"go/ast"
	"log"
	"strings"
	"text/template"
)

const gojayAnnotation = "//gojay:json"
const genFileSuffix = "_gojay.go"

var pkgTpl *template.Template
var gojayImport = []byte("import \"github.com/francoispqt/gojay\"\n")

type Gen struct {
	b        *strings.Builder
	pkg      string
	src      string
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

// NewGen returns a new generator
func NewGen(p string) *Gen {
	g := &Gen{
		src:      p,
		b:        &strings.Builder{},
		genTypes: make(map[string]*ast.TypeSpec),
	}
	return g
}

func (g *Gen) writePkg(pkg string) error {
	err := pkgTpl.Execute(g.b, struct {
		PkgName string
	}{
		PkgName: pkg,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) writeGojayImport() error {
	_, err := g.b.Write(gojayImport)
	if err != nil {
		return err
	}
	return nil
}

func (g *Gen) Gen() error {
	// write package
	err := g.writePkg(g.pkg)
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
		// is struct
		case *ast.StructType:
			err = g.genStruct(s.Name.String(), t)
			if err != nil {
				return err
			}
		// is array
		case *ast.ArrayType:
			err = g.genArray(s.Name.String(), t)
			if err != nil {
				return err
			}
		// is map
		case *ast.MapType:
			// TODO: generate for map type
		}
	}
	return nil
}
