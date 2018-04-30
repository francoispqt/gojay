package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderBoolTrue(t *testing.T) {
	enc := BorrowEncoder(nil)
	defer enc.Release()
	b, err := enc.EncodeBool(true)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, "true", string(b), "string(b) must be equal to 'true'")
}

func TestEncoderBoolFalse(t *testing.T) {
	enc := BorrowEncoder(nil)
	defer enc.Release()
	b, err := enc.EncodeBool(false)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, "false", string(b), "string(b) must be equal to 'false'")
}

func TestEncoderBoolPoolError(t *testing.T) {
	enc := BorrowEncoder(nil)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
		assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledEncoderError")
	}()
	_, _ = enc.EncodeBool(false)
	assert.True(t, false, "should not be called as it should have panicked")
}
