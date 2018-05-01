package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderInt(t *testing.T) {
	r, err := Marshal(1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderIntEncodeAPI(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	err := enc.EncodeInt(1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		builder.String(),
		"Result of marshalling is different as the one expected")
}
func TestEncoderIntEncodeAPIPoolError(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err should not be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
	}()
	_ = enc.EncodeInt(1)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
func TestEncoderIntEncodeAPIWriteError(t *testing.T) {
	w := TestWriterError("")
	enc := NewEncoder(w)
	err := enc.EncodeInt(1)
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "Test Error", err.Error(), "err should be of type InvalidUsagePooledEncoderError")
}

func TestEncoderInt64(t *testing.T) {
	r, err := Marshal(int64(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderInt64EncodeAPI(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	err := enc.EncodeInt64(int64(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		builder.String(),
		"Result of marshalling is different as the one expected")
}
func TestEncoderInt64EncodeAPIPoolError(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err should not be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
	}()
	_ = enc.EncodeInt64(1)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
func TestEncoderInt64EncodeAPIWriteError(t *testing.T) {
	w := TestWriterError("")
	enc := NewEncoder(w)
	err := enc.EncodeInt64(1)
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "Test Error", err.Error(), "err should be of type InvalidUsagePooledEncoderError")
}

func TestEncoderInt32(t *testing.T) {
	r, err := Marshal(int32(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderInt16(t *testing.T) {
	r, err := Marshal(int16(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderInt8(t *testing.T) {
	r, err := Marshal(int8(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderUint64(t *testing.T) {
	r, err := Marshal(uint64(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderUint32(t *testing.T) {
	r, err := Marshal(uint32(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderUint16(t *testing.T) {
	r, err := Marshal(uint16(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderUint8(t *testing.T) {
	r, err := Marshal(uint8(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderFloat(t *testing.T) {
	r, err := Marshal(1.1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1.1`,
		string(r),
		"Result of marshalling is different as the one expected")
}
func TestEncoderFloatEncodeAPI(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	err := enc.EncodeFloat(float64(1.1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1.1`,
		builder.String(),
		"Result of marshalling is different as the one expected")
}
func TestEncoderFloatEncodeAPIPoolError(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err should not be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
	}()
	_ = enc.EncodeFloat(1.1)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
func TestEncoderFloatEncodeAPIWriteError(t *testing.T) {
	w := TestWriterError("")
	enc := NewEncoder(w)
	err := enc.EncodeFloat(1.1)
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "Test Error", err.Error(), "err should be of type InvalidUsagePooledEncoderError")
}

func TestEncoderFloat32EncodeAPI(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	err := enc.EncodeFloat32(float32(1.12))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1.12`,
		builder.String(),
		"Result of marshalling is different as the one expected")
}
func TestEncoderFloat32EncodeAPIPoolError(t *testing.T) {
	builder := &strings.Builder{}
	enc := NewEncoder(builder)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err should not be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
	}()
	_ = enc.EncodeFloat32(float32(1.1))
	assert.True(t, false, "should not be called as decoder should have panicked")
}
func TestEncoderFloat32EncodeAPIWriteError(t *testing.T) {
	w := TestWriterError("")
	enc := NewEncoder(w)
	err := enc.EncodeFloat32(float32(1.1))
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "Test Error", err.Error(), "err should be of type InvalidUsagePooledEncoderError")
}

func TestEncoderIntPooledError(t *testing.T) {
	v := 1
	enc := BorrowEncoder(nil)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
		assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = enc.EncodeInt(v)
	assert.True(t, false, "should not be called as it should have panicked")
}

func TestEncoderFloatPooledError(t *testing.T) {
	v := 1.1
	enc := BorrowEncoder(nil)
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
		assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = enc.EncodeFloat(v)
	assert.True(t, false, "should not be called as it should have panicked")
}
