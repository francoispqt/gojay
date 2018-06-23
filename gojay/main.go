//+build !test

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dst = flag.String("o", "", "destination file to output generated implementations")
var src = flag.String("s", "", "source dir or file")
var types = flag.String("t", "", "types to generate")

func hasAnnotation(fP string) bool {
	b, err := ioutil.ReadFile(fP)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Contains(string(b), gojayAnnotation)
}

// getPath returns either the path given as argument or current working directory
func getPath() (string, error) {
	var err error
	var p string
	if *src != "" { // if src is present parse from src
		p, err = filepath.Abs(*src)
		if err != nil {
			return "", err
		}
	} else if len(os.Args) > 1 { // else if there is a command line arg, use it as path to a package $GOPATH/src/os.Args[1]
		p, err = filepath.Abs(os.Getenv("GOPATH") + "/src/" + os.Args[1])
		if err != nil {
			return "", err
		}
	} else {
		p, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}
	return p, nil
}

func getTypes() (t []string) {
	if *types != "" { // if src is present parse from src
		return strings.Split(*types, ",")
	} else if *src == "" && *dst == "" && len(os.Args) > 2 { // else if there is a command line arg, use it as path to a package $GOPATH/src/os.Args[1]
		return strings.Split(os.Args[2], ",")
	}
	return t
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
	// if has dst file, write to file
	if *dst != "" {
		f, err := os.OpenFile(*dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
			return
		}
		f.WriteString(g.b.String())
		return
	}
	// else just print to stdout
	os.Stdout.WriteString(g.b.String())
}
