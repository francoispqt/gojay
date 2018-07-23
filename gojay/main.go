//+build !test

package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dst = flag.String("o", "", "destination file to output generated implementations")
var src = flag.String("s", "", "source dir or file (absolute or relative path)")
var pkg = flag.String("p", "", "go package")
var types = flag.String("t", "", "types to generate")

var ErrNoPathProvided = errors.New("You must provide a path or a package name")

type stringWriter interface {
	WriteString(string) (int, error)
}

func hasAnnotation(fP string) bool {
	b, err := ioutil.ReadFile(fP)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Contains(string(b), gojayAnnotation)
}

func resolvePath(p string) (string, error) {
	if fullPath, err := filepath.Abs(p); err != nil {
		return "", err
	} else if _, err := os.Stat(fullPath); err != nil {
		return "", err
	} else {
		return fullPath, nil
	}
}

// getPath returns either the path given as argument or current working directory
func getPath() (string, error) {
	// if pkg is set, resolve pkg path
	if *pkg != "" {
		return resolvePath(os.Getenv("GOPATH") + "/src/" + *pkg)
	} else if *src != "" { // if src is present parse from src
		return resolvePath(*src)
	} else if len(os.Args) > 1 { // else if there is a command line arg, use it as path to a package $GOPATH/src/os.Args[1]
		return resolvePath(os.Getenv("GOPATH") + "/src/" + os.Args[1])
	}
	return "", ErrNoPathProvided
}

// getTypes returns the types to be parsed
func getTypes() (t []string) {
	if *types != "" { // if src is present parse from src
		return strings.Split(*types, ",")
	} else if *src == "." && *dst == "" && len(os.Args) > 2 { // else if there is a command line arg, use it as path to a package $GOPATH/src/os.Args[1]
		return strings.Split(os.Args[2], ",")
	}
	return t
}

// getOutput returns the output
func getOutput() (stringWriter, error) {
	if *dst != "" {
		p, err := filepath.Abs(*dst)
		if err != nil {
			return nil, err
		}
		return os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	} else if len(os.Args) > 3 && *src == "" && *types == "" {
		p, err := filepath.Abs(os.Args[3])
		if err != nil {
			return nil, err
		}
		return os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	}
	return os.Stdout, nil
}

func parseArgs() (p string, t []string, err error) {
	flag.Parse()
	p, err = getPath()
	if err != nil {
		return p, t, err
	}
	t = getTypes()
	return p, t, err
}

func main() {
	p, t, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}
	// parse source files
	g := NewGen(p, t)
	err = g.parse()
	if err != nil {
		log.Fatal(err)
		return
	}
	// generate output
	err = g.Gen()
	if err != nil {
		log.Fatal(err)
		return
	}
	var o stringWriter
	o, err = getOutput()
	if err != nil {
		log.Fatal(err)
		return
	}
	// write content to output
	o.WriteString(g.b.String())
}
