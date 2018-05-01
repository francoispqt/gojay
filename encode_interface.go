package gojay

import (
	"fmt"
	"reflect"
)

// Encode encodes a value to JSON.
//
// If Encode cannot find a way to encode the type to JSON
// it will return an InvalidMarshalError.
func (enc *Encoder) Encode(v interface{}) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	switch vt := v.(type) {
	case string:
		return enc.EncodeString(vt)
	case bool:
		return enc.EncodeBool(vt)
	case MarshalerArray:
		return enc.EncodeArray(vt)
	case MarshalerObject:
		return enc.EncodeObject(vt)
	case int:
		return enc.EncodeInt(vt)
	case int64:
		return enc.EncodeInt64(vt)
	case int32:
		return enc.EncodeInt(int(vt))
	case int8:
		return enc.EncodeInt(int(vt))
	case uint64:
		return enc.EncodeInt(int(vt))
	case uint32:
		return enc.EncodeInt(int(vt))
	case uint16:
		return enc.EncodeInt(int(vt))
	case uint8:
		return enc.EncodeInt(int(vt))
	case float64:
		return enc.EncodeFloat(vt)
	case float32:
		return enc.EncodeFloat32(vt)
	default:
		return InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
}

// AddInterface adds an interface{} to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInterface(value interface{}) error {
	switch vt := value.(type) {
	case string:
		return enc.AddString(vt)
	case bool:
		return enc.AddBool(vt)
	case MarshalerArray:
		return enc.AddArray(vt)
	case MarshalerObject:
		return enc.AddObject(vt)
	case int:
		return enc.AddInt(vt)
	case int64:
		return enc.AddInt(int(vt))
	case int32:
		return enc.AddInt(int(vt))
	case int8:
		return enc.AddInt(int(vt))
	case uint64:
		return enc.AddInt(int(vt))
	case uint32:
		return enc.AddInt(int(vt))
	case uint16:
		return enc.AddInt(int(vt))
	case uint8:
		return enc.AddInt(int(vt))
	case float64:
		return enc.AddFloat(vt)
	case float32:
		return enc.AddFloat32(vt)
	}

	return nil
}

// AddInterfaceKey adds an interface{} to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInterfaceKey(key string, value interface{}) error {
	switch vt := value.(type) {
	case string:
		return enc.AddStringKey(key, vt)
	case bool:
		return enc.AddBoolKey(key, vt)
	case MarshalerArray:
		return enc.AddArrayKey(key, value.(MarshalerArray))
	case MarshalerObject:
		return enc.AddObjectKey(key, value.(MarshalerObject))
	case int:
		return enc.AddIntKey(key, vt)
	case int64:
		return enc.AddIntKey(key, int(vt))
	case int32:
		return enc.AddIntKey(key, int(vt))
	case int16:
		return enc.AddIntKey(key, int(vt))
	case int8:
		return enc.AddIntKey(key, int(vt))
	case uint64:
		return enc.AddIntKey(key, int(vt))
	case uint32:
		return enc.AddIntKey(key, int(vt))
	case uint16:
		return enc.AddIntKey(key, int(vt))
	case uint8:
		return enc.AddIntKey(key, int(vt))
	case float64:
		return enc.AddFloatKey(key, vt)
	case float32:
		return enc.AddFloat32Key(key, vt)
	}

	return nil
}
