package gojay

import (
	"fmt"
	"io"
	"reflect"
	"time"
)

// UnmarshalJSONArray parses the JSON-encoded data and stores the result in the value pointed to by v.
//
// v must implement UnmarshalerJSONArray.
//
// If a JSON value is not appropriate for a given target type, or if a JSON number
// overflows the target type, UnmarshalJSONArray skips that field and completes the unmarshaling as best it can.
func UnmarshalJSONArray(data []byte, v UnmarshalerJSONArray) error {
	dec := borrowDecoder(nil, 0)
	defer dec.Release()
	dec.data = make([]byte, len(data))
	copy(dec.data, data)
	dec.length = len(data)
	_, err := dec.decodeArray(v)
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

// UnmarshalJSONObject parses the JSON-encoded data and stores the result in the value pointed to by v.
//
// v must implement UnmarshalerJSONObject.
//
// If a JSON value is not appropriate for a given target type, or if a JSON number
// overflows the target type, UnmarshalJSONObject skips that field and completes the unmarshaling as best it can.
func UnmarshalJSONObject(data []byte, v UnmarshalerJSONObject) error {
	dec := borrowDecoder(nil, 0)
	defer dec.Release()
	dec.data = make([]byte, len(data))
	copy(dec.data, data)
	dec.length = len(data)
	_, err := dec.decodeObject(v)
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// If v is nil, not a pointer, or not an implementation of UnmarshalerJSONObject or UnmarshalerJSONArray
// Unmarshal returns an InvalidUnmarshalError.
//
// Unmarshal uses the inverse of the encodings that Marshal uses, allocating maps, slices, and pointers as necessary, with the following additional rules:
// To unmarshal JSON into a pointer, Unmarshal first handles the case of the JSON being the JSON literal null.
// In that case, Unmarshal sets the pointer to nil.
// Otherwise, Unmarshal unmarshals the JSON into the value pointed at by the pointer.
// If the pointer is nil, Unmarshal allocates a new value for it to point to.
//
// To Unmarshal JSON into a struct, Unmarshal requires the struct to implement UnmarshalerJSONObject.
//
// To unmarshal a JSON array into a slice, Unmarshal requires the slice to implement UnmarshalerJSONArray.
//
// Unmarshal JSON does not allow yet to unmarshall an interface value
// If a JSON value is not appropriate for a given target type, or if a JSON number
// overflows the target type, Unmarshal skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns an UnmarshalTypeError describing the earliest such error. In any case, it's not guaranteed that all the remaining fields following the problematic one will be unmarshaled into the target object.
func Unmarshal(data []byte, v interface{}) error {
	var err error
	var dec *Decoder
	switch vt := v.(type) {
	case *string:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeString(vt)
	case **string:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeStringNull(vt)
	case *int:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt(vt)
	case **int:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeIntNull(vt)
	case *int8:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt8(vt)
	case **int8:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt8Null(vt)
	case *int16:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt16(vt)
	case **int16:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt16Null(vt)
	case *int32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt32(vt)
	case **int32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt32Null(vt)
	case *int64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt64(vt)
	case **int64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeInt64Null(vt)
	case *uint8:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint8(vt)
	case **uint8:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint8Null(vt)
	case *uint16:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint16(vt)
	case **uint16:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint16Null(vt)
	case *uint32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint32(vt)
	case **uint32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint32Null(vt)
	case *uint64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint64(vt)
	case **uint64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeUint64Null(vt)
	case *float64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeFloat64(vt)
	case **float64:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeFloat64Null(vt)
	case *float32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeFloat32(vt)
	case **float32:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeFloat32Null(vt)
	case *bool:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeBool(vt)
	case **bool:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.decodeBoolNull(vt)
	case UnmarshalerJSONObject:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = make([]byte, len(data))
		copy(dec.data, data)
		_, err = dec.decodeObject(vt)
	case UnmarshalerJSONArray:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = make([]byte, len(data))
		copy(dec.data, data)
		_, err = dec.decodeArray(vt)
	case *interface{}:
		dec = borrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = make([]byte, len(data))
		copy(dec.data, data)
		err = dec.decodeInterface(vt)
	default:
		return InvalidUnmarshalError(fmt.Sprintf(invalidUnmarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
	defer dec.Release()
	if err != nil {
		return err
	}
	return dec.err
}

// UnmarshalerJSONObject is the interface to implement for a struct to be
// decoded
type UnmarshalerJSONObject interface {
	UnmarshalJSONObject(*Decoder, string) error
	NKeys() int
}

// UnmarshalerJSONArray is the interface to implement for a slice or an array to be
// decoded
type UnmarshalerJSONArray interface {
	UnmarshalJSONArray(*Decoder) error
}

// A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	r        io.Reader
	data     []byte
	err      error
	isPooled byte
	called   byte
	child    byte
	cursor   int
	length   int
	keysDone int
}

// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v interface{}) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	var err error
	switch vt := v.(type) {
	case *string:
		err = dec.decodeString(vt)
	case **string:
		err = dec.decodeStringNull(vt)
	case *int:
		err = dec.decodeInt(vt)
	case **int:
		err = dec.decodeIntNull(vt)
	case *int8:
		err = dec.decodeInt8(vt)
	case **int8:
		err = dec.decodeInt8Null(vt)
	case *int16:
		err = dec.decodeInt16(vt)
	case **int16:
		err = dec.decodeInt16Null(vt)
	case *int32:
		err = dec.decodeInt32(vt)
	case **int32:
		err = dec.decodeInt32Null(vt)
	case *int64:
		err = dec.decodeInt64(vt)
	case **int64:
		err = dec.decodeInt64Null(vt)
	case *uint8:
		err = dec.decodeUint8(vt)
	case **uint8:
		err = dec.decodeUint8Null(vt)
	case *uint16:
		err = dec.decodeUint16(vt)
	case **uint16:
		err = dec.decodeUint16Null(vt)
	case *uint32:
		err = dec.decodeUint32(vt)
	case **uint32:
		err = dec.decodeUint32Null(vt)
	case *uint64:
		err = dec.decodeUint64(vt)
	case **uint64:
		err = dec.decodeUint64Null(vt)
	case *float64:
		err = dec.decodeFloat64(vt)
	case **float64:
		err = dec.decodeFloat64Null(vt)
	case *float32:
		err = dec.decodeFloat32(vt)
	case **float32:
		err = dec.decodeFloat32Null(vt)
	case *bool:
		err = dec.decodeBool(vt)
	case **bool:
		err = dec.decodeBoolNull(vt)
	case UnmarshalerJSONObject:
		_, err = dec.decodeObject(vt)
	case UnmarshalerJSONArray:
		_, err = dec.decodeArray(vt)
	case *EmbeddedJSON:
		err = dec.decodeEmbeddedJSON(vt)
	case *interface{}:
		err = dec.decodeInterface(vt)
	default:
		return InvalidUnmarshalError(fmt.Sprintf(invalidUnmarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
	if err != nil {
		return err
	}
	return dec.err
}

// Reset resets the decoder to free the buffer an reuse the same decoder
func (dec *Decoder) Reset() {
	dec.called = 0
	dec.keysDone = 0
	dec.cursor = 0
	dec.err = nil
	dec.r = nil
	dec.length = 0
	dec.isPooled = 0
	dec.data = make([]byte, 512)
}

// SetReader sets the reader of the decoder
func (dec *Decoder) SetReader(r io.Reader) {
	dec.r = r
}

// ADD VALUES FUNCTIONS

// AddInt decodes the next key to an *int.
// If next key value overflows int, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddInt(v *int) error {
	return dec.Int(v)
}

// AddIntNull decodes the next key to an *int.
// If next key value overflows int, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddIntNull(v **int) error {
	return dec.IntNull(v)
}

// AddInt8 decodes the next key to an *int.
// If next key value overflows int8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddInt8(v *int8) error {
	return dec.Int8(v)
}

// AddInt8Null decodes the next key to an *int.
// If next key value overflows int8, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddInt8Null(v **int8) error {
	return dec.Int8Null(v)
}

// AddInt16 decodes the next key to an *int.
// If next key value overflows int16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddInt16(v *int16) error {
	return dec.Int16(v)
}

// AddInt16Null decodes the next key to an *int.
// If next key value overflows int16, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddInt16Null(v **int16) error {
	return dec.Int16Null(v)
}

