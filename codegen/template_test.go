package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ExpandTemplate(t *testing.T) {
	var dictionary = map[int]string{
		1: `type {{.TypeName}} {{.SourceTypeName}}`,
	}
	expaded, err := expandTemplate("test", dictionary, 1, struct {
		TypeName       string
		SourceTypeName string
	}{"A", "B"})

	assert.Nil(t, err)
	assert.Equal(t, "type A B", expaded)
}
