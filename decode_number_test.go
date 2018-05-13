package gojay

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderInt(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult int
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-positive",
			json:           "100",
			expectedResult: 100,
		},
		{
			name:           "basic-positive2",
			json:           "1039405",
			expectedResult: 1039405,
		},
		{
			name:           "basic-negative",
			json:           "-2",
			expectedResult: -2,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-big",
			json:           "9223372036854775807",
			expectedResult: 9223372036854775807,
		},
		{
			name:           "basic-big-overflow",
			json:           "9223372036854775808",
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
		{
			name:           "basic-big-overflow2",
			json:           "92233720368547758089",
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
		{
			name:           "basic-big-overflow3",
			json:           "92233720368547758089 ",
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
		{
			name:           "basic-negative2",
			json:           "-2349557",
			expectedResult: -2349557,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: -7,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876 ",
			expectedResult: -7,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876a",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-exponent-positive-positive-exp",
			json:           "1e2",
			expectedResult: 100,
		},
		{
			name:           "basic-exponent-positive-positive-exp2",
			json:           "5e+06",
			expectedResult: 5000000,
		},
		{
			name:           "basic-exponent-positive-positive-exp3",
			json:           "3e+3",
			expectedResult: 3000,
		},
		{
			name:           "basic-exponent-positive-positive-exp4",
			json:           "8e+005",
			expectedResult: 800000,
		},
		{
			name:           "basic-exponent-positive-negative-exp",
			json:           "1e-2",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp2",
			json:           "5e-6",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp3",
			json:           "3e-3",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp4",
			json:           "8e-005",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-negative-positive-exp",
			json:           "-1e2",
			expectedResult: -100,
		},
		{
			name:           "basic-exponent-negative-positive-exp2",
			json:           "-5e+06",
			expectedResult: -5000000,
		},
		{
			name:           "basic-exponent-negative-positive-exp3",
			json:           "-3e03",
			expectedResult: -3000,
		},
		{
			name:           "basic-exponent-negative-positive-exp4",
			json:           "-8e+005",
			expectedResult: -800000,
		},
		{
			name:           "error1",
			json:           "132zz4",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "negative-error2",
			json:           " -1213xdde2323 ",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "error3",
			json:           "-8e+00$aa5",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v int
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult, v, fmt.Sprintf("v must be equal to %d", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := int(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeInt(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")
	})
	t.Run("decoder-api", func(t *testing.T) {
		var v int
		dec := NewDecoder(strings.NewReader(`33`))
		defer dec.Release()
		err := dec.DecodeInt(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, int(33), v, "v must be equal to 33")
	})
	t.Run("decoder-api-invalid-json", func(t *testing.T) {
		var v int
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeInt(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}
func TestDecoderInt64(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult int64
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-positive",
			json:           "100",
			expectedResult: 100,
		},
		{
			name:           "basic-positive2",
			json:           "1039405",
			expectedResult: 1039405,
		},
		{
			name:           "basic-negative",
			json:           "-2",
			expectedResult: -2,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-big",
			json:           "9223372036854775807",
			expectedResult: 9223372036854775807,
		},
		{
			name:           "basic-big-overflow",
			json:           "9223372036854775808",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-big-overflow2",
			json:           "92233720368547758089",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-big-overflow3",
			json:           "92233720368547758089 ",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-negative2",
			json:           "-2349557",
			expectedResult: -2349557,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: -7,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876a",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-exponent-positive-positive-exp",
			json:           "1e2",
			expectedResult: 100,
		},
		{
			name:           "basic-exponent-positive-positive-exp2",
			json:           "5e+06 ",
			expectedResult: 5000000,
		},
		{
			name:           "basic-exponent-positive-positive-exp3",
			json:           "3e+3",
			expectedResult: 3000,
		},
		{
			name:           "basic-exponent-positive-positive-exp4",
			json:           "8e+005",
			expectedResult: 800000,
		},
		{
			name:           "basic-exponent-positive-negative-exp",
			json:           "1e-2 ",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp2",
			json:           "5e-6",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp3",
			json:           "3e-3",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp4",
			json:           "8e-005",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-negative-positive-exp",
			json:           "-1e2",
			expectedResult: -100,
		},
		{
			name:           "basic-exponent-negative-positive-exp2",
			json:           "-5e+06",
			expectedResult: -5000000,
		},
		{
			name:           "basic-exponent-negative-positive-exp2",
			json:           "-5.4e+06",
			expectedResult: -5400000,
		},
		{
			name:           "basic-exponent-negative-positive-exp3",
			json:           "-3e03",
			expectedResult: -3000,
		},
		{
			name:           "basic-exponent-negative-positive-exp4",
			json:           "-8e+005",
			expectedResult: -800000,
		},
		{
			name:           "error1",
			json:           "132zz4",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v int64
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult, v, fmt.Sprintf("v must be equal to %d", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := int64(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeInt64(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")
	})
	t.Run("decoder-api", func(t *testing.T) {
		var v int64
		dec := NewDecoder(strings.NewReader(`33`))
		defer dec.Release()
		err := dec.DecodeInt64(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, int64(33), v, "v must be equal to 33")
	})
	t.Run("decoder-api-invalid-json", func(t *testing.T) {
		var v int64
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeInt64(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}
func TestDecoderUint64(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult uint64
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-positive",
			json:           "100",
			expectedResult: 100,
		},
		{
			name:           "basic-positive2",
			json:           "1039405",
			expectedResult: 1039405,
		},
		{
			name:           "basic-negative",
			json:           "-2",
			expectedResult: 2,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-big",
			json:           "18446744073709551615",
			expectedResult: 18446744073709551615,
		},
		{
			name:           "basic-big-overflow",
			json:           "18446744073709551616",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-big-overflow2",
			json:           "184467440737095516161",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-negative2",
			json:           "-2349557",
			expectedResult: 2349557,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: 7,
		},
		{
			name:           "error1",
			json:           "132zz4",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v uint64
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult, v, fmt.Sprintf("v must be equal to %d", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := uint64(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeUint64(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")
	})
	t.Run("decoder-api", func(t *testing.T) {
		var v uint64
		dec := NewDecoder(strings.NewReader(`33`))
		defer dec.Release()
		err := dec.DecodeUint64(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, uint64(33), v, "v must be equal to 33")
	})
	t.Run("decoder-api-json-error", func(t *testing.T) {
		var v uint64
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeUint64(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}
func TestDecoderInt32(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult int32
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-positive",
			json:           "100",
			expectedResult: 100,
		},
		{
			name:           "basic-positive2",
			json:           "1039405",
			expectedResult: 1039405,
		},
		{
			name:           "basic-negative",
			json:           "-2",
			expectedResult: -2,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-negative2",
			json:           "-2349557",
			expectedResult: -2349557,
		},
		{
			name:           "basic-big",
			json:           "2147483647",
			expectedResult: 2147483647,
		},
		{
			name:           "basic-big-overflow",
			json:           " 2147483648",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-big-overflow2",
			json:           "21474836483",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: -7,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876a",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-exponent-positive-positive-exp",
			json:           "1.2E2",
			expectedResult: 120,
		},
		{
			name:           "basic-exponent-positive-positive-exp1",
			json:           "3.5e+005",
			expectedResult: 350000,
		},
		{
			name:           "basic-exponent-positive-positive-exp2",
			json:           "5e+06",
			expectedResult: 5000000,
		},
		{
			name:           "basic-exponent-positive-positive-exp3",
			json:           "3e+3",
			expectedResult: 3000,
		},
		{
			name:           "basic-exponent-positive-positive-exp4",
			json:           "8e+005 ",
			expectedResult: 800000,
		},
		{
			name:           "basic-exponent-positive-negative-exp",
			json:           "1e-2 ",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp2",
			json:           "5E-6",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp3",
			json:           "3e-3",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-positive-negative-exp4",
			json:           "8e-005",
			expectedResult: 0,
		},
		{
			name:           "basic-exponent-negative-positive-exp",
			json:           "-1e2",
			expectedResult: -100,
		},
		{
			name:           "basic-exponent-negative-positive-exp2",
			json:           "-5e+06",
			expectedResult: -5000000,
		},
		{
			name:           "basic-exponent-negative-positive-exp3",
			json:           "-3e03",
			expectedResult: -3000,
		},
		{
			name:           "basic-exponent-negative-positive-exp4",
			json:           "-8e+005",
			expectedResult: -800000,
		},
		{
			name:           "error3",
			json:           "-8e+00$aa5",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v int32
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult, v, fmt.Sprintf("v must be equal to %d", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := int32(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeInt32(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")

	})
	t.Run("decoder-api", func(t *testing.T) {
		var v int32
		dec := NewDecoder(strings.NewReader(`33`))
		defer dec.Release()
		err := dec.DecodeInt32(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, int32(33), v, "v must be equal to 33")
	})
	t.Run("decoder-api-invalid-json", func(t *testing.T) {
		var v int32
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeInt32(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}

func TestDecoderUint32(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult uint32
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-positive",
			json:           "100",
			expectedResult: 100,
		},
		{
			name:           "basic-positive2",
			json:           "1039405",
			expectedResult: 1039405,
		},
		{
			name:           "basic-negative",
			json:           "-2",
			expectedResult: 2,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-negative2",
			json:           "-2349557",
			expectedResult: 2349557,
		},
		{
			name:           "basic-big",
			json:           "4294967295",
			expectedResult: 4294967295,
		},
		{
			name:           "basic-big-overflow",
			json:           " 4294967298",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-big-overflow2",
			json:           "42949672983",
			expectedResult: 0,
			err:            true,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: 7,
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
		{
			name:           "invalid-json",
			json:           `123invalid`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v uint32
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult, v, fmt.Sprintf("v must be equal to %d", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := uint32(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeUint32(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")
	})
	t.Run("decoder-api", func(t *testing.T) {
		var v uint32
		dec := NewDecoder(strings.NewReader(`33`))
		defer dec.Release()
		err := dec.DecodeUint32(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, uint32(33), v, "v must be equal to 33")
	})
	t.Run("decoder-api-json-error", func(t *testing.T) {
		var v uint32
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeUint32(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}

func TestDecoderFloat64(t *testing.T) {
	testCases := []struct {
		name           string
		json           string
		expectedResult float64
		err            bool
		errType        interface{}
	}{
		{
			name:           "basic-exponent-positive-positive-exp",
			json:           "1e2",
			expectedResult: 100,
		},
		{
			name:           "basic-exponent-positive-positive-exp2",
			json:           "5e+06",
			expectedResult: 5000000,
		},
		{
			name:           "basic-exponent-positive-positive-exp3",
			json:           "3e+3",
			expectedResult: 3000,
		},
		{
			name:           "basic-null",
			json:           "null",
			expectedResult: 0,
		},
		{
			name:           "basic-null-err",
			json:           "nxll",
			expectedResult: 0,
			err:            true,
			errType:        InvalidJSONError(""),
		},
		{
			name:           "basic-exponent-positive-positive-exp4",
			json:           "8e+005",
			expectedResult: 800000,
		},
		{
			name:           "basic-exponent-positive-negative-exp",
			json:           "1e-2",
			expectedResult: 0.01,
		},
		{
			name:           "basic-exponent-positive-negative-exp2",
			json:           "5e-6",
			expectedResult: 0.000005,
		},
		{
			name:           "basic-exponent-positive-negative-exp3",
			json:           "3e-3",
			expectedResult: 0.003,
		},
		{
			name:           "basic-exponent-positive-negative-exp4",
			json:           "8e-005",
			expectedResult: 0.00008,
		},
		{
			name:           "basic-exponent-negative-positive-exp",
			json:           "-1e2",
			expectedResult: -100,
		},
		{
			name:           "basic-exponent-negative-positive-exp2",
			json:           "-5e+06",
			expectedResult: -5000000,
		},
		{
			name:           "basic-exponent-negative-positive-exp3",
			json:           "-3e03",
			expectedResult: -3000,
		},
		{
			name:           "basic-exponent-negative-positive-exp4",
			json:           "-8e+005",
			expectedResult: -800000,
		},
		{
			name:           "basic-exponent-negative-positive-exp4",
			json:           "-8.2e-005",
			expectedResult: -0.000082,
		},
		{
			name:           "basic-float",
			json:           "2.4595",
			expectedResult: 2.4595,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876",
			expectedResult: -7.8876,
		},
		{
			name:           "basic-float",
			json:           "2.4595e1",
			expectedResult: 24.595,
		},
		{
			name:           "basic-float2",
			json:           "-7.8876e002",
			expectedResult: -788.76,
		},
		{
			name:           "invalid-type",
			json:           `"string"`,
			expectedResult: 0,
			err:            true,
			errType:        InvalidTypeError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json := []byte(testCase.json)
			var v float64
			err := Unmarshal(json, &v)
			if testCase.err {
				assert.NotNil(t, err, "Err must not be nil")
				if testCase.errType != nil {
					assert.IsType(
						t,
						testCase.errType,
						err,
						fmt.Sprintf("err should be of type %s", reflect.TypeOf(err).String()),
					)
				}
			} else {
				assert.Nil(t, err, "Err must be nil")
			}
			assert.Equal(t, testCase.expectedResult*1000000, math.Round(v*1000000), fmt.Sprintf("v must be equal to %f", testCase.expectedResult))
		})
	}
	t.Run("pool-error", func(t *testing.T) {
		result := float64(1)
		dec := NewDecoder(nil)
		dec.Release()
		defer func() {
			err := recover()
			assert.NotNil(t, err, "err shouldnot be nil")
			assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		}()
		_ = dec.DecodeFloat64(&result)
		assert.True(t, false, "should not be called as decoder should have panicked")
	})
	t.Run("decoder-api", func(t *testing.T) {
		var v float64
		dec := NewDecoder(strings.NewReader(`1.25`))
		defer dec.Release()
		err := dec.DecodeFloat64(&v)
		assert.Nil(t, err, "Err must be nil")
		assert.Equal(t, 1.25, v, "v must be equal to 1.25")
	})
	t.Run("decoder-api-json-error", func(t *testing.T) {
		var v float64
		dec := NewDecoder(strings.NewReader(``))
		defer dec.Release()
		err := dec.DecodeFloat64(&v)
		assert.NotNil(t, err, "Err must not be nil")
		assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
	})
}
