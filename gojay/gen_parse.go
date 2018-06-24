//+build !test

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func (g *Gen) parse() error {
	var f, err = os.Stat(g.src)
	if err != nil {
		return err
	}
	if f.IsDir() {
		err = g.parseDir()
	} else {
		err = g.parseFile()
	}
	return err
}

func (g *Gen) parseDir() error {
	// parse the given path
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, g.src, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	// range across packages
	for pkgName, pkg := range pkgs {
		v := newVisitor(g, pkgName)
		g.pkg = pkgName
		// range on files in package
		for _, f := range pkg.Files {
			ast.Walk(v, f)
			if err != nil {
				return err
			}
		}
		g.vis = v
	}
	return nil
}

func (g *Gen) parseFile() error {
	// parse the given path
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, g.src, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	g.pkg = f.Name.Name
	v := newVisitor(g, g.pkg)
	ast.Walk(v, f)
	if err != nil {
		return err
	}
	g.vis = v
	return nil
}
