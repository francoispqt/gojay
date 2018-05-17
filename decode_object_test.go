package gojay

import (
	"io"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeObjectBasic(t *testing.T) {

}

func TestDecodeObjectComplex(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult testObjectComplex
		err            bool
		errType        interface{}
	}{
		{
			name: "basic",
			json: `{"testSubObject":{},"testSubSliceInts":[1,2]}`,
			expectedResult: testObjectComplex{
				testSubObject:    &testObject{},
				testSubSliceInts: &testSliceInts{1, 2},
			},
			err: false,
		},
		{
			name: "complex",
			json: `{"testSubObject":{"testStr":"some string","testInt":124465,"testUint16":120, "testUint8":15,"testInt16":-135,"testInt8":-23},"testSubSliceInts":[1,2],"testStr":"some \\n string"}`,
			expectedResult: testObjectComplex{
				testSubObject: &testObject{
					testStr:    "some string",
					testInt:    124465,
					testUint16: 120,
					testUint8:  15,
					testInt16:  -135,
					testInt8:   -23,
				},
				testSubSliceInts: &testSliceInts{1, 2},
				testStr:          "some \n string",
			},
			err: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := testObjectComplex{
				testSubObject:    &testObject{},
				testSubSliceInts: &testSliceInts{},
			}
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			log.Print(s, testCase.name)
			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, testCase.expectedResult, s, "value at given index should be the same as expected results")
		})
	}
}

type TestObj struct {
	test        int
	test2       int
	test3       string
	test4       string
	test5       float64
	testArr     testSliceObjects
	testSubObj  *TestSubObj
	testSubObj2 *TestSubObj
}

type TestSubObj struct {
	test3          int
	test4          int
	test5          string
	testSubSubObj  *TestSubObj
	testSubSubObj2 *TestSubObj
}

func (t *TestSubObj) UnmarshalJSONObject(dec *Decoder, key string) error {
	switch key {
	case "test":
		return dec.AddInt(&t.test3)
	case "test2":
		return dec.AddInt(&t.test4)
	case "test3":
		return dec.AddString(&t.test5)
	case "testSubSubObj":
		t.testSubSubObj = &TestSubObj{}
		return dec.AddObject(t.testSubSubObj)
	case "testSubSubObj2":
		t.testSubSubObj2 = &TestSubObj{}
		return dec.AddObject(t.testSubSubObj2)
	}
	return nil
}

func (t *TestSubObj) NKeys() int {
	return 0
}

func (t *TestObj) UnmarshalJSONObject(dec *Decoder, key string) error {
	switch key {
	case "test":
		return dec.AddInt(&t.test)
	case "test2":
		return dec.AddInt(&t.test2)
	case "test3":
		return dec.AddString(&t.test3)
	case "test4":
		return dec.AddString(&t.test4)
	case "test5":
		return dec.AddFloat(&t.test5)
	case "testSubObj":
		t.testSubObj = &TestSubObj{}
		return dec.AddObject(t.testSubObj)
	case "testSubObj2":
		t.testSubObj2 = &TestSubObj{}
		return dec.AddObject(t.testSubObj2)
	case "testArr":
		return dec.AddArray(&t.testArr)
	}
	return nil
}

func (t *TestObj) NKeys() int {
	return 8
}

