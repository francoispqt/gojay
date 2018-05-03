package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderIntBasic(t *testing.T) {
	json := []byte(`124`)
	var v int
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 124, v, "v must be equal to 124")
}
func TestDecoderIntNegative(t *testing.T) {
	json := []byte(` -124 `)
	var v int
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, -124, v, "v must be equal to -124")
}
func TestDecoderIntNegativeError(t *testing.T) {
	json := []byte(` -12x4 `)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderIntNull(t *testing.T) {
	json := []byte(`null`)
	var v int
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int(0), v, "v must be equal to 0")
}
func TestDecoderIntInvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeErrorr")
}
func TestDecoderIntInvalidJSON(t *testing.T) {
	json := []byte(`123n`)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderIntBig(t *testing.T) {
	json := []byte(`9223372036854775807`)
	var v int
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 9223372036854775807, v, "v must be equal to 9223372036854775807")
}
func TestDecoderIntOverfow(t *testing.T) {
	json := []byte(`9223372036854775808`)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as int is overflowing")
	assert.Equal(t, 0, v, "v must be equal to 0")
}
func TestDecoderIntOverfow2(t *testing.T) {
	json := []byte(`92233720368547758089 `)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as int is overflowing")
	assert.Equal(t, 0, v, "v must be equal to 0")
}
func TestDecoderIntOverfow3(t *testing.T) {
	json := []byte(`92233720368547758089 `)
	var v int
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as int is overflowing")
	assert.Equal(t, 0, v, "v must be equal to 0")
}
func TestDecoderIntPoolError(t *testing.T) {
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
}
func TestDecoderIntDecoderAPI(t *testing.T) {
	var v int
	dec := NewDecoder(strings.NewReader(`33`))
	defer dec.Release()
	err := dec.DecodeInt(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int(33), v, "v must be equal to 33")
}

