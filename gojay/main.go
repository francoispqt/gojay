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
	files, err := getFiles()
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, f, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		v := &vis{pkg: node.Name.String(), specs: make([]*ast.TypeSpec, 0), file: f}
		ast.Walk(v, node)
		err = v.gen()
		if err != nil {
			log.Fatal(err)
		}
	}
}
