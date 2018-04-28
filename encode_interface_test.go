package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var encoderTestCases = []struct {
	v            interface{}
	expectations func(t *testing.T, b []byte, err error)
}{
	{
		v: 100,
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: int64(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: int32(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: int8(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: uint64(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: uint32(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: uint16(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: uint8(100),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100", string(b), "string(b) should equal 100")
		},
	},
	{
		v: float64(100.12),
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "100.12", string(b), "string(b) should equal 100.12")
		},
	},
	{
		v: true,
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, "true", string(b), "string(b) should equal true")
		},
	},
	{
		v: "hello world",
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, `"hello world"`, string(b), `string(b) should equal "hello world"`)
		},
	},
	{
		v: "hello world",
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, `"hello world"`, string(b), `string(b) should equal "hello world"`)
		},
	},
	{
		v: &TestEncodingArrStrings{"hello world", "foo bar"},
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, `["hello world","foo bar"]`, string(b), `string(b) should equal ["hello world","foo bar"]`)
		},
	},
	{
		v: &testObject{"漢字", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.1, 1.1, true},
		expectations: func(t *testing.T, b []byte, err error) {
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, `{"testStr":"漢字","testInt":1,"testInt64":1,"testInt32":1,"testInt16":1,"testInt8":1,"testUint64":1,"testUint32":1,"testUint16":1,"testUint8":1,"testFloat64":1.1,"testFloat32":1.1,"testBool":true}`, string(b), `string(b) should equal {"testStr":"漢字","testInt":1,"testInt64":1,"testInt32":1,"testInt16":1,"testInt8":1,"testUint64":1,"testUint32":1,"testUint16":1,"testUint8":1,"testFloat64":1.1,"testFloat32":1.1,"testBool":true}`)
		},
	},
	{
		v: &struct{}{},
		expectations: func(t *testing.T, b []byte, err error) {
			assert.NotNil(t, err, "err should be nil")
			assert.IsType(t, InvalidMarshalError(""), err, "err should be of type InvalidMarshalError")
		},
	},
}

func TestEncoderInterfaceAllTypesDecoderAPI(t *testing.T) {
	for _, test := range encoderTestCases {
		enc := BorrowEncoder()
		b, err := enc.Encode(test.v)
		enc.Release()
		test.expectations(t, b, err)
	}
}

func TestEncoderInterfaceAllTypesMarshalAPI(t *testing.T) {
	for _, test := range encoderTestCases {
		b, err := Marshal(test.v)
		test.expectations(t, b, err)
	}
}
