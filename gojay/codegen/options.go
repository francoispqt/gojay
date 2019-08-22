package codegen

import (
	"flag"
	"strings"

	"github.com/go-errors/errors"
	"github.com/viant/toolbox"
	"github.com/viant/toolbox/url"
)

type Options struct {
	Source      string
	Dest        string
	Types       []string
	PoolObjects bool
	TagName     string
	Pkg         string
	ExtImports  []ExtImport
}

type ExtImport struct {
	Identifier string
	Path       string
}

func (o *Options) Validate() error {
	if o.Source == "" {
		return errors.New("Source was empty")
	}
	if len(o.Types) == 0 {
		return errors.New("Types was empty")
	}
	return nil
}

const (
	optionKeySource      = "s"
	optionKeyDest        = "o"
	optionKeyTypes       = "t"
	optionKeyTagName     = "a"
	optionKeyPoolObjects = "p"
	optionKeyPkg         = "pkg"
	optionKeyExtImports  = "imp"
)

//NewOptionsWithFlagSet creates a new options for the supplide flagset
func NewOptionsWithFlagSet(set *flag.FlagSet) *Options {
	toolbox.Dump(set)

	var result = &Options{}
	result.Dest = set.Lookup(optionKeyDest).Value.String()
	result.Source = set.Lookup(optionKeySource).Value.String()
	result.PoolObjects = toolbox.AsBoolean(set.Lookup(optionKeyPoolObjects).Value.String())
	result.TagName = set.Lookup(optionKeyTagName).Value.String()
	result.Types = strings.Split(set.Lookup(optionKeyTypes).Value.String(), ",")
	result.Pkg = set.Lookup(optionKeyPkg).Value.String()
	if result.Source == "" {
		result.Source = url.NewResource(".").ParsedURL.Path
	}
	var extImportSlice = strings.Split(set.Lookup(optionKeyExtImports).Value.String(), ",")
	result.ExtImports = make([]ExtImport, len(extImportSlice))
	for _, impStr := range extImportSlice {
		var imp = strings.Split(impStr, ":")
		var extImp = ExtImport{
			Path: imp[0],
		}
		if len(imp) > 1 {
			extImp = ExtImport{
				Identifier: imp[0],
				Path:       imp[1],
			}
		}
		result.ExtImports = append(result.ExtImports, extImp)
	}
	return result
}
