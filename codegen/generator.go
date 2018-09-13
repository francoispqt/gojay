package codegen

import (
	"fmt"
	"github.com/viant/toolbox"
	"io/ioutil"
	"strings"
)

type Generator struct {
	fileInfo      *toolbox.FileSetInfo
	types         map[string]string
	structTypes   map[string]string
	sliceTypes    map[string]string
	pooledObjects map[string]string
	poolInit      map[string]string
	imports       map[string]bool
	filedInit     []string
	Pkg           string
	Code          string
	Init          string
	Imports       string
	options       *Options
}

func (g *Generator) Type(typeName string) *toolbox.TypeInfo {
	return g.fileInfo.Type(typeName)
}

func (g *Generator) addImport(pkg string) {
	g.imports[`"`+pkg+`"`] = true
}

func (g *Generator) init() {
	g.filedInit = []string{}
	g.imports = map[string]bool{}
	g.pooledObjects = map[string]string{}
	g.structTypes = map[string]string{}
	g.sliceTypes = map[string]string{}
	g.poolInit = map[string]string{}
	g.addImport("github.com/francoispqt/gojay")
}

func (g *Generator) Generate(options *Options) error {
	if err := options.Validate(); err != nil {
		return err
	}
	g.options = options
	g.init()

	if options.PoolObjects {
		g.addImport("sync")
	}
	if err := g.readPackageCode(options.Source); err != nil {
		return err
	}
	for _, rootType := range options.Types {
		if err := g.generateStructCode(rootType); err != nil {
			return err
		}
	}
	g.Imports = strings.Join(toolbox.MapKeysToStringSlice(g.imports), "\n")
	return g.writeCode()
}

func (g *Generator) writeCode() error {
	var generatedCode = []string{}
	for _, code := range g.pooledObjects {
		generatedCode = append(generatedCode, code)
	}
	generatedCode = append(generatedCode, "")
	for _, code := range g.sliceTypes {
		generatedCode = append(generatedCode, code)
	}
	generatedCode = append(generatedCode, "")
	for _, code := range g.structTypes {
		generatedCode = append(generatedCode, code)
	}

	for _, code := range g.poolInit {
		if g.Init != "" {
			g.Init += "\n"
		}
		g.Init += code
	}
	g.Code = strings.Join(generatedCode, "\n")

	expandedCode, err := expandBlockTemplate(fileCode, g)
	if err != nil {
		return err
	}
	if g.options.Dest == "" {
		fmt.Print(expandedCode)
		return nil
	}
	return ioutil.WriteFile(g.options.Dest, []byte(expandedCode), 0644)
}

func (g *Generator) generatePrimitiveArray(field *Field) error {
	key := field.ComponentType + toolbox.AsString(field.IsPointerComponent)
	if _, ok := g.sliceTypes[key]; ok {
		return nil
	}
	code, err := expandBlockTemplate(baseTypeSlice, field)
	g.sliceTypes[key] = code
	return err
}

func (g *Generator) generateObjectArray(field *Field) error {
	if _, ok := g.sliceTypes[field.RawComponentType]; ok {
		return nil
	}

	if err := g.generateStructCode(field.ComponentType); err != nil {
		return err
	}
	code, err := expandBlockTemplate(structTypeSlice, field)
	if err != nil {
		return err
	}
	g.sliceTypes[field.RawComponentType] = code
	return err
}

func (g *Generator) generatePool(structType string) error {
	if !g.options.PoolObjects {
		return nil
	}
	var err error
	if g.pooledObjects[structType], err = expandBlockTemplate(poolVar, struct {
		PoolName string
	}{getPoolName(structType)}); err == nil {
		g.poolInit[structType], err = expandBlockTemplate(poolInit, struct {
			PoolName string
			Type     string
		}{getPoolName(structType), structType})
	}
	return err

}

func (g *Generator) generateStructCode(structType string) error {
	structType = normalizeTypeName(structType)
	typeInfo := g.Type(structType)
	if typeInfo == nil {
		return nil
	}
	if _, hasCode := g.structTypes[structType]; hasCode {
		return nil
	}
	g.generatePool(structType)
	aStruct := NewStruct(typeInfo, g)
	code, err := aStruct.Generate()
	if err != nil {
		return err
	}
	g.structTypes[structType] = code
	return nil
}

func (g *Generator) readPackageCode(pkgPath string) error {
	var err error
	fragments := strings.Split(pkgPath, "/")
	g.Pkg = fragments[len(fragments)-1]
	g.fileInfo, err = toolbox.NewFileSetInfo(pkgPath)
	return err
}
