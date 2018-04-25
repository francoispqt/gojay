package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderBoolTrue(t *testing.T) {
	json := []byte(`true`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, true, v, "v must be equal to true")
}

func TestDecoderBoolFalse(t *testing.T) {
	json := []byte(`false`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, false, v, "v must be equal to false")
}

func TestDecoderBoolInvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v bool
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, false, v, "v must be equal to false as it is zero val")
}

func TestDecoderBoolNonBooleanJSONFalse(t *testing.T) {
	json := []byte(`null`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, false, v, "v must be equal to true")
}

func TestDecoderBoolInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	var v bool
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}
