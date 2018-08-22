package gojay

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeInterfaceBasic(t *testing.T) {
	testCases := []struct {
		name            string
		json            string
		expectedResult  interface{}
		err             bool
		errType         interface{}
		skipCheckResult bool
	}{
		{
			name:           "array",
			json:           `[1,2,3]`,
			expectedResult: []interface{}([]interface{}{float64(1), float64(2), float64(3)}),
			err:            false,
		},
		{
			name:           "object",
			json:           `{"testStr": "hello world!"}`,
			expectedResult: map[string]interface{}(map[string]interface{}{"testStr": "hello world!"}),
			err:            false,
		},
		{
			name:            "array-error",
			json:            `["h""o","l","a"]`,
			err:             true,
			errType:         &json.SyntaxError{},
			skipCheckResult: true,
		},
		{
			name:            "object-error",
			json:            `{"testStr" "hello world!"}`,
			err:             true,
			errType:         &json.SyntaxError{},
			skipCheckResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var i interface{}
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&i)
			if testCase.err {
				t.Log(err)
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			assert.Nil(t, err, "err should be nil")
			if !testCase.skipCheckResult {
				assert.Equal(t, testCase.expectedResult, i, "value at given index should be the same as expected results")
			}
		})
	}
}

func TestDecodeInterfaceObject(t *testing.T) {
	testCases := []struct {
		name            string
		json            string
		expectedResult  testObject
		err             bool
		errType         interface{}
		skipCheckResult bool
	}{
		{
			name: "basic-array",
			json: `{
        "testStr": "hola",
        "testInterface": ["h","o","l","a"],
      }`,
			expectedResult: testObject{
				testStr:       "hola",
				testInterface: []interface{}([]interface{}{"h", "o", "l", "a"}),
			},
			err: false,
		},
		{
			name: "basic-string",
			json: `{
        "testInterface": "漢字",
      }`,
			expectedResult: testObject{
				testInterface: interface{}("漢字"),
			},
			err: false,
		},
		{
			name: "basic-interface",
			json: `{
        "testInterface": {
          "string": "prost"
        },
      }`,
			expectedResult: testObject{
				testInterface: map[string]interface{}{"string": "prost"},
			},
			err: false,
		},
		{
			name: "complex-interface",
			json: `{
        "testInterface": {
          "number": 1988,
          "string": "prost",
          "array": ["h","o","l","a"],
          "object": {
            "k": "v",
            "a": [1,2,3]
          },
          "array-of-objects": [
            {"k": "v"},
            {"a": "b"}
          ]
        },
      }`,
			expectedResult: testObject{
				testInterface: map[string]interface{}{
					"array-of-objects": []interface{}{
						map[string]interface{}{"k": "v"},
						map[string]interface{}{"a": "b"},
					},
					"number": float64(1988),
					"string": "prost",
					"array":  []interface{}{"h", "o", "l", "a"},
					"object": map[string]interface{}{
						"k": "v",
						"a": []interface{}{float64(1), float64(2), float64(3)},
					},
				},
			},
			err: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := testObject{}
			dec := BorrowDecoder(strings.NewReader(testCase.json))
			defer dec.Release()
			err := dec.Decode(&s)
			if testCase.err {
				t.Log(err)
				assert.NotNil(t, err, "err should not be nil")
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be of the given type")
				}
				return
			}
			assert.Nil(t, err, "err should be nil")
			if !testCase.skipCheckResult {
				assert.Equal(t, testCase.expectedResult, s, "value at given index should be the same as expected results")
			}
		})
	}
}
