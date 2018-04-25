package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	testStr     string
	testInt     int
	testInt64   int64
	testInt32   int32
	testInt16   int16
	testInt8    int8
	testUint64  uint64
	testUint32  uint32
	testUint16  uint16
	testUint8   uint8
	testFloat64 float64
	testFloat32 float32
	testBool    bool
}

func (t *testObject) IsNil() bool {
	return t == nil
}

func (t *testObject) MarshalObject(enc *Encoder) {
	enc.AddStringKey("testStr", t.testStr)
	enc.AddIntKey("testInt", t.testInt)
	enc.AddIntKey("testInt64", int(t.testInt64))
	enc.AddIntKey("testInt32", int(t.testInt32))
	enc.AddIntKey("testInt16", int(t.testInt16))
	enc.AddIntKey("testInt8", int(t.testInt8))
	enc.AddIntKey("testUint64", int(t.testUint64))
	enc.AddIntKey("testUint32", int(t.testUint32))
	enc.AddIntKey("testUint16", int(t.testUint16))
	enc.AddIntKey("testUint8", int(t.testUint8))
	enc.AddFloatKey("testFloat64", t.testFloat64)
	enc.AddFloat32Key("testFloat32", t.testFloat32)
	enc.AddBoolKey("testBool", t.testBool)
}

func TestEncodeBasicObject(t *testing.T) {
	r, err := Marshal(&testObject{"漢字", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.1, 1.1, true})
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"testStr":"漢字","testInt":1,"testInt64":1,"testInt32":1,"testInt16":1,"testInt8":1,"testUint64":1,"testUint32":1,"testUint16":1,"testUint8":1,"testFloat64":1.1,"testFloat32":1.1,"testBool":true}`,
		string(r),
		"Result of marshalling is different as the one expected",
	)
}

type TestEncoding struct {
	test          string
	test2         string
	testInt       int
	testBool      bool
	testF32       float32
	testF64       float64
	testInterface interface{}
	testArr       TestEncodingArr
	sub           *SubObject
}

func (t *TestEncoding) IsNil() bool {
	return t == nil
}

func (t *TestEncoding) MarshalObject(enc *Encoder) {
	enc.AddStringKey("test", t.test)
	enc.AddStringKey("test2", t.test2)
	enc.AddIntKey("testInt", t.testInt)
	enc.AddBoolKey("testBool", t.testBool)
	enc.AddArrayKey("testArr", t.testArr)
	enc.AddInterfaceKey("testF64", t.testF64)
	enc.AddInterfaceKey("testF32", t.testF32)
	enc.AddInterfaceKey("testInterface", t.testInterface)
	enc.AddObjectKey("sub", t.sub)
}

type SubObject struct {
	test1    int
	test2    string
	test3    float64
	testBool bool
	sub      *SubObject
}

func (t *SubObject) IsNil() bool {
	return t == nil
}

func (t *SubObject) MarshalObject(enc *Encoder) {
	enc.AddIntKey("test1", t.test1)
	enc.AddStringKey("test2", t.test2)
	enc.AddFloatKey("test3", t.test3)
	enc.AddBoolKey("testBool", t.testBool)
	enc.AddObjectKey("sub", t.sub)
}

func TestEncoderComplexObject(t *testing.T) {
	v := &TestEncoding{
		test:          "hello world",
		test2:         "foobar",
		testInt:       1,
		testBool:      true,
		testF32:       120.53,
		testF64:       120.15,
		testInterface: true,
		testArr: TestEncodingArr{
			&TestEncoding{
				test: "1",
			},
		},
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
	}
	r, err := MarshalObject(v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"test":"hello world","test2":"foobar","testInt":1,"testBool":true,"testArr":[{"test":"1","test2":"","testInt":0,"testBool":false,"testArr":[],"testF64":0,"testF32":0}],"testF64":120.15,"testF32":120.53,"testInterface":true,"sub":{"test1":10,"test2":"hello world","test3":1.23543,"testBool":true,"sub":{"test1":10,"test2":"hello world","test3":0,"testBool":false}}}`,
		string(r),
		"Result of marshalling is different as the one expected",
	)
}

type testEncodingObjInterfaces struct {
	interfaceVal interface{}
}

func (t *testEncodingObjInterfaces) IsNil() bool {
	return t == nil
}

func (t *testEncodingObjInterfaces) MarshalObject(enc *Encoder) {
	enc.AddInterfaceKey("interfaceVal", t.interfaceVal)
}

func TestObjInterfaces(t *testing.T) {
	v := testEncodingObjInterfaces{"string"}
	r, err := Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":"string"}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{1}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{int64(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{int32(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{int16(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{int8(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{uint64(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{uint32(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{uint16(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{uint8(1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{float64(1.1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1.1}`,
		string(r),
		"Result of marshalling is different as the one expected")
	v = testEncodingObjInterfaces{float32(1.1)}
	r, err = Marshal(&v)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(
		t,
		`{"interfaceVal":1.1}`,
		string(r),
		"Result of marshalling is different as the one expected")
}
