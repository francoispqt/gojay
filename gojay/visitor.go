package main

import (
	"go/ast"
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
	specs        []*ast.TypeSpec
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
			v.specs = append(v.specs, n)
		}
		v.commentFound = false
		return v
	case *ast.StructType:
		v.commentFound = false
		return nil
	}
	return nil
}
