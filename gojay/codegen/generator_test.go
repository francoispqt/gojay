package codegen

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"log"
	"path"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {

	parent := path.Join(toolbox.CallerDirectory(3), "test")

	var useCases = []struct {
		description string
		options     *Options
		hasError    bool
	}{
		{
			description: "basic struct code generation",
			options: &Options{
				Source: path.Join(parent, "basic_struct"),
				Types:  []string{"Message"},
				Dest:   path.Join(parent, "basic_struct", "encoding.go"),
			},
		},

		{
			description: "struct with pool code generation",
			options: &Options{
				Source:      path.Join(parent, "pooled_struct"),
				Types:       []string{"Message"},
				Dest:        path.Join(parent, "pooled_struct", "encoding.go"),
				PoolObjects: true,
			},
		},
		{
			description: "struct with embedded type code generation",
			options: &Options{
				Source:      path.Join(parent, "embedded_struct"),
				Types:       []string{"Message"},
				Dest:        path.Join(parent, "embedded_struct", "encoding.go"),
				PoolObjects: false,
			},
		},
		{
			description: "struct with json annotation and time/foarmat|layouat generation",
			options: &Options{
				Source:      path.Join(parent, "annotated_struct"),
				Types:       []string{"Message"},
				Dest:        path.Join(parent, "annotated_struct", "encoding.go"),
				PoolObjects: false,
				TagName:     "json",
			},
		},
	}

	for _, useCase := range useCases {
		gen := NewGenerator(useCase.options)
		err := gen.Generate()
		if useCase.hasError {
			assert.NotNil(t, err, useCase.description)
			continue
		}
		if !assert.Nil(t, err, useCase.description) {
			log.Fatal(err)
			continue
		}
	}

}
