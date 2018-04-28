package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderStringBasic(t *testing.T) {
	json := []byte(`"string"`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "string", v, "v must be equal to 'string'")
}

func TestDecoderStringComplex(t *testing.T) {
	json := []byte(`  "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char", v, "v is not equal to the value expected")
}

func TestDecoderStringNull(t *testing.T) {
	json := []byte(`null`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to ''")
}

func TestDecoderStringInvalidJSON(t *testing.T) {
	json := []byte(`"invalid JSONs`)
	var v string
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderStringInvalidType(t *testing.T) {
	json := []byte(`1`)
	var v string
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidTypeError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderStringDecoderAPI(t *testing.T) {
	var v string
	dec := NewDecoder(strings.NewReader(`"hello world!"`))
	defer dec.Release()
	err := dec.DecodeString(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "hello world!", v, "v must be equal to 'hello world!'")
}

func TestDecoderStringPoolError(t *testing.T) {
	result := ""
	dec := NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = dec.DecodeString(&result)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
