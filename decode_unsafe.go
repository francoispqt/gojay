package gojay

import (
	"fmt"
	"reflect"
)

// Unsafe is the structure holding the unsafe version of the API.
// The difference between unsafe api and regular api is that the regular API
// copies the buffer passed to Unmarshal functions to a new internal buffer.
// Making it safer because internally GoJay uses unsafe.Pointer to transform slice of bytes into a string.
var Unsafe = decUnsafe{}

type decUnsafe struct{}

func (u decUnsafe) UnmarshalArray(data []byte, v UnmarshalerArray) error {
	dec := BorrowDecoder(nil, 0)
	dec.data = data
	dec.length = len(data)
	_, err := dec.DecodeArray(v)
	dec.Release()
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

func (u decUnsafe) UnmarshalObject(data []byte, v UnmarshalerObject) error {
	dec := BorrowDecoder(nil, 0)
	dec.data = data
	dec.length = len(data)
	_, err := dec.DecodeObject(v)
	dec.Release()
	if err != nil {
		return err
	}
	if dec.err != nil {
		return dec.err
	}
	return nil
}

func (u decUnsafe) Unmarshal(data []byte, v interface{}) error {
	var err error
	var dec *Decoder
	switch vt := v.(type) {
	case *string:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeString(vt)
	case *int:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt(vt)
	case *int32:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt32(vt)
	case *uint32:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeUint32(vt)
	case *int64:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeInt64(vt)
	case *uint64:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeUint64(vt)
	case *float64:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeFloat64(vt)
	case *bool:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		err = dec.DecodeBool(vt)
	case UnmarshalerObject:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		_, err = dec.DecodeObject(vt)
	case UnmarshalerArray:
		dec = BorrowDecoder(nil, 0)
		dec.length = len(data)
		dec.data = data
		_, err = dec.DecodeArray(vt)
	default:
		return InvalidUnmarshalError(fmt.Sprintf(invalidUnmarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
	defer dec.Release()
	if err != nil {
		return err
	}
	return dec.err
}
