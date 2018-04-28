package gojay

import (
	"strings"
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
	assert.Equal(t, false, v, "v must be equal to false")
}

func TestDecoderBoolInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	var v bool
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}
func TestDecoderBoolDecoderAPI(t *testing.T) {
	var v bool
	dec := NewDecoder(strings.NewReader("true"))
	defer dec.Release()
	err := dec.DecodeBool(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, true, v, "v must be equal to true")
}

func TestDecoderBoolPoolError(t *testing.T) {
	v := true
	dec := NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = dec.DecodeBool(&v)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
