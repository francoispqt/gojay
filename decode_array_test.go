package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSliceStrings []string

func (t *testSliceStrings) UnmarshalArray(dec *Decoder) error {
	str := ""
	if err := dec.AddString(&str); err != nil {
		return err
	}
	*t = append(*t, str)
	return nil
}

type testSliceInts []*int

func (t *testSliceInts) UnmarshalArray(dec *Decoder) error {
	i := 0
	ptr := &i
	*t = append(*t, ptr)
	return dec.AddInt(ptr)
}

type testSliceObj []*TestObj

func (t *testSliceObj) UnmarshalArray(dec *Decoder) error {
	obj := &TestObj{}
	*t = append(*t, obj)
	return dec.AddObject(obj)
}

type testChannelArray chan *TestObj

func (c *testChannelArray) UnmarshalArray(dec *Decoder) error {
	obj := &TestObj{}
	if err := dec.AddObject(obj); err != nil {
		return err
	}
	*c <- obj
	return nil
}

func TestDecoderSliceOfStringsBasic(t *testing.T) {
	json := []byte(`["string","string1"]`)
	testArr := testSliceStrings{}
	err := Unmarshal(json, &testArr)
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, "string", testArr[0], "testArr[0] should be 'string'")
	assert.Equal(t, "string1", testArr[1], "testArr[1] should be 'string1'")
}

func TestDecoderSliceNull(t *testing.T) {
	json := []byte(`null`)
	v := &testSliceStrings{}
	err := Unmarshal(json, v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, len(*v), 0, "v must be of len 0")
}

func TestDecoderSliceArrayOfIntsBasic(t *testing.T) {
	json := []byte(`[
		1,
		2
	]`)
	testArr := testSliceInts{}
	err := UnmarshalArray(json, &testArr)
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, 1, *testArr[0], "testArr[0] should be 1")
	assert.Equal(t, 2, *testArr[1], "testArr[1] should be 2")
}

func TestDecoderSliceArrayOfIntsBigInts(t *testing.T) {
	json := []byte(`[
		789034384533530523,
		545344023293232032
	]`)
	testArr := testSliceInts{}
	err := UnmarshalArray(json, &testArr)
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, 789034384533530523, *testArr[0], "testArr[0] should be 789034384533530523")
	assert.Equal(t, 545344023293232032, *testArr[1], "testArr[1] should be 545344023293232032")
}

func TestDecoderSliceOfObjectsBasic(t *testing.T) {
	json := []byte(`[
		{
			"test": 245,
			"test2": -246,
			"test3": "string"
		},
		{
			"test": 247,
			"test2": 248,
			"test3": "string"
		},
		{
			"test": 777,
			"test2": 456,
			"test3": "string"
		}
	]`)
	testArr := testSliceObj{}
	err := Unmarshal(json, &testArr)
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 3, "testArr should be of len 2")
	assert.Equal(t, 245, testArr[0].test, "testArr[0] should be 245")
	assert.Equal(t, -246, testArr[0].test2, "testArr[0] should be 246")
	assert.Equal(t, "string", testArr[0].test3, "testArr[0].test3 should be 'string'")
	assert.Equal(t, 247, testArr[1].test, "testArr[1] should be 247")
	assert.Equal(t, 248, testArr[1].test2, "testArr[1] should be 248")
	assert.Equal(t, "string", testArr[1].test3, "testArr[1].test3 should be 'string'")
	assert.Equal(t, 777, testArr[2].test, "testArr[2] should be 777")
	assert.Equal(t, 456, testArr[2].test2, "testArr[2] should be 456")
	assert.Equal(t, "string", testArr[2].test3, "testArr[2].test3 should be 'string'")
}

func TestDecodeSliceInvalidType(t *testing.T) {
	result := testSliceObj{}
	err := UnmarshalArray([]byte(`{}`), &result)
	assert.NotNil(t, err, "err should not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err should be of type InvalidTypeError")
	assert.Equal(t, "Cannot unmarshall to array, wrong char '{' found at pos 0", err.Error(), "err should not be nil")
}

func TestDecoderChannelOfObjectsBasic(t *testing.T) {
	json := []byte(`[
		{
			"test": 245,
			"test2": -246,
			"test3": "string"
		},
		{
			"test": 247,
			"test2": 248,
			"test3": "string"
		},
		{
			"test": 777,
			"test2": 456,
			"test3": "string"
		}
	]`)
	testChan := testChannelArray(make(chan *TestObj, 3))
	err := UnmarshalArray(json, &testChan)
	assert.Nil(t, err, "Err must be nil")
	ct := 0
	l := len(testChan)
	for _ = range testChan {
		ct++
		if ct == l {
			break
		}
	}
	assert.Equal(t, ct, 3)
}

func TestDecoderSliceInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	testArr := testSliceInts{}
	err := UnmarshalArray(json, &testArr)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}