func assertResult(t *testing.T, v *TestObj, err error) {
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 245, v.test, "v.test must be equal to 245")
	assert.Equal(t, 246, v.test2, "v.test2 must be equal to 246")
	assert.Equal(t, "string", v.test3, "v.test3 must be equal to 'string'")
	assert.Equal(t, "complex string with spaces and some slashes\"", v.test4, "v.test4 must be equal to 'string'")
	assert.Equal(t, -1.15657654376543, v.test5, "v.test5 must be equal to 1.15")
	assert.Len(t, v.testArr, 2, "v.testArr must be of len 2")
	// assert.Equal(t, v.testArr[0].test, 245, "v.testArr[0].test must be equal to 245")
	// assert.Equal(t, v.testArr[0].test2, 246, "v.testArr[0].test must be equal to 246")
	// assert.Equal(t, v.testArr[1].test, 245, "v.testArr[0].test must be equal to 245")
	// assert.Equal(t, v.testArr[1].test2, 246, "v.testArr[0].test must be equal to 246")

	assert.Equal(t, 121, v.testSubObj.test3, "v.testSubObj.test3 must be equal to 121")
	assert.Equal(t, 122, v.testSubObj.test4, "v.testSubObj.test4 must be equal to 122")
	assert.Equal(t, "string", v.testSubObj.test5, "v.testSubObj.test5 must be equal to 'string'")
	assert.Equal(t, 150, v.testSubObj.testSubSubObj.test3, "v.testSubObj.testSubSubObj.test3 must be equal to 150")
	assert.Equal(t, 150, v.testSubObj.testSubSubObj2.test3, "v.testSubObj.testSubSubObj2.test3 must be equal to 150")

	assert.Equal(t, 122, v.testSubObj2.test3, "v.testSubObj2.test3 must be equal to 121")
	assert.Equal(t, 123, v.testSubObj2.test4, "v.testSubObj2.test4 must be equal to 122")
	assert.Equal(t, "string", v.testSubObj2.test5, "v.testSubObj2.test5 must be equal to 'string'")
	assert.Equal(t, 151, v.testSubObj2.testSubSubObj.test3, "v.testSubObj2.testSubSubObj.test must be equal to 150")
}

func TestDecoderObject(t *testing.T) {
	json := []byte(`{
		"test": 245,
		"test2": 246,
		"test3": "string",
		"test4": "complex string with spaces and some slashes\"",
		"test5": -1.15657654376543,
		"testNull": null,
		"testArr": [
			{
				"test": 245,
				"test2": 246
			},
			{
				"test": 245,
				"test2": 246
			}
		],
		"testSubObj": {
			"test": 121,
			"test2": 122,
			"testNull": null,
			"testSubSubObj": {
				"test": 150,
				"testNull": null
			},
			"testSubSubObj2": {
				"test": 150
			},
			"test3": "string"
			"testNull": null,
		},
		"testSubObj2": {
			"test": 122,
			"test3": "string"
			"testSubSubObj": {
				"test": 151
			},
			"test2": 123
		}
	}`)
	v := &TestObj{}
	err := Unmarshal(json, v)
	assertResult(t, v, err)
}

func TestDecodeObjectNull(t *testing.T) {
	json := []byte(`null`)
	v := &TestObj{}
	err := Unmarshal(json, v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, v.test, 0, "v.test must be 0 val")
}

var jsonComplex = []byte(`{
	"test": "{\"test\":\"1\",\"test1\":2}",
	"test2\\n": "\\\\\\\\\\n",
	"testArrSkip": ["testString with escaped \\\" quotes"],
	"testSkipString": "skip \\ string with \\n escaped char \" ",
	"testSkipObject": {
		"testSkipSubObj": {
			"test": "test"
		}
	},
	"testSkipNumber": 123.23,
	"testSkipNumber2": 123.23 ,
	"testBool": true,
	"testSkipBoolTrue": true,
	"testSkipBoolFalse": false,
	"testSkipBoolNull": null,
	"testSub": {
		"test": "{\"test\":\"1\",\"test1\":2}",
		"test2\\n": "[1,2,3]",
		"test3": 1,
		"testObjSkip": {
			"test": "test string with escaped \" quotes"
		},
		"testStrSkip" : "test"
	},
	"testBoolSkip": false,
	"testObjInvalidType": "somestring",
	"testArrSkip2": [[],["someString"]],
	"test3": 1
}`)

type jsonObjectComplex struct {
	Test               string
	Test2              string
	Test3              int
	Test4              bool
	testSub            *jsonObjectComplex
	testObjInvalidType *jsonObjectComplex
}

func (j *jsonObjectComplex) UnmarshalJSONObject(dec *Decoder, key string) error {
	switch key {
	case "test":
		return dec.AddString(&j.Test)
	case "test2\n":
		return dec.AddString(&j.Test2)
	case "test3":
		return dec.AddInt(&j.Test3)
	case "testBool":
		return dec.AddBool(&j.Test4)
	case "testSub":
		j.testSub = &jsonObjectComplex{}
		return dec.AddObject(j.testSub)
	case "testObjInvalidType":
		j.testObjInvalidType = &jsonObjectComplex{}
		return dec.AddObject(j.testObjInvalidType)
	}
	return nil
}

