package gojay

import (
	"strings"
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

type testObjectWithUnknownType struct {
	unknownType struct{}
}

func (t *testObjectWithUnknownType) IsNil() bool {
	return t == nil
}

func (t *testObjectWithUnknownType) MarshalObject(enc *Encoder) {
	enc.AddInterfaceKey("unknownType", t.unknownType)
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
	enc.AddInterfaceKey("testArr", t.testArr)
	enc.AddInterfaceKey("testF64", t.testF64)
	enc.AddInterfaceKey("testF32", t.testF32)
	enc.AddInterfaceKey("testInterface", t.testInterface)
	enc.AddInterfaceKey("sub", t.sub)
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

type testEncodingObjInterfaces struct {
	interfaceVal interface{}
}

func (t *testEncodingObjInterfaces) IsNil() bool {
	return t == nil
}

func (t *testEncodingObjInterfaces) MarshalObject(enc *Encoder) {
	enc.AddInterfaceKey("interfaceVal", t.interfaceVal)
}

func TestEncoderObjectEncodeAPI(t *testing.T) {
	t.Run("encode-basic", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeObject(&testObject{"漢字", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.1, 1.1, true})
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"testStr":"漢字","testInt":1,"testInt64":1,"testInt32":1,"testInt16":1,"testInt8":1,"testUint64":1,"testUint32":1,"testUint16":1,"testUint8":1,"testFloat64":1.1,"testFloat32":1.1,"testBool":true}`,
			builder.String(),
			"Result of marshalling is different as the one expected",
		)
	})
}

func TestEncoderObjectMarshalAPI(t *testing.T) {
	t.Run("marshal-basic", func(t *testing.T) {
		r, err := Marshal(&testObject{"漢字", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.1, 1.1, true})
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"testStr":"漢字","testInt":1,"testInt64":1,"testInt32":1,"testInt16":1,"testInt8":1,"testUint64":1,"testUint32":1,"testUint16":1,"testUint8":1,"testFloat64":1.1,"testFloat32":1.1,"testBool":true}`,
			string(r),
			"Result of marshalling is different as the one expected",
		)
	})

	t.Run("marshal-complex", func(t *testing.T) {
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
			`{"test":"hello world","test2":"foobar","testInt":1,"testBool":true,"testArr":[{"test":"1","test2":"","testInt":0,"testBool":false,"testArr":[],"testF64":0,"testF32":0,"sub":{}}],"testF64":120.15,"testF32":120.53,"testInterface":true,"sub":{"test1":10,"test2":"hello world","test3":1.23543,"testBool":true,"sub":{"test1":10,"test2":"hello world","test3":0,"testBool":false,"sub":{}}}}`,
			string(r),
			"Result of marshalling is different as the one expected",
		)
	})

	t.Run("marshal-interface-string", func(t *testing.T) {
		v := testEncodingObjInterfaces{"string"}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":"string"}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-int", func(t *testing.T) {
		v := testEncodingObjInterfaces{1}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-int64", func(t *testing.T) {
		v := testEncodingObjInterfaces{int64(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-int32", func(t *testing.T) {
		v := testEncodingObjInterfaces{int32(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-int16", func(t *testing.T) {
		v := testEncodingObjInterfaces{int16(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-int8", func(t *testing.T) {
		v := testEncodingObjInterfaces{int8(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-uint64", func(t *testing.T) {
		v := testEncodingObjInterfaces{uint64(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-uint32", func(t *testing.T) {
		v := testEncodingObjInterfaces{uint32(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-uint16", func(t *testing.T) {
		v := testEncodingObjInterfaces{uint16(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-uint8", func(t *testing.T) {
		v := testEncodingObjInterfaces{uint8(1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-float64", func(t *testing.T) {
		v := testEncodingObjInterfaces{float64(1.1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1.1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
	t.Run("marshal-interface-float32", func(t *testing.T) {
		v := testEncodingObjInterfaces{float32(1.1)}
		r, err := Marshal(&v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"interfaceVal":1.1}`,
			string(r),
			"Result of marshalling is different as the one expected")
	})
}

type TestObectOmitEmpty struct {
	nonNiler           int
	testInt            int
	testFloat          float64
	testFloat32        float32
	testString         string
	testBool           bool
	testObectOmitEmpty *TestObectOmitEmpty
	testObect          *TestObectOmitEmpty
}

func (t *TestObectOmitEmpty) IsNil() bool {
	return t == nil
}

func (t *TestObectOmitEmpty) MarshalObject(enc *Encoder) {
	enc.AddIntKeyOmitEmpty("testInt", t.testInt)
	enc.AddIntKeyOmitEmpty("testIntNotEmpty", 1)
	enc.AddFloatKeyOmitEmpty("testFloat", t.testFloat)
	enc.AddFloatKeyOmitEmpty("testFloatNotEmpty", 1.1)
	enc.AddFloat32KeyOmitEmpty("testFloat32", t.testFloat32)
	enc.AddFloat32KeyOmitEmpty("testFloat32NotEmpty", 1.1)
	enc.AddStringKeyOmitEmpty("testString", t.testString)
	enc.AddStringKeyOmitEmpty("testStringNotEmpty", "foo")
	enc.AddBoolKeyOmitEmpty("testBool", t.testBool)
	enc.AddBoolKeyOmitEmpty("testBoolNotEmpty", true)
	enc.AddObjectKeyOmitEmpty("testObect", t.testObect)
	enc.AddObjectKeyOmitEmpty("testObectOmitEmpty", t.testObectOmitEmpty)
	enc.AddArrayKeyOmitEmpty("testArrayOmitEmpty", TestEncodingArrStrings{})
	enc.AddArrayKeyOmitEmpty("testArray", TestEncodingArrStrings{"foo"})
}

type TestObectOmitEmptyInterface struct{}

func (t *TestObectOmitEmptyInterface) IsNil() bool {
	return t == nil
}

func (t *TestObectOmitEmptyInterface) MarshalObject(enc *Encoder) {
	enc.AddInterfaceKeyOmitEmpty("testInt", 0)
	enc.AddInterfaceKeyOmitEmpty("testInt64", int64(0))
	enc.AddInterfaceKeyOmitEmpty("testInt32", int32(0))
	enc.AddInterfaceKeyOmitEmpty("testInt16", int16(0))
	enc.AddInterfaceKeyOmitEmpty("testInt8", int8(0))
	enc.AddInterfaceKeyOmitEmpty("testUint8", uint8(0))
	enc.AddInterfaceKeyOmitEmpty("testUint16", uint16(0))
	enc.AddInterfaceKeyOmitEmpty("testUint32", uint32(0))
	enc.AddInterfaceKeyOmitEmpty("testUint64", uint64(0))
	enc.AddInterfaceKeyOmitEmpty("testIntNotEmpty", 1)
	enc.AddInterfaceKeyOmitEmpty("testFloat", 0)
	enc.AddInterfaceKeyOmitEmpty("testFloatNotEmpty", 1.1)
	enc.AddInterfaceKeyOmitEmpty("testFloat32", float32(0))
	enc.AddInterfaceKeyOmitEmpty("testFloat32NotEmpty", float32(1.1))
	enc.AddInterfaceKeyOmitEmpty("testString", "")
	enc.AddInterfaceKeyOmitEmpty("testStringNotEmpty", "foo")
	enc.AddInterfaceKeyOmitEmpty("testBool", false)
	enc.AddInterfaceKeyOmitEmpty("testBoolNotEmpty", true)
	enc.AddInterfaceKeyOmitEmpty("testObectOmitEmpty", nil)
	enc.AddInterfaceKeyOmitEmpty("testObect", &TestEncoding{})
	enc.AddInterfaceKeyOmitEmpty("testArr", &TestEncodingArrStrings{})
}

func TestEncoderObjectOmitEmpty(t *testing.T) {
	t.Run("encoder-omit-empty-all-types", func(t *testing.T) {
		v := &TestObectOmitEmpty{
			nonNiler:  1,
			testInt:   0,
			testObect: &TestObectOmitEmpty{testInt: 1},
		}
		r, err := MarshalObject(v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"testIntNotEmpty":1,"testFloatNotEmpty":1.1,"testFloat32NotEmpty":1.1,"testStringNotEmpty":"foo","testBoolNotEmpty":true,"testObect":{"testInt":1,"testIntNotEmpty":1,"testFloatNotEmpty":1.1,"testFloat32NotEmpty":1.1,"testStringNotEmpty":"foo","testBoolNotEmpty":true,"testArray":["foo"]},"testArray":["foo"]}`,
			string(r),
			"Result of marshalling is different as the one expected",
		)
	})

	t.Run("encoder-omit-empty-interface", func(t *testing.T) {
		v := &TestObectOmitEmptyInterface{}
		r, err := MarshalObject(v)
		assert.Nil(t, err, "Error should be nil")
		assert.Equal(
			t,
			`{"testIntNotEmpty":1,"testFloatNotEmpty":1.1,"testFloat32NotEmpty":1.1,"testStringNotEmpty":"foo","testBoolNotEmpty":true,"testObect":{"test":"","test2":"","testInt":0,"testBool":false,"testArr":[],"testF64":0,"testF32":0,"sub":{}}}`,
			string(r),
			"Result of marshalling is different as the one expected",
		)
	})
}

func TestEncoderObjectEncodeAPIError(t *testing.T) {
	t.Run("interface-key-error", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		err := enc.EncodeObject(&testObjectWithUnknownType{struct{}{}})
		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, "Invalid type struct {} provided to Marshal", err.Error(), "err.Error() should be 'Invalid type struct {} provided to Marshal'")
	})
	t.Run("write-error", func(t *testing.T) {
		w := TestWriterError("")
		enc := NewEncoder(w)
		err := enc.EncodeObject(&testObject{"漢字", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1.1, 1.1, true})
		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, "Test Error", err.Error(), "err.Error() should be 'Test Error'")
	})
	t.Run("interface-error", func(t *testing.T) {
		builder := &strings.Builder{}
		enc := NewEncoder(builder)
		enc.AddInterfaceKeyOmitEmpty("test", struct{}{})
		assert.NotNil(t, enc.err, "enc.Err() should not be nil")
	})
	t.Run("pool-error", func(t *testing.T) {
		v := &TestEncoding{}
		enc := BorrowEncoder(nil)
		enc.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledEncoderError(""), err, "err should be of type InvalidUsagePooledEncoderError")
			assert.Equal(t, "Invalid usage of pooled encoder", err.(InvalidUsagePooledEncoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = enc.EncodeObject(v)
		assert.True(t, false, "should not be called as it should have panicked")
	})
}
