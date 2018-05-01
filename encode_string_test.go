package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderStringEncodeAPI(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString("hello world")
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"hello world"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("utf8", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString("漢字")
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"漢字"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
}

func TestEncoderStringEncodeAPIErrors(t *testing.T) {
	t.Run("pool-error", func(t *testing.T) {
		v := ""
		enc := BorrowEncoder(nil)
		enc.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
			assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = enc.EncodeString(v)
		assert.True(t, false, "should not be called as it should have panicked")
	})
	t.Run("write-error", func(t *testing.T) {
		v := "test"
		w := TestWriterError("")
		enc := BorrowEncoder(w)
		defer enc.Release()
		err := enc.EncodeString(v)
		assert.NotNil(t, err, "err should not be nil")
	})
}

func TestEncoderStringMarshalAPI(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		r, err := Marshal("string")
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"string"`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("utf8", func(t *testing.T) {
		r, err := Marshal("漢字")
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"漢字"`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
}