func (j *jsonObjectComplex) NKeys() int {
	return 6
}

func TestDecodeObjComplex(t *testing.T) {
	result := jsonObjectComplex{}
	err := UnmarshalJSONObject(jsonComplex, &result)
	assert.NotNil(t, err, "err should not be as invalid type as been encountered nil")
	assert.Equal(t, `Cannot unmarshal to struct, wrong char '"' found at pos 639`, err.Error(), "err should not be as invalid type as been encountered nil")
	assert.Equal(t, `{"test":"1","test1":2}`, result.Test, "result.Test is not expected value")
	assert.Equal(t, "\\\\\\\\\n", result.Test2, "result.Test2 is not expected value")
	assert.Equal(t, 1, result.Test3, "result.test3 is not expected value")
	assert.Equal(t, `{"test":"1","test1":2}`, result.testSub.Test, "result.testSub.test is not expected value")
	assert.Equal(t, `[1,2,3]`, result.testSub.Test2, "result.testSub.test2 is not expected value")
	assert.Equal(t, 1, result.testSub.Test3, "result.testSub.test3 is not expected value")
	assert.Equal(t, true, result.Test4, "result.Test4 is not expected value, should be true")
}

type jsonDecodePartial struct {
	Test  string
	Test2 string
}

func (j *jsonDecodePartial) UnmarshalJSONObject(dec *Decoder, key string) error {
	switch key {
	case "test":
		return dec.AddString(&j.Test)
	case `test2`:
		return dec.AddString(&j.Test2)
	}
	return nil
}

func (j *jsonDecodePartial) NKeys() int {
	return 2
}

func TestDecodeObjectPartial(t *testing.T) {
	result := jsonDecodePartial{}
	dec := NewDecoder(nil)
	dec.data = []byte(`{
		"test": "test",
		"test2": "test",
		"testArrSkip": ["test"],
		"testSkipString": "test",
		"testSkipNumber": 123.23
	}`)
	dec.length = len(dec.data)
	err := dec.DecodeObject(&result)
	assert.Nil(t, err, "err should be nil")
	assert.NotEqual(t, len(dec.data), dec.cursor)
}

