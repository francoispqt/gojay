package gojay

import (
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

func TestEncoderInt64(t *testing.T) {
	r, err := Marshal(int64(1))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`1`,
		string(r),
		"Result of marshalling is different as the one expected")
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
