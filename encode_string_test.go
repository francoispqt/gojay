package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderString(t *testing.T) {
	r, err := Marshal("string")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`"string"`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderStringUTF8(t *testing.T) {
	r, err := Marshal("漢字")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`"漢字"`,
		string(r),
		"Result of marshalling is different as the one expected")
}