func TestDecoderObjectInvalidJSON(t *testing.T) {
	result := jsonDecodePartial{}
	dec := NewDecoder(nil)
	dec.data = []byte(`{
		"test2": "test",
		"testArrSkip": ["test"],
		"testSkipString": "testInvalidJSON\\\\
	}`)
	dec.length = len(dec.data)
	err := dec.DecodeObject(&result)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

type myMap map[string]string

func (m myMap) UnmarshalJSONObject(dec *Decoder, k string) error {
	str := ""
	err := dec.AddString(&str)
	if err != nil {
		return err
	}
	m[k] = str
	return nil
}

// return 0 to parse all keys
func (m myMap) NKeys() int {
	return 0
}

func TestDecoderObjectMap(t *testing.T) {
	json := `{
		"test": "string",
		"test2": "string",
		"test3": "string",
		"test4": "string",
		"test5": "string",
	}`
	m := myMap(make(map[string]string))
	dec := BorrowDecoder(strings.NewReader(json))
	err := dec.Decode(m)

	assert.Nil(t, err, "err should be nil")
	assert.Len(t, m, 5, "len of m should be 5")
}

func TestDecoderObjectDecoderAPI(t *testing.T) {
	json := `{
		"test": 245,
		"test2": 246,
		"test3": "string",
		"test4": "complex string with spaces and some slashes\"",
		"test5": -1.15657654376543,
		"testNull": null,
		"testArr": [
			{
				"test": 245,
				"test2": 246
			},
			{
				"test": 245,
				"test2": 246
			}
		],
		"testSubObj": {
			"test": 121,
			"test2": 122,
			"testNull": null,
			"testSubSubObj": {
				"test": 150,
				"testNull": null
			},
			"testSubSubObj2": {
				"test": 150
			},
			"test3": "string"
			"testNull": null,
		},
		"testSubObj2": {
			"test": 122,
			"test3": "string"
			"testSubSubObj": {
				"test": 151
			},
			"test2": 123
		}
	}`
	v := &TestObj{}
	dec := NewDecoder(strings.NewReader(json))
	err := dec.DecodeObject(v)
	assertResult(t, v, err)
}

type ReadCloser struct {
	json []byte
}

func (r *ReadCloser) Read(b []byte) (int, error) {
	copy(b, r.json)
	return len(r.json), io.EOF
}

func TestDecoderObjectDecoderAPIReadCloser(t *testing.T) {
	readCloser := ReadCloser{
		json: []byte(`{
			"test": "string",
			"test2": "string",
			"test3": "string",
			"test4": "string",
			"test5": "string",
		}`),
	}
	m := myMap(make(map[string]string))
	dec := NewDecoder(&readCloser)
	err := dec.DecodeObject(m)
	assert.Nil(t, err, "err should be nil")
	assert.Len(t, m, 5, "len of m should be 5")
}

func TestDecoderObjectDecoderAPIFuncReadCloser(t *testing.T) {
	readCloser := ReadCloser{
		json: []byte(`{
			"test": "string",
			"test2": "string",
			"test3": "string",
			"test4": "string",
			"test5": "string",
		}`),
	}
	m := myMap(make(map[string]string))
	dec := NewDecoder(&readCloser)
	err := dec.DecodeObject(DecodeObjectFunc(func(dec *Decoder, k string) error {
		str := ""
		err := dec.AddString(&str)
		if err != nil {
			return err
		}
		m[k] = str
		return nil
	}))
	assert.Nil(t, err, "err should be nil")
	assert.Len(t, m, 5, "len of m should be 5")
}

func TestDecoderObjectDecoderInvalidJSONError(t *testing.T) {
	v := &TestObj{}
	dec := NewDecoder(strings.NewReader(`{"err:}`))
	err := dec.DecodeObject(v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderObjectDecoderInvalidJSONError2(t *testing.T) {
	v := &TestSubObj{}
	dec := NewDecoder(strings.NewReader(`{"err:}`))
	err := dec.DecodeObject(v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderObjectDecoderInvalidJSONError3(t *testing.T) {
	v := &TestSubObj{}
	dec := NewDecoder(strings.NewReader(`{"err":"test}`))
	err := dec.DecodeObject(v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderObjectDecoderInvalidJSONError4(t *testing.T) {
	testArr := testSliceInts{}
	dec := NewDecoder(strings.NewReader(`hello`))
	err := dec.DecodeArray(&testArr)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderObjectPoolError(t *testing.T) {
	result := jsonDecodePartial{}
	dec := NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = dec.DecodeObject(&result)
	assert.True(t, false, "should not be called as decoder should have panicked")
}

func TestSkipData(t *testing.T) {
	testCases := []struct {
		name string
		err  bool
		json string
	}{
		{
			name: "skip-bool-false-err",
			json: `fulse`,
			err:  true,
		},
		{
			name: "skip-bool-true-err",
			json: `trou`,
			err:  true,
		},
		{
			name: "skip-bool-null-err",
			json: `nil`,
			err:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.skipData()
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
			} else {
				assert.Nil(t, err, "err should be nil")
			}
		})
	}
	t.Run("error-invalid-json", func(t *testing.T) {
		dec := NewDecoder(strings.NewReader(""))
		err := dec.skipData()
		assert.NotNil(t, err, "err should not be nil as data is empty")
		assert.IsType(t, InvalidJSONError(""), err, "err should of type InvalidJSONError")
	})
	t.Run("skip-array-error-invalid-json", func(t *testing.T) {
		dec := NewDecoder(strings.NewReader(""))
		_, err := dec.skipArray()
		assert.NotNil(t, err, "err should not be nil as data is empty")
		assert.IsType(t, InvalidJSONError(""), err, "err should of type InvalidJSONError")
	})
}
