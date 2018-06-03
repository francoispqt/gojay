package main

import (
	"go/ast"
	"os"
	"strings"
)

func docContains(n *ast.CommentGroup, s string) bool {
	for _, d := range n.List {
		if strings.Contains(d.Text, s) {
			return true
		}
	}
	return false
}

type vis struct {
	pkg          string
	specs        map[string]*ast.TypeSpec
	files        map[string]map[string]*ast.TypeSpec
	file         string
	commentFound bool
}

func (v *vis) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		v.commentFound = false
		return v
	case *ast.File:
		v.commentFound = false
		return v
	case *ast.GenDecl:
		if n.Doc != nil {
			v.commentFound = docContains(n.Doc, gojayAnnotation)
		}
		return v
	case *ast.TypeSpec:
		if v.commentFound {
			v.specs[n.Name.Name] = n
			if v.files[v.file] == nil {
				v.files[v.file] = make(map[string]*ast.TypeSpec)
			}
			v.files[v.file][n.Name.Name] = n
		}
		v.commentFound = false
		return v
	case *ast.StructType:
		v.commentFound = false
		return v
	}
	return v
}

func (v *vis) gen() error {
	for fileName, genTypes := range v.files {
		// open the file
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		defer f.Close()
		g := &gen{f, genTypes, v}
		err = g.gen(v.pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewVisitor(pkgName string) *vis {
	return &vis{
		pkg:   pkgName,
		specs: make(map[string]*ast.TypeSpec),
		files: make(map[string]map[string]*ast.TypeSpec),
	}
}
