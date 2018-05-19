package gojay

type testObject struct {
	testStr       string
	testInt       int
	testInt64     int64
	testInt32     int32
	testInt16     int16
	testInt8      int8
	testUint64    uint64
	testUint32    uint32
	testUint16    uint16
	testUint8     uint8
	testFloat64   float64
	testFloat32   float32
	testBool      bool
	testSubObject *testObject
	testSubArray  testSliceInts
}

// make sure it implements interfaces
var _ MarshalerJSONObject = &testObject{}
var _ UnmarshalerJSONObject = &testObject{}

func (t *testObject) IsNil() bool {
	return t == nil
}

func (t *testObject) MarshalJSONObject(enc *Encoder) {
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

func (t *testObject) UnmarshalJSONObject(dec *Decoder, k string) error {
	switch k {
	case "testStr":
		return dec.AddString(&t.testStr)
	case "testInt":
		return dec.AddInt(&t.testInt)
	case "testInt64":
		return dec.AddInt64(&t.testInt64)
	case "testInt32":
		return dec.AddInt32(&t.testInt32)
	case "testInt16":
		return dec.AddInt16(&t.testInt16)
	case "testInt8":
		return dec.AddInt8(&t.testInt8)
	case "testUint64":
		return dec.AddUint64(&t.testUint64)
	case "testUint32":
		return dec.AddUint32(&t.testUint32)
	case "testUint16":
		return dec.AddUint16(&t.testUint16)
	case "testUint8":
		return dec.AddUint8(&t.testUint8)
	case "testFloat64":
		return dec.AddFloat(&t.testFloat64)
	case "testFloat32":
		return dec.AddFloat32(&t.testFloat32)
	case "testBool":
		return dec.AddBool(&t.testBool)
	}
	return nil
}

func (t *testObject) NKeys() int {
	return 13
}

type testObjectComplex struct {
	testSubObject    *testObject
	testSubSliceInts *testSliceInts
	testStr          string
	testSubObject2   *testObjectComplex
}

func (t *testObjectComplex) IsNil() bool {
	return t == nil
}

func (t *testObjectComplex) MarshalJSONObject(enc *Encoder) {
	enc.AddObjectKey("testSubObject", t.testSubObject)
	enc.AddStringKey("testStr", t.testStr)
	enc.AddObjectKey("testStr", t.testSubObject2)
}

func (t *testObjectComplex) UnmarshalJSONObject(dec *Decoder, k string) error {
	switch k {
	case "testSubObject":
		return dec.AddObject(t.testSubObject)
	case "testSubSliceInts":
		return dec.AddArray(t.testSubSliceInts)
	case "testStr":
		return dec.AddString(&t.testStr)
	case "testSubObject2":
		return dec.AddObject(t.testSubObject2)
	}
	return nil
}

func (t *testObjectComplex) NKeys() int {
	return 4
}

// make sure it implements interfaces
var _ MarshalerJSONObject = &testObjectComplex{}
var _ UnmarshalerJSONObject = &testObjectComplex{}
