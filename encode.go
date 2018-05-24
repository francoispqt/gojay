package gojay

import (
	"io"
	"fmt"
	"reflect"
	"encoding/json"
)

// MarshalJSONObject returns the JSON encoding of v.
//
// It takes a struct implementing Marshaler to a JSON slice of byte
// it returns a slice of bytes and an error.
// Example with an Marshaler:
//	type TestStruct struct {
//		id int
//	}
//	func (s *TestStruct) MarshalJSONObject(enc *gojay.Encoder) {
//		enc.AddIntKey("id", s.id)
//	}
//	func (s *TestStruct) IsNil() bool {
//		return s == nil
//	}
//
// 	func main() {
//		test := &TestStruct{
//			id: 123456,
//		}
//		b, _ := gojay.Marshal(test)
// 		fmt.Println(b) // {"id":123456}
//	}
func MarshalJSONObject(v MarshalerJSONObject) ([]byte, error) {
	enc := BorrowEncoder(nil)
	enc.grow(512)
	defer enc.Release()
	return enc.encodeObject(v)
}

// MarshalJSONArray returns the JSON encoding of v.
//
// It takes an array or a slice implementing Marshaler to a JSON slice of byte
// it returns a slice of bytes and an error.
// Example with an Marshaler:
// 	type TestSlice []*TestStruct
//
// 	func (t TestSlice) MarshalJSONArray(enc *Encoder) {
//		for _, e := range t {
//			enc.AddObject(e)
//		}
//	}
//
//	func main() {
//		test := &TestSlice{
//			&TestStruct{123456},
//			&TestStruct{7890},
// 		}
// 		b, _ := Marshal(test)
//		fmt.Println(b) // [{"id":123456},{"id":7890}]
//	}
func MarshalJSONArray(v MarshalerJSONArray) ([]byte, error) {
	enc := BorrowEncoder(nil)
	enc.grow(512)
	enc.writeByte('[')
	v.(MarshalerJSONArray).MarshalJSONArray(enc)
	enc.writeByte(']')
	defer enc.Release()
	return enc.buf, nil
}

// Marshal returns the JSON encoding of v.
//
// Marshal takes interface v and encodes it according to its type.
// Basic example with a string:
// 	b, err := gojay.Marshal("test")
//	fmt.Println(b) // "test"
//
// If v implements Marshaler or Marshaler interface
// it will call the corresponding methods.
//
// If a struct, slice, or array is passed and does not implement these interfaces
// it will return a a non nil InvalidUnmarshalError error.
// Example with an Marshaler:
//	type TestStruct struct {
//		id int
//	}
//	func (s *TestStruct) MarshalJSONObject(enc *gojay.Encoder) {
//		enc.AddIntKey("id", s.id)
//	}
//	func (s *TestStruct) IsNil() bool {
//		return s == nil
//	}
//
// 	func main() {
//		test := &TestStruct{
//			id: 123456,
//		}
//		b, _ := gojay.Marshal(test)
// 		fmt.Println(b) // {"id":123456}
//	}
func Marshal(v interface{}) ([]byte, error) {
	return marshal(v, false)
}

// MarshalAny returns the JSON encoding of v.
//
// MarshalAny takes interface v and encodes it according to its type.
// Basic example with a string:
// 	b, err := gojay.Marshal("test")
//	fmt.Println(b) // "test"
//
// If v implements Marshaler or Marshaler interface
// it will call the corresponding methods.
//
// If it cannot find any supported type it will be marshalled though default Go "json" package.
// Warning, this function can be slower, than a default "Marshal"
//
//	type TestStruct struct {
//		id int
//	}
//
// 	func main() {
//		test := &TestStruct{
//			id: 123456,
//		}
//		b, _ := gojay.Marshal(test)
// 		fmt.Println(b) // {"id": 123456}
//	}
func MarshalAny(v interface{}) ([]byte, error) {
	return marshal(v, true)
}

func marshal(v interface{}, any bool) ([]byte, error) {
	var (
		enc = BorrowEncoder(nil)

		buf []byte
		err error
	)

	buf, err = func() ([]byte, error) {
		switch vt := v.(type) {
		case MarshalerJSONObject:
			return enc.encodeObject(vt)
		case MarshalerJSONArray:
			return enc.encodeArray(vt)
		case string:
			return enc.encodeString(vt)
		case bool:
			return enc.encodeBool(vt)
		case int:
			return enc.encodeInt(vt)
		case int64:
			return enc.encodeInt64(vt)
		case int32:
			return enc.encodeInt(int(vt))
		case int16:
			return enc.encodeInt(int(vt))
		case int8:
			return enc.encodeInt(int(vt))
		case uint64:
			return enc.encodeInt(int(vt))
		case uint32:
			return enc.encodeInt(int(vt))
		case uint16:
			return enc.encodeInt(int(vt))
		case uint8:
			return enc.encodeInt(int(vt))
		case float64:
			return enc.encodeFloat(vt)
		case float32:
			return enc.encodeFloat32(vt)
		case *EmbeddedJSON:
			return enc.encodeEmbeddedJSON(vt)
		default:
			if any {
				return json.Marshal(vt)
			}

			return nil, InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, reflect.TypeOf(vt).String()))
		}
	} ()

	enc.Release()
	return buf, err
}

// MarshalerJSONObject is the interface to implement for struct to be encoded
type MarshalerJSONObject interface {
	MarshalJSONObject(enc *Encoder)
	IsNil() bool
}

// MarshalerJSONArray is the interface to implement
// for a slice or an array to be encoded
type MarshalerJSONArray interface {
	MarshalJSONArray(enc *Encoder)
	IsNil() bool
}

// An Encoder writes JSON values to an output stream.
type Encoder struct {
	buf      []byte
	isPooled byte
	w        io.Writer
	err      error
}

// AppendBytes allows a modular usage by appending bytes manually to the current state of the buffer.
func (enc *Encoder) AppendBytes(b []byte) {
	enc.writeBytes(b)
}

// AppendByte allows a modular usage by appending a single byte manually to the current state of the buffer.
func (enc *Encoder) AppendByte(b byte) {
	enc.writeByte(b)
}

// Buf returns the Encoder's buffer.
func (enc *Encoder) Buf() []byte {
	return enc.buf
}

// Write writes to the io.Writer and resets the buffer.
func (enc *Encoder) Write() (int, error) {
	i, err := enc.w.Write(enc.buf)
	enc.buf = enc.buf[:0]
	return i, err
}

func (enc *Encoder) getPreviousRune() byte {
	last := len(enc.buf) - 1
	return enc.buf[last]
}
