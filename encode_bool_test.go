package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderBoolMarshalAPI(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		b, err := Marshal(true)
		assert.Nil(t, err, "err must be nil")
		assert.Equal(t, "true", string(b), "string(b) must be equal to 'true'")
	})
	t.Run("false", func(t *testing.T) {
		b, err := Marshal(false)
		assert.Nil(t, err, "err must be nil")
		assert.Equal(t, "false", string(b), "string(b) must be equal to 'false'")
	})
}

func TestEncoderBoolEncodeAPI(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := BorrowEncoder(builder)
		defer enc.Release()
		err := enc.EncodeBool(true)
		assert.Nil(t, err, "err must be nil")
		assert.Equal(t, "true", builder.String(), "string(b) must be equal to 'true'")
	})
	t.Run("false", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := BorrowEncoder(builder)
		defer enc.Release()
		err := enc.EncodeBool(false)
		assert.Nil(t, err, "err must be nil")
		assert.Equal(t, "false", builder.String(), "string(b) must be equal to 'false'")
	})
}

func TestEncoderBoolErrors(t *testing.T) {
	t.Run("pool-error", func(t *testing.T) {
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
	})
	t.Run("encode-api-write-error", func(t *testing.T) {
		v := true
		w := TestWriterError("")
		enc := BorrowEncoder(w)
		defer enc.Release()
		err := enc.EncodeBool(v)
		assert.NotNil(t, err, "err should not be nil")
	})
}
