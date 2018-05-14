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
		err := enc.EncodeString("漢字𩸽")
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"漢字𩸽"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("utf8-multibyte", func(t *testing.T) {
		str := "テュールスト マーティン ヤコブ 😁"
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"テュールスト マーティン ヤコブ 😁"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("escaped-sequence1", func(t *testing.T) {
		str := `テュールスト マ\ーテ
ィン ヤコブ 😁`
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"テュールスト マ\\ーテ\nィン ヤコブ 😁"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("escaped-sequence2", func(t *testing.T) {
		str := `テュールスト マ\ーテ
ィン ヤコブ 😁	`
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"テュールスト マ\\ーテ\nィン ヤコブ 😁\t"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("escaped-sequence3", func(t *testing.T) {
		str := "hello \r world 𝄞"
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"hello \r world 𝄞"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("escaped-sequence3", func(t *testing.T) {
		str := "hello \b world 𝄞"
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`"hello \b world 𝄞"`,
			builder.String(),
			"Result of marshalling is different as the one expected")
	})
	t.Run("escaped-sequence3", func(t *testing.T) {
		str := "hello \f world 𝄞"
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeString(str)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			"\"hello \\f world 𝄞\"",
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