// AddInt32 decodes the next key to an *int.
// If next key value overflows int32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddInt32(v *int32) error {
	return dec.Int32(v)
}

// AddInt32Null decodes the next key to an *int.
// If next key value overflows int32, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddInt32Null(v **int32) error {
	return dec.Int32Null(v)
}

// AddInt64 decodes the next key to an *int.
// If next key value overflows int64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddInt64(v *int64) error {
	return dec.Int64(v)
}

// AddInt64Null decodes the next key to an *int.
// If next key value overflows int64, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddInt64Null(v **int64) error {
	return dec.Int64Null(v)
}

// AddUint8 decodes the next key to an *int.
// If next key value overflows uint8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddUint8(v *uint8) error {
	return dec.Uint8(v)
}

// AddUint8Null decodes the next key to an *int.
// If next key value overflows uint8, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddUint8Null(v **uint8) error {
	return dec.Uint8Null(v)
}

// AddUint16 decodes the next key to an *int.
// If next key value overflows uint16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddUint16(v *uint16) error {
	return dec.Uint16(v)
}

// AddUint16Null decodes the next key to an *int.
// If next key value overflows uint16, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddUint16Null(v **uint16) error {
	return dec.Uint16Null(v)
}

// AddUint32 decodes the next key to an *int.
// If next key value overflows uint32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddUint32(v *uint32) error {
	return dec.Uint32(v)
}

