package gojay

import (
	"fmt"
	"io"
	"reflect"
)

// UnmarshalArray parses the JSON-encoded data and stores the result in the value pointed to by v.
//
// v must implement UnmarshalerArray.
//
// If a JSON value is not appropriate for a given target type, or if a JSON number
// overflows the target type, UnmarshalArray skips that field and completes the unmarshaling as best it can.
func UnmarshalArray(data []byte, v UnmarshalerArray) error {
	dec := newDecoder(nil, 0)
	dec.data = data
	dec.length = len(data)
	_, err := dec.DecodeArray(v)
	dec.addToPool()
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

// UnmarshalObject parses the JSON-encoded data and stores the result in the value pointed to by v.
//
// v must implement UnmarshalerObject.
//
// If a JSON value is not appropriate for a given target type, or if a JSON number
// overflows the target type, UnmarshalObject skips that field and completes the unmarshaling as best it can.
func UnmarshalObject(data []byte, v UnmarshalerObject) error {
	dec := newDecoder(nil, 0)
	dec.data = data
	dec.length = len(data)
	_, err := dec.DecodeObject(v)
	dec.addToPool()
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// If v is nil, not a pointer, or not an implementation of UnmarshalerObject or UnmarshalerArray
// Unmarshal returns an InvalidUnmarshalError.
//
// Unmarshal uses the inverse of the encodings that Marshal uses, allocating maps, slices, and pointers as necessary, with the following additional rules:
// To unmarshal JSON into a pointer, Unmarshal first handles the case of the JSON being the JSON literal null.
// In that case, Unmarshal sets the pointer to nil.
// Otherwise, Unmarshal unmarshals the JSON into the value pointed at by the pointer.
// If the pointer is nil, Unmarshal allocates a new value for it to point to.
//
// To Unmarshal JSON into a struct, Unmarshal requires the struct to implement UnmarshalerObject.
//
// To unmarshal a JSON array into a slice, Unmarshal requires the slice to implement UnmarshalerArray.
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
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeString(vt)
	case *int:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt(vt)
	case *int32:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt32(vt)
	case *uint32:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeUint32(vt)
	case *int64:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt64(vt)
	case *uint64:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeUint64(vt)
	case *float64:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeFloat64(vt)
	case *bool:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeBool(vt)
	case UnmarshalerObject:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		_, err = dec.DecodeObject(vt)
	case UnmarshalerArray:
		dec = newDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		_, err = dec.DecodeArray(vt)
	default:
		return InvalidUnmarshalError(fmt.Sprintf(invalidUnmarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
	defer dec.addToPool()
	if err != nil {
		return err
	}
	return dec.err
}

// UnmarshalerObject is the interface to implement for a struct to be
// decoded
type UnmarshalerObject interface {
	UnmarshalObject(*Decoder, string) error
	NKeys() int
}

// UnmarshalerArray is the interface to implement for a slice or an array to be
// decoded
type UnmarshalerArray interface {
	UnmarshalArray(*Decoder) error
}

// UnmarshalerStream is the interface to implement for a slice, an array or a slice
// to decode a line delimited JSON to.
type UnmarshalerStream interface {
	UnmarshalStream(*StreamDecoder) error
}

// A Decoder reads and decodes JSON values from an input stream.
type Decoder struct {
	data     []byte
	cursor   int
	length   int
	keysDone int
	called   byte
	child    byte
	err      error
	r        io.Reader
}

// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) Decode(v interface{}) error {
	switch vt := v.(type) {
	case *string:
		return dec.DecodeString(vt)
	case *int:
		return dec.DecodeInt(vt)
	case *int32:
		return dec.DecodeInt32(vt)
	case *uint32:
		return dec.DecodeUint32(vt)
	case *int64:
		return dec.DecodeInt64(vt)
	case *uint64:
		return dec.DecodeUint64(vt)
	case *float64:
		return dec.DecodeFloat64(vt)
	case *bool:
		return dec.DecodeBool(vt)
	case UnmarshalerObject:
		_, err := dec.DecodeObject(vt)
		return err
	case UnmarshalerArray:
		_, err := dec.DecodeArray(vt)
		return err
	default:
		return InvalidUnmarshalError(fmt.Sprintf(invalidUnmarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
}

// ADD VALUES FUNCTIONS

// AddInt decodes the next key to an *int.
// If next key value overflows int, an InvalidTypeError error will be returned.
func (dec *Decoder) AddInt(v *int) error {
	err := dec.DecodeInt(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// AddFloat decodes the next key to a *float64.
// If next key value overflows float64, an InvalidTypeError error will be returned.
func (dec *Decoder) AddFloat(v *float64) error {
	err := dec.DecodeFloat64(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// AddBool decodes the next key to a *bool.
// If next key is neither null nor a JSON boolean, an InvalidTypeError will be returned.
// If next key is null, bool will be false.
func (dec *Decoder) AddBool(v *bool) error {
	err := dec.DecodeBool(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// AddString decodes the next key to a *string.
// If next key is not a JSON string nor null, InvalidTypeError will be returned.
func (dec *Decoder) AddString(v *string) error {
	err := dec.DecodeString(v)
	if err != nil {
		return err
	}
	dec.called |= 1
	return nil
}

// AddObject decodes the next key to a UnmarshalerObject.
func (dec *Decoder) AddObject(value UnmarshalerObject) error {
	initialKeysDone := dec.keysDone
	initialChild := dec.child
	dec.keysDone = 0
	dec.called = 0
	dec.child |= 1
	newCursor, err := dec.DecodeObject(value)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.keysDone = initialKeysDone
	dec.child = initialChild
	dec.called |= 1
	return nil
}

// AddArray decodes the next key to a UnmarshalerArray.
func (dec *Decoder) AddArray(value UnmarshalerArray) error {
	newCursor, err := dec.DecodeArray(value)
	if err != nil {
		return err
	}
	dec.cursor = newCursor
	dec.called |= 1
	return nil
}

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
		// idea is to append data from reader at the end
		n, err := dec.r.Read(dec.data[dec.length:])
		if err != nil || n == 0 {
			return false
		}
		dec.length = dec.length + n
		return true
	}
	return false
}

func (dec *Decoder) nextChar() byte {
	for dec.cursor < dec.length || dec.read() {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			dec.cursor = dec.cursor + 1
			continue
		}
		d := dec.data[dec.cursor]
		return d
	}
	return 0
}