func TestDecoderIntInvalidJSONError(t *testing.T) {
	var v int
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeInt(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecoderInt32Basic(t *testing.T) {
	json := []byte(`124`)
	var v int32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(124), v, "v must be equal to 124")
}
func TestDecoderInt32Negative(t *testing.T) {
	json := []byte(`-124 `)
	var v int32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(-124), v, "v must be equal to -124")
}
func TestDecoderInt32NegativeError(t *testing.T) {
	json := []byte(`-12x4 `)
	var v int32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderInt32Null(t *testing.T) {
	json := []byte(`null`)
	var v int32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(0), v, "v must be equal to 0")
}
func TestDecoderInt32InvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v int32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeErrorr")
}
func TestDecoderInt32InvalidJSON(t *testing.T) {
	json := []byte(`123n`)
	var v int32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderInt32Big(t *testing.T) {
	json := []byte(`2147483647`)
	var v int32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "err must not be nil as int32 does not overflow")
	assert.Equal(t, int32(2147483647), v, "int32 must be equal to 2147483647")
}
func TestDecoderInt32Overflow(t *testing.T) {
	json := []byte(` 2147483648`)
	var v int32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderInt32Overflow2(t *testing.T) {
	json := []byte(`21474836483`)
	var v int32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderInt32PoolError(t *testing.T) {
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
}
func TestDecoderInt32tDecoderAPI(t *testing.T) {
	var v int32
	dec := NewDecoder(strings.NewReader(`33`))
	defer dec.Release()
	err := dec.DecodeInt32(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(33), v, "v must be equal to 33")
}

func TestDecoderInt32InvalidJSONError(t *testing.T) {
	var v int32
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeInt32(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecoderUint32Basic(t *testing.T) {
	json := []byte(`124 `)
	var v uint32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(124), v, "v must be equal to 124")
}
func TestDecoderUint32Null(t *testing.T) {
	json := []byte(`null`)
	var v uint32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(0), v, "v must be equal to 0")
}
func TestDecoderUint32InvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v uint32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeErrorr")
}
func TestDecoderUint32InvalidJSON(t *testing.T) {
	json := []byte(`123n`)
	var v uint32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderUint32Big(t *testing.T) {
	json := []byte(`4294967295 `)
	var v uint32
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "err must not be nil as uint32 does not overflow")
	assert.Equal(t, uint32(4294967295), v, "err must be of type InvalidTypeError")
}
func TestDecoderUint32Overflow(t *testing.T) {
	json := []byte(`4294967298`)
	var v uint32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as uint32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}

func TestDecoderUint32Overflow2(t *testing.T) {
	json := []byte(`42949672983`)
	var v uint32
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as uint32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderUint32PoolError(t *testing.T) {
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
}
func TestDecoderUint32tDecoderAPI(t *testing.T) {
	var v uint32
	dec := NewDecoder(strings.NewReader(`33`))
	defer dec.Release()
	err := dec.DecodeUint32(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(33), v, "v must be equal to 33")
}

func TestDecoderUint32InvalidJSONError(t *testing.T) {
	var v uint32
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeUint32(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecoderInt64Basic(t *testing.T) {
	json := []byte(`124 `)
	var v int64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(124), v, "v must be equal to 124")
}
func TestDecoderInt64Negative(t *testing.T) {
	json := []byte(`-124`)
	var v int64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(-124), v, "v must be equal to -124")
}
func TestDecoderInt64Null(t *testing.T) {
	json := []byte(`null`)
	var v int64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(0), v, "v must be equal to 0")
}
func TestDecoderInt64InvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v int64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeErrorr")
}
func TestDecoderInt64InvalidJSON(t *testing.T) {
	json := []byte(`123n`)
	var v int64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderInt64Big(t *testing.T) {
	json := []byte(`9223372036854775807`)
	var v int64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "err must not be nil as int64 does not overflow")
	assert.Equal(t, int64(9223372036854775807), v, "err must be of type InvalidTypeError")
}
func TestDecoderInt64Overflow(t *testing.T) {
	json := []byte(`9223372036854775808`)
	var v int64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int64 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderInt64Overflow2(t *testing.T) {
	json := []byte(`92233720368547758082`)
	var v int64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int64 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderInt64PoolError(t *testing.T) {
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
}
func TestDecoderInt64DecoderAPI(t *testing.T) {
	var v int64
	dec := NewDecoder(strings.NewReader(`33`))
	defer dec.Release()
	err := dec.DecodeInt64(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(33), v, "v must be equal to 33")
}
func TestDecoderInt64InvalidJSONError(t *testing.T) {
	var v int64
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeInt64(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecoderUint64Basic(t *testing.T) {
	json := []byte(` 124 `)
	var v uint64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(124), v, "v must be equal to 124")
}
func TestDecoderUint64Null(t *testing.T) {
	json := []byte(`null`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(0), v, "v must be equal to 0")
}
func TestDecoderUint64InvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeErrorr")
}
func TestDecoderUint64InvalidJSON(t *testing.T) {
	json := []byte(`123n`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err must be of type InvalidJSONError")
}
func TestDecoderUint64Big(t *testing.T) {
	json := []byte(`18446744073709551615`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "err must not be nil as uint64 does not overflow")
	assert.Equal(t, uint64(18446744073709551615), v, "err must be of type InvalidTypeError")
}
func TestDecoderUint64Overflow(t *testing.T) {
	json := []byte(`18446744073709551616`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderUint64Overflow2(t *testing.T) {
	json := []byte(`184467440737095516161`)
	var v uint64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil as int32 overflows")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type InvalidTypeError")
}
func TestDecoderUint64PoolError(t *testing.T) {
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
}
func TestDecoderUint64tDecoderAPI(t *testing.T) {
	var v uint64
	dec := NewDecoder(strings.NewReader(`33`))
	defer dec.Release()
	err := dec.DecodeUint64(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(33), v, "v must be equal to 33")
}

func TestDecoderUint64InvalidJSONError(t *testing.T) {
	var v uint64
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeUint64(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}

func TestDecoderFloatBasic(t *testing.T) {
	json := []byte(`100.11 `)
	var v float64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 100.11, v, "v must be equal to 100.11")
}
func TestDecoderFloatBasic2(t *testing.T) {
	json := []byte(` 100.11 `)
	var v float64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 100.11, v, "v must be equal to 100.11")
}
func TestDecoderFloatBasic3(t *testing.T) {
	json := []byte(` 100 `)
	var v float64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float64(100), v, "v must be equal to 100.11")
}

func TestDecoderFloatBig(t *testing.T) {
	json := []byte(`89899843.3493493 `)
	var v float64
	err := Unmarshal(json, &v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 89899843.3493493, v, "v must be equal to 8989984340.3493493")
}

func TestDecoderFloatInvalidType(t *testing.T) {
	json := []byte(`"string"`)
	var v float64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "err must not be nil")
	assert.IsType(t, InvalidTypeError(""), err, "err must be of type *strconv.NumError")
}

func TestDecoderFloatInvalidJSON(t *testing.T) {
	json := []byte(`hello`)
	var v float64
	err := Unmarshal(json, &v)
	assert.NotNil(t, err, "Err must not be nil as JSON is invalid")
	assert.IsType(t, InvalidJSONError(""), err, "err message must be 'Invalid JSON'")
}
func TestDecoderFloatDecoderAPI(t *testing.T) {
	var v float64
	dec := NewDecoder(strings.NewReader(`1.25`))
	defer dec.Release()
	err := dec.DecodeFloat64(&v)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, 1.25, v, "v must be equal to 1.25")
}
func TestDecoderFloatPoolError(t *testing.T) {
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
}

func TestDecoderFloatInvalidJSONError(t *testing.T) {
	var v float64
	dec := NewDecoder(strings.NewReader(``))
	defer dec.Release()
	err := dec.DecodeFloat64(&v)
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, InvalidJSONError(""), err, "err should be of type InvalidJSONError")
}