// AddUint32Null decodes the next key to an *int.
// If next key value overflows uint32, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddUint32Null(v **uint32) error {
	return dec.Uint32Null(v)
}

// AddUint64 decodes the next key to an *int.
// If next key value overflows uint64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddUint64(v *uint64) error {
	return dec.Uint64(v)
}

// AddUint64Null decodes the next key to an *int.
// If next key value overflows uint64, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddUint64Null(v **uint64) error {
	return dec.Uint64Null(v)
}

// AddFloat decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddFloat(v *float64) error {
	return dec.Float64(v)
}

// AddFloatNull decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddFloatNull(v **float64) error {
	return dec.Float64Null(v)
}

// AddFloat64 decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddFloat64(v *float64) error {
	return dec.Float64(v)
}

// AddFloat64Null decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddFloat64Null(v **float64) error {
	return dec.Float64Null(v)
}

// AddFloat32 decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) AddFloat32(v *float32) error {
	return dec.Float32(v)
}

// AddFloat32Null decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddFloat32Null(v **float32) error {
	return dec.Float32Null(v)
}

// AddBool decodes the next key to a *bool.
// If next key is neither null nor a JSON boolean, an InvalidUnmarshalError will be returned.
// If next key is null, bool will be false.
func (dec *Decoder) AddBool(v *bool) error {
	return dec.Bool(v)
}

// AddBoolNull decodes the next key to a *bool.
// If next key is neither null nor a JSON boolean, an InvalidUnmarshalError will be returned.
// If next key is null, bool will be false.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddBoolNull(v **bool) error {
	return dec.BoolNull(v)
}

// AddString decodes the next key to a *string.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) AddString(v *string) error {
	return dec.String(v)
}

// AddStringNull decodes the next key to a *string.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) AddStringNull(v **string) error {
	return dec.StringNull(v)
}

// AddObject decodes the next key to a UnmarshalerJSONObject.
func (dec *Decoder) AddObject(v UnmarshalerJSONObject) error {
	return dec.Object(v)
}

