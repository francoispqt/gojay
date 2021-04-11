package main

import (
	"flag"
	"log"

	"github.com/francoispqt/gojay/gojay/codegen"
)

var pkg = flag.String("pkg", "", "the package name of the generated file")
var dst = flag.String("o", "", "destination file to output generated code")
var src = flag.String("s", "", "source dir or file (absolute or relative path)")
var types = flag.String("t", "", "types to generate")
var annotation = flag.String("a", "json", "annotation tag (default json)")
var poolObjects = flag.String("p", "", "generate code to reuse objects using sync.Pool")
var errOnUnknown = flag.String("e", "", "generate code to error on unknown fields")

func main() {
	flag.Parse()
	options := codegen.NewOptionsWithFlagSet(flag.CommandLine)
	gen := codegen.NewGenerator(options)
	if err := gen.Generate(); err != nil {
		log.Fatal(err)
	}
}
