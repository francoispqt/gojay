package gojay

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderInt64(t *testing.T) {
	var testCasesBasic = []struct {
		name         string
		v            int64
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int64(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt64,
			expectedJSON: "[9223372036854775807,9223372036854775807]",
		},
		{
			name:         "big",
			v:            int64(0),
			expectedJSON: "[0,0]",
		},
	}
	for _, testCase := range testCasesBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int64(testCase.v)
				enc.AddInt64(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}

	var testCasesOmitEmpty = []struct {
		name         string
		v            int64
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int64(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt64,
			expectedJSON: "[9223372036854775807,9223372036854775807]",
		},
		{
			name:         "big",
			v:            int64(0),
			expectedJSON: "[]",
		},
	}
	for _, testCase := range testCasesOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int64OmitEmpty(testCase.v)
				enc.AddInt64OmitEmpty(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesKeyBasic = []struct {
		name         string
		v            int64
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int64(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt64,
			expectedJSON: `{"foo":9223372036854775807,"bar":9223372036854775807}`,
		},
		{
			name:         "big",
			v:            int64(0),
			expectedJSON: `{"foo":0,"bar":0}`,
		},
	}
	for _, testCase := range testCasesKeyBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int64Key("foo", testCase.v)
				enc.AddInt64Key("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}

	var testCasesKeyOmitEmpty = []struct {
		name         string
		v            int64
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int64(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt64,
			expectedJSON: `{"foo":9223372036854775807,"bar":9223372036854775807}`,
		},
		{
			name:         "big",
			v:            int64(0),
			expectedJSON: `{}`,
		},
	}
	for _, testCase := range testCasesKeyOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int64KeyOmitEmpty("foo", testCase.v)
				enc.AddInt64KeyOmitEmpty("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
}

func TestEncoderInt32(t *testing.T) {
	var testCasesBasic = []struct {
		name         string
		v            int32
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int32(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt32,
			expectedJSON: "[2147483647,2147483647]",
		},
		{
			name:         "big",
			v:            int32(0),
			expectedJSON: "[0,0]",
		},
	}
	for _, testCase := range testCasesBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int32(testCase.v)
				enc.AddInt32(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesOmitEmpty = []struct {
		name         string
		v            int32
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int32(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt32,
			expectedJSON: "[2147483647,2147483647]",
		},
		{
			name:         "big",
			v:            int32(0),
			expectedJSON: "[]",
		},
	}
	for _, testCase := range testCasesOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int32OmitEmpty(testCase.v)
				enc.AddInt32OmitEmpty(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesKeyBasic = []struct {
		name         string
		v            int32
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int32(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt32,
			expectedJSON: `{"foo":2147483647,"bar":2147483647}`,
		},
		{
			name:         "big",
			v:            int32(0),
			expectedJSON: `{"foo":0,"bar":0}`,
		},
	}
	for _, testCase := range testCasesKeyBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int32Key("foo", testCase.v)
				enc.AddInt32Key("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}

	var testCasesKeyOmitEmpty = []struct {
		name         string
		v            int32
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int32(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt32,
			expectedJSON: `{"foo":2147483647,"bar":2147483647}`,
		},
		{
			name:         "big",
			v:            int32(0),
			expectedJSON: `{}`,
		},
	}
	for _, testCase := range testCasesKeyOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int32KeyOmitEmpty("foo", testCase.v)
				enc.AddInt32KeyOmitEmpty("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
}

func TestEncoderInt16(t *testing.T) {
	var testCasesBasic = []struct {
		name         string
		v            int16
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int16(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt16,
			expectedJSON: "[32767,32767]",
		},
		{
			name:         "big",
			v:            int16(0),
			expectedJSON: "[0,0]",
		},
	}
	for _, testCase := range testCasesBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int16(testCase.v)
				enc.AddInt16(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesOmitEmpty = []struct {
		name         string
		v            int16
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int16(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt16,
			expectedJSON: "[32767,32767]",
		},
		{
			name:         "big",
			v:            int16(0),
			expectedJSON: "[]",
		},
	}
	for _, testCase := range testCasesOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int16OmitEmpty(testCase.v)
				enc.AddInt16OmitEmpty(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesKeyBasic = []struct {
		name         string
		v            int16
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int16(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt16,
			expectedJSON: `{"foo":32767,"bar":32767}`,
		},
		{
			name:         "big",
			v:            int16(0),
			expectedJSON: `{"foo":0,"bar":0}`,
		},
	}
	for _, testCase := range testCasesKeyBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int16Key("foo", testCase.v)
				enc.AddInt16Key("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}

	var testCasesKeyOmitEmpty = []struct {
		name         string
		v            int16
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int16(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt16,
			expectedJSON: `{"foo":32767,"bar":32767}`,
		},
		{
			name:         "big",
			v:            int16(0),
			expectedJSON: `{}`,
		},
	}
	for _, testCase := range testCasesKeyOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int16KeyOmitEmpty("foo", testCase.v)
				enc.AddInt16KeyOmitEmpty("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
}

func TestEncoderInt8(t *testing.T) {
	var testCasesBasic = []struct {
		name         string
		v            int8
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int8(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt8,
			expectedJSON: "[127,127]",
		},
		{
			name:         "big",
			v:            int8(0),
			expectedJSON: "[0,0]",
		},
	}
	for _, testCase := range testCasesBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int8(testCase.v)
				enc.AddInt8(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesOmitEmpty = []struct {
		name         string
		v            int8
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int8(1),
			expectedJSON: "[1,1]",
		},
		{
			name:         "big",
			v:            math.MaxInt8,
			expectedJSON: "[127,127]",
		},
		{
			name:         "big",
			v:            int8(0),
			expectedJSON: "[]",
		},
	}
	for _, testCase := range testCasesOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeArrayFunc(func(enc *Encoder) {
				enc.Int8OmitEmpty(testCase.v)
				enc.AddInt8OmitEmpty(testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
	var testCasesKeyBasic = []struct {
		name         string
		v            int8
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int8(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt8,
			expectedJSON: `{"foo":127,"bar":127}`,
		},
		{
			name:         "big",
			v:            int8(0),
			expectedJSON: `{"foo":0,"bar":0}`,
		},
	}
	for _, testCase := range testCasesKeyBasic {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int8Key("foo", testCase.v)
				enc.AddInt8Key("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}

	var testCasesKeyOmitEmpty = []struct {
		name         string
		v            int8
		expectedJSON string
	}{
		{
			name:         "basic",
			v:            int8(1),
			expectedJSON: `{"foo":1,"bar":1}`,
		},
		{
			name:         "big",
			v:            math.MaxInt8,
			expectedJSON: `{"foo":127,"bar":127}`,
		},
		{
			name:         "big",
			v:            int8(0),
			expectedJSON: `{}`,
		},
	}
	for _, testCase := range testCasesKeyOmitEmpty {
		t.Run(testCase.name, func(t *testing.T) {
			var b = &strings.Builder{}
			var enc = NewEncoder(b)
			enc.Encode(EncodeObjectFunc(func(enc *Encoder) {
				enc.Int8KeyOmitEmpty("foo", testCase.v)
				enc.AddInt8KeyOmitEmpty("bar", testCase.v)
			}))
			assert.Equal(t, testCase.expectedJSON, b.String())
		})
	}
}
