package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const gojayAnnotation = "//gojay:json"

func hasAnnotation(fP string) bool {
	b, err := ioutil.ReadFile(fP)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Contains(string(b), gojayAnnotation)
}

func getPath() (string, error) {
	p := os.Args[1]
	return filepath.Abs(p)
}

func getFiles() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Gojay generator takes one argument, 0 given")
	}
	p, err := getPath()
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}
	r := make([]string, 0)
	for _, f := range files {
		fP := filepath.Join(p, f.Name())
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") && hasAnnotation(fP) {
			r = append(r, fP)
		}
	}
	return r, nil
}

func main() {
	p, err := getPath()
	if err != nil {
		log.Fatal(err)
	}
	// for _, f := range files {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, p, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	for pkgName, pkg := range pkgs {
		v := NewVisitor(pkgName)
		for fileName, f := range pkg.Files {
			v.file = fileName[:len(fileName)-3] + genFileSuffix
			ast.Walk(v, f)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = v.gen()
		if err != nil {
			log.Fatal(err)
		}
	}
}
