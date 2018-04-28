package gojay

import (
	"fmt"
	"reflect"
)

// Encode encodes a value to JSON.
//
// If Encode cannot find a way to encode the type to JSON
// it will return an InvalidMarshalError.
func (enc *Encoder) Encode(v interface{}) ([]byte, error) {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	switch vt := v.(type) {
	case string:
		return enc.encodeString(vt)
	case bool:
		return enc.encodeBool(vt)
	case MarshalerArray:
		return enc.encodeArray(vt)
	case MarshalerObject:
		return enc.encodeObject(vt)
	case int:
		return enc.encodeInt(int64(vt))
	case int64:
		return enc.encodeInt(vt)
	case int32:
		return enc.encodeInt(int64(vt))
	case int8:
		return enc.encodeInt(int64(vt))
	case uint64:
		return enc.encodeInt(int64(vt))
	case uint32:
		return enc.encodeInt(int64(vt))
	case uint16:
		return enc.encodeInt(int64(vt))
	case uint8:
		return enc.encodeInt(int64(vt))
	case float64:
		return enc.encodeFloat(vt)
	case float32:
		return enc.encodeFloat(float64(vt))
	default:
		return nil, InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
}

// AddInterface adds an interface{} to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInterface(value interface{}) error {
	switch value.(type) {
	case string:
		return enc.AddString(value.(string))
	case bool:
		return enc.AddBool(value.(bool))
	case MarshalerArray:
		return enc.AddArray(value.(MarshalerArray))
	case MarshalerObject:
		return enc.AddObject(value.(MarshalerObject))
	case int:
		return enc.AddInt(value.(int))
	case int64:
		return enc.AddInt(int(value.(int64)))
	case int32:
		return enc.AddInt(int(value.(int32)))
	case int8:
		return enc.AddInt(int(value.(int8)))
	case uint64:
		return enc.AddInt(int(value.(uint64)))
	case uint32:
		return enc.AddInt(int(value.(uint32)))
	case uint16:
		return enc.AddInt(int(value.(uint16)))
	case uint8:
		return enc.AddInt(int(value.(uint8)))
	case float64:
		return enc.AddFloat(value.(float64))
	case float32:
		return enc.AddFloat(float64(value.(float32)))
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
