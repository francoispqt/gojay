package gojay

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderStringBasic(t *testing.T) {
	json := []byte(`"string"`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "string", v, "v must be equal to 'string'")
}

func TestDecoderStringEmpty(t *testing.T) {
	json := []byte(``)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to 'string'")
}

func TestDecoderStringNullInvalid(t *testing.T) {
	json := []byte(`nall`)
	var v string
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to 'string'")
}

func TestDecoderStringComplex(t *testing.T) {
	json := []byte(`  "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "string with spaces and \"escape\"d \"quotes\" and escaped line returns \n and escaped \\\\ escaped char", v, "v is not equal to the value expected")
}

func TestDecoderStringNull(t *testing.T) {
	json := []byte(`null`)
	var v string
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to ''")
}

func TestDecoderStringInvalidJSON(t *testing.T) {
	json := []byte(`"invalid JSONs`)
	var v string
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderStringInvalidType(t *testing.T) {
	json := []byte(`1`)
	var v string
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidTypeError(""), err, "err message must be 'Invalid JSON'")
}

func TestDecoderStringDecoderAPI(t *testing.T) {
	var v string
	dec := NewDecoder(strings.NewReader(`"hello world!"`))
	defer dec.Release()
	err := dec.DecodeString(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "hello world!", v, "v must be equal to 'hello world!'")
}

func TestDecoderStringPoolError(t *testing.T) {
	// reset the pool to make sure it's not full
	decPool = sync.Pool{
		New: func() interface{} {
			return NewDecoder(nil)
		},
	}
	result := ""
	dec := NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = dec.DecodeString(&result)
	assert.True(t, false, "should not be called as decoder should have panicked")
}

func TestDecoderSkipEscapedStringError(t *testing.T) {
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.skipEscapedString()
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}

func TestDecoderSkipEscapedStringError2(t *testing.T) {
	dec := NewDecoder(strings.NewReader(`\"`))
	defer dec.Release()
	err := dec.skipEscapedString()
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}

func TestDecoderSkipEscapedStringError3(t *testing.T) {
	dec := NewDecoder(strings.NewReader(`invalid`))
	defer dec.Release()
	err := dec.skipEscapedString()
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}

func TestDecoderSkipStringError(t *testing.T) {
	dec := NewDecoder(strings.NewReader(`invalid`))
	defer dec.Release()
	err := dec.skipString()
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}

func TestParseEscapedString(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult string
		err            bool
		errType        interface{}
	}{
		{
			name:           "escape quote err",
			json:           `"test string \" escaped"`,
			expectedResult: `test string " escaped`,
			err:            false,
		},
		{
			name:           "escape quote err2",
			json:           `"test string \\t escaped"`,
			expectedResult: "test string \t escaped",
			err:            false,
		},
		{
			name:           "escape quote err2",
			json:           `"test string \\r escaped"`,
			expectedResult: "test string \r escaped",
			err:            false,
		},
		{
			name:           "escape quote err2",
			json:           `"test string \\b escaped"`,
			expectedResult: "test string \b escaped",
			err:            false,
		},
		{
			name:           "escape quote err",
			json:           `"test string \\n escaped"`,
			expectedResult: "test string \n escaped",
			err:            false,
		},
		{
			name:           "escape quote err",
			json:           `"test string \\" escaped"`,
			expectedResult: ``,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "escape quote err",
			json:           `"test string \\\l escaped"`,
			expectedResult: ``,
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			str := ""
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.Decode(&str)
			if testCase.err {
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of expected type")
				}
				log.Print(err)
			} else {
				assert.Nil(t, err, "err should be nil")
			}
			assert.Equal(t, testCase.expectedResult, str, fmt.Sprintf("str should be equal to '%s'", testCase.expectedResult))
		})
	}

}

func TestSkipString(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult string
		err            bool
		errType        interface{}
	}{
		{
			name:           "escape quote err",
			json:           `test string \\" escaped"`,
			expectedResult: ``,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "escape quote err",
			json:           `test string \\\l escaped"`,
			expectedResult: ``,
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}

	for _, testCase := range testCases {
		str := ""
		dec := NewDecoder(strings.NewReader(testCase.json))
		err := dec.skipString()
		if testCase.err {
			assert.NotNil(t, err, "err should not be nil")
			if testCase.errType != nil {
				assert.IsType(t, testCase.errType, err, "err should be of expected type")
			}
			log.Print(err)
		} else {
			assert.Nil(t, err, "err should be nil")
		}
		assert.Equal(t, testCase.expectedResult, str, fmt.Sprintf("str should be equal to '%s'", testCase.expectedResult))
	}
}
