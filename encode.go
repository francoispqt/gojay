package gojay

import (
	"fmt"
	"io"
	"reflect"
)

// MarshalObject returns the JSON encoding of v.
//
// It takes a struct implementing Marshaler to a JSON slice of byte
// it returns a slice of bytes and an error.
// Example with an Marshaler:
//	type TestStruct struct {
//		id int
//	}
//	func (s *TestStruct) MarshalObject(enc *gojay.Encoder) {
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
func MarshalObject(v MarshalerObject) ([]byte, error) {
	enc := newEncoder()
	defer enc.Release()
	return enc.encodeObject(v)
}

// MarshalArray returns the JSON encoding of v.
//
// It takes an array or a slice implementing Marshaler to a JSON slice of byte
// it returns a slice of bytes and an error.
// Example with an Marshaler:
// 	type TestSlice []*TestStruct
//
// 	func (t TestSlice) MarshalArray(enc *Encoder) {
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
func MarshalArray(v MarshalerArray) ([]byte, error) {
	enc := newEncoder()
	enc.grow(200)
	enc.writeByte('[')
	v.(MarshalerArray).MarshalArray(enc)
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
// it will return a a non nil InvalidTypeError error.
// Example with an Marshaler:
//	type TestStruct struct {
//		id int
//	}
//	func (s *TestStruct) MarshalObject(enc *gojay.Encoder) {
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
	switch vt := v.(type) {
	case MarshalerObject:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeObject(vt)
	case MarshalerArray:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeArray(vt)
	case string:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeString(vt)
	case bool:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeBool(vt)
	case int:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(vt)
	case int64:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt64(vt)
	case int32:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case int16:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case int8:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case uint64:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case uint32:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case uint16:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case uint8:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeInt(int(vt))
	case float64:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeFloat(vt)
	case float32:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeFloat32(vt)
	case *EmbeddedJSON:
		enc := BorrowEncoder(nil)
		defer enc.Release()
		return enc.encodeEmbeddedJSON(vt)
	default:
		return nil, InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
}

// MarshalerObject is the interface to implement for struct to be encoded
type MarshalerObject interface {
	MarshalObject(enc *Encoder)
	IsNil() bool
}

// MarshalerArray is the interface to implement
// for a slice or an array to be encoded
type MarshalerArray interface {
	MarshalArray(enc *Encoder)
	IsNil() bool
}

// An Encoder writes JSON values to an output stream.
type Encoder struct {
	buf      []byte
	isPooled byte
	w        io.Writer
	err      error
}

func (enc *Encoder) Init(w io.Writer) {
	enc.w = w
	enc.buf = make([]byte, 0, 512)
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
