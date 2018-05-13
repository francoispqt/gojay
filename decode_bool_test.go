package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderBool(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult bool
		expectations   func(t *testing.T, v bool, err error)
	}{
		{
			name: "true-basic",
			json: "true",
			expectations: func(t *testing.T, v bool, err error) {
				assert.Nil(t, err, "err should be nil")
				assert.True(t, v, "result should be true")
			},
		},
		{
			name: "false-basic",
			json: "false",
			expectations: func(t *testing.T, v bool, err error) {
				assert.Nil(t, err, "err should be nil")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "null-basic",
			json: "null",
			expectations: func(t *testing.T, v bool, err error) {
				assert.Nil(t, err, "err should be nil")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "true-error",
			json: "taue",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "true-error2",
			json: "trae",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "true-error3",
			json: "trua",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "true-error4",
			json: "truea",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			name: "true-error4",
			json: "t",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "fulse",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "fause",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "falze",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "falso",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "falsea",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "f",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "nall",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "nual",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "nula",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "nulle",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
		{
			json: "n",
			expectations: func(t *testing.T, v bool, err error) {
				assert.NotNil(t, err, "err should be nil")
				assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
				assert.False(t, v, "result should be false")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v bool
			err := Unmarshal(json, &v)
			testCase.expectations(t, v, err)
		})
	}
}

func TestDecoderBoolTrue(t *testing.T) {
	json := []byte(`true`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, true, v, "v must be equal to true")
}

func TestDecoderBoolFalse(t *testing.T) {
	json := []byte(`false`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, false, v, "v must be equal to false")
}

func TestDecoderBoolInvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v bool
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, false, v, "v must be equal to false as it is zero val")
}

func TestDecoderBoolNonBooleanJSONFalse(t *testing.T) {
	json := []byte(`null`)
	var v bool
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, false, v, "v must be equal to false")
}

func TestDecoderBoolInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	var v bool
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}
func TestDecoderBoolDecoderAPI(t *testing.T) {
	var v bool
	dec := BorrowDecoder(strings.NewReader("true"))
	defer dec.Release()
	err := dec.DecodeBool(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, true, v, "v must be equal to true")
}

func TestDecoderBoolPoolError(t *testing.T) {
	v := true
	dec := NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	_ = dec.DecodeBool(&v)
	assert.True(t, false, "should not be called as decoder should have panicked")
}