// AddObjectNull decodes the next key to a UnmarshalerJSONObject.
func (dec *Decoder) AddObjectNull(v interface{}) error {
	return dec.ObjectNull(v)
}

// AddArray decodes the next key to a UnmarshalerJSONArray.
func (dec *Decoder) AddArray(v UnmarshalerJSONArray) error {
	return dec.Array(v)
}

// AddArray decodes the next key to a UnmarshalerJSONArray.
func (dec *Decoder) AddArrayNull(v UnmarshalerJSONArray) error {
	return dec.ArrayNull(v)
}

// AddInterface decodes the next key to a interface{}.
func (dec *Decoder) AddInterface(v *interface{}) error {
	return dec.Interface(v)
}

// Int decodes the next key to an *int.
// If next key value overflows int, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int(v *int) error {
	err := dec.decodeInt(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// IntNull decodes the next key to an *int.
// If next key value overflows int, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) IntNull(v **int) error {
	err := dec.decodeIntNull(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int8 decodes the next key to an *int.
// If next key value overflows int8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int8(v *int8) error {
	err := dec.decodeInt8(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int8Null decodes the next key to an *int.
// If next key value overflows int8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int8Null(v **int8) error {
	err := dec.decodeInt8Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int16 decodes the next key to an *int.
// If next key value overflows int16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int16(v *int16) error {
	err := dec.decodeInt16(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int16Null decodes the next key to an *int.
// If next key value overflows int16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int16Null(v **int16) error {
	err := dec.decodeInt16Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int32 decodes the next key to an *int.
// If next key value overflows int32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int32(v *int32) error {
	err := dec.decodeInt32(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int32Null decodes the next key to an *int.
// If next key value overflows int32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int32Null(v **int32) error {
	err := dec.decodeInt32Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int64 decodes the next key to an *int.
// If next key value overflows int64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int64(v *int64) error {
	err := dec.decodeInt64(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Int64Null decodes the next key to an *int.
// If next key value overflows int64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Int64Null(v **int64) error {
	err := dec.decodeInt64Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint8 decodes the next key to an *int.
// If next key value overflows uint8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint8(v *uint8) error {
	err := dec.decodeUint8(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint8Null decodes the next key to an *int.
// If next key value overflows uint8, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint8Null(v **uint8) error {
	err := dec.decodeUint8Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint16 decodes the next key to an *int.
// If next key value overflows uint16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint16(v *uint16) error {
	err := dec.decodeUint16(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint16Null decodes the next key to an *int.
// If next key value overflows uint16, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint16Null(v **uint16) error {
	err := dec.decodeUint16Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint32 decodes the next key to an *int.
// If next key value overflows uint32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint32(v *uint32) error {
	err := dec.decodeUint32(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint32Null decodes the next key to an *int.
// If next key value overflows uint32, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint32Null(v **uint32) error {
	err := dec.decodeUint32Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint64 decodes the next key to an *int.
// If next key value overflows uint64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint64(v *uint64) error {
	err := dec.decodeUint64(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Uint64Null decodes the next key to an *int.
// If next key value overflows uint64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Uint64Null(v **uint64) error {
	err := dec.decodeUint64Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Float decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Float(v *float64) error {
	return dec.Float64(v)
}

// FloatNull decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) FloatNull(v **float64) error {
	return dec.Float64Null(v)
}

// Float64 decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Float64(v *float64) error {
	err := dec.decodeFloat64(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Float64Null decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Float64Null(v **float64) error {
	err := dec.decodeFloat64Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Float32 decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Float32(v *float32) error {
	err := dec.decodeFloat32(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Float32Null decodes the next key to a *float64.
// If next key value overflows float64, an InvalidUnmarshalError error will be returned.
func (dec *Decoder) Float32Null(v **float32) error {
	err := dec.decodeFloat32Null(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Bool decodes the next key to a *bool.
// If next key is neither null nor a JSON boolean, an InvalidUnmarshalError will be returned.
// If next key is null, bool will be false.
func (dec *Decoder) Bool(v *bool) error {
	err := dec.decodeBool(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// BoolNull decodes the next key to a *bool.
// If next key is neither null nor a JSON boolean, an InvalidUnmarshalError will be returned.
// If next key is null, bool will be false.
func (dec *Decoder) BoolNull(v **bool) error {
	err := dec.decodeBoolNull(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// String decodes the next key to a *string.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) String(v *string) error {
	err := dec.decodeString(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// StringNull decodes the next key to a **string.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
// If a `null` is encountered, gojay does not change the value of the pointer.
func (dec *Decoder) StringNull(v **string) error {
	err := dec.decodeStringNull(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// AddTime decodes the next key to a *time.Time with the given format
func (dec *Decoder) AddTime(v *time.Time, format string) error {
	return dec.Time(v, format)
}

// Time decodes the next key to a *time.Time with the given format
func (dec *Decoder) Time(v *time.Time, format string) error {
	err := dec.decodeTime(v, format)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Object decodes the next key to a UnmarshalerJSONObject.
func (dec *Decoder) Object(value UnmarshalerJSONObject) error {
	initialKeysDone := dec.keysDone
	initialChild := dec.child
	dec.keysDone = 0
	dec.called = 0
	dec.child |= 1
	newCursor, err := dec.decodeObject(value)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.keysDone = initialKeysDone
	dec.child = initialChild
	dec.called |= 1
	return nil
}

// ObjectNull decodes the next key to a UnmarshalerJSONObject.
// v should be a pointer to an UnmarshalerJSONObject,
// if `null` value is encountered in JSON, it will leave the value v untouched,
// else it will create a new instance of the UnmarshalerJSONObject behind v.
func (dec *Decoder) ObjectNull(v interface{}) error {
	initialKeysDone := dec.keysDone
	initialChild := dec.child
	dec.keysDone = 0
	dec.called = 0
	dec.child |= 1
	newCursor, err := dec.decodeObjectNull(v)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.keysDone = initialKeysDone
	dec.child = initialChild
	dec.called |= 1
	return nil
}

// Array decodes the next key to a UnmarshalerJSONArray.
func (dec *Decoder) Array(v UnmarshalerJSONArray) error {
	newCursor, err := dec.decodeArray(v)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.called |= 1
	return nil
}

// ArrayNull decodes the next key to a UnmarshalerJSONArray.
// v should be a pointer to an UnmarshalerJSONArray,
// if `null` value is encountered in JSON, it will leave the value v untouched,
// else it will create a new instance of the UnmarshalerJSONArray behind v.
func (dec *Decoder) ArrayNull(v interface{}) error {
	newCursor, err := dec.decodeArrayNull(v)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.called |= 1
	return nil
}

// Interface decodes the next key to an interface{}.
func (dec *Decoder) Interface(value *interface{}) error {
	err := dec.decodeInterface(value)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// Array decodes the next key to a UnmarshalerJSONArray.
// func (dec *Decoder) ArrayNull(factory func() UnmarshalerJSONArray) error {
// 	newCursor, err := dec.decodeArrayNull(factory)
// 	if err != nil {
// 		return err
// 	}
// 	dec.cursor = newCursor
// 	dec.called |= 1
// 	return nil
// }

// Non exported

func isDigit(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func (dec *Decoder) read() bool {
	if dec.r != nil {
		// if we reach the end, double the buffer to ensure there's always more space
		if len(dec.data) == dec.length {
			nLen := dec.length * 2
			Buf := make([]byte, nLen, nLen)
			copy(Buf, dec.data)
			dec.data = Buf
		}
		var n int
		var err error
		for n == 0 {
			n, err = dec.r.Read(dec.data[dec.length:])
			if err != nil {
				if err != io.EOF {
					dec.err = err
					return false
				}
				if n == 0 {
					return false
				}
				dec.length = dec.length + n
				return true
			}
		}
		dec.length = dec.length + n
		return true
	}
	return false
}

func (dec *Decoder) nextChar() byte {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		}
		d := dec.data[dec.cursor]
		return d
	}
	return 0
}
