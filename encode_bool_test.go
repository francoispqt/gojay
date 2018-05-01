package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderBoolTrue(t *testing.T) {
	builder := &strings.Builder{}
	enc := BorrowEncoder(builder)
	defer enc.Release()
	err := enc.EncodeBool(true)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, "true", builder.String(), "string(b) must be equal to 'true'")
}

func TestEncoderBoolFalse(t *testing.T) {
	builder := &strings.Builder{}
	enc := BorrowEncoder(builder)
	defer enc.Release()
	err := enc.EncodeBool(false)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, "false", builder.String(), "string(b) must be equal to 'false'")
}

func TestEncoderBoolPoolError(t *testing.T) {
	builder := &strings.Builder{}
	enc := BorrowEncoder(builder)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
		assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledEncoderError")
	}()
	_ = enc.EncodeBool(false)
	assert.True(t, false, "should not be called as it should have panicked")
}
func TestEncoderBoolPoolEncoderAPIWriteError(t *testing.T) {
	v := true
	w := TestWriterError("")
	enc := BorrowEncoder(w)
	defer enc.Release()
	err := enc.EncodeBool(v)
	assert.NotNil(t, err, "err should not be nil")
}
