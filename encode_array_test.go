package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEncodingArrStrings []string

func (t TestEncodingArrStrings) MarshalArray(enc *Encoder) {
	for _, e := range t {
		enc.AddString(e)
	}
}

type TestEncodingArr []*TestEncoding

func (t TestEncodingArr) MarshalArray(enc *Encoder) {
	for _, e := range t {
		enc.AddObject(e)
	}
}
func TestEncoderArrayObjects(t *testing.T) {
	v := &TestEncodingArr{
		&TestEncoding{
			test:          "hello world",
			test2:         "漢字",
			testInt:       1,
			testBool:      true,
			testInterface: 1,
			sub: &SubObject{
				test1:    10,
				test2:    "hello world",
				test3:    1.23543,
				testBool: true,
				sub: &SubObject{
					test1:    10,
					testBool: false,
					test2:    "hello world",
				},
			},
		},
		&TestEncoding{
			test:     "hello world",
			test2:    "漢字",
			testInt:  1,
			testBool: true,
			sub: &SubObject{
				test1:    10,
				test2:    "hello world",
				test3:    1.23543,
				testBool: true,
				sub: &SubObject{
					test1:    10,
					testBool: false,
					test2:    "hello world",
				},
			},
		},
	}
	r, err := Marshal(v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`[{"test":"hello world","test2":"漢字","testInt":1,"testBool":true,`+
			`"testArr":[],"testF64":0,"testF32":0,"testInterface":1,"sub":{"test1":10,"test2":"hello world",`+
			`"test3":1.23543,"testBool":true,"sub":{"test1":10,"test2":"hello world",`+
			`"test3":0,"testBool":false}}},{"test":"hello world","test2":"漢字","testInt":1,`+
			`"testBool":true,"testArr":[],"testF64":0,"testF32":0,"sub":{"test1":10,"test2":"hello world","test3":1.23543,`+
			`"testBool":true,"sub":{"test1":10,"test2":"hello world","test3":0,"testBool":false}}}]`,
		string(r),
		"Result of marshalling is different as the one expected")
}

type testEncodingArrInterfaces []interface{}

func (t testEncodingArrInterfaces) MarshalArray(enc *Encoder) {
	for _, e := range t {
		enc.AddInterface(e)
	}
}

func TestEncoderArrayInterfaces(t *testing.T) {
	v := &testEncodingArrInterfaces{
		1,
		int64(1),
		int32(1),
		int16(1),
		int8(1),
		uint64(1),
		uint32(1),
		uint16(1),
		uint8(1),
		float64(1.31),
		// float32(1.31),
		&TestEncodingArr{},
		true,
		"test",
		&TestEncoding{
			test:     "hello world",
			test2:    "foobar",
			testInt:  1,
			testBool: true,
		},
	}
	r, err := MarshalArray(v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`[1,1,1,1,1,1,1,1,1.31,[],true,"test",{"test":"hello world","test2":"foobar","testInt":1,"testBool":true,"testArr":[],"testF64":0,"testF32":0}]`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderArrayInterfacesEncoderAPI(t *testing.T) {
	v := &testEncodingArrInterfaces{
		1,
		int64(1),
		int32(1),
		int16(1),
		int8(1),
		uint64(1),
		uint32(1),
		uint16(1),
		uint8(1),
		float64(1.31),
		// float32(1.31),
		&TestEncodingArr{},
		true,
		"test",
		&TestEncoding{
			test:     "hello world",
			test2:    "foobar",
			testInt:  1,
			testBool: true,
		},
	}
	enc := BorrowEncoder()
	defer enc.Release()
	r, err := enc.EncodeArray(v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`[1,1,1,1,1,1,1,1,1.31,[],true,"test",{"test":"hello world","test2":"foobar","testInt":1,"testBool":true,"testArr":[],"testF64":0,"testF32":0}]`,
		string(r),
		"Result of marshalling is different as the one expected")
}

func TestEncoderArrayPooledError(t *testing.T) {
	v := &testEncodingArrInterfaces{}
	enc := BorrowEncoder()
	enc.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
		assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
	}()
	_, _ = enc.EncodeArray(v)
	assert.True(t, false, "should not be called as it should have panicked")
}
