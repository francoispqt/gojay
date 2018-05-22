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
		return enc.EncodeUint64(vt)
	case uint32:
		return enc.EncodeUint(uint(vt))
	case uint16:
		return enc.EncodeUint(uint(vt))
	case uint8:
		return enc.EncodeUint(uint(vt))
	case float64:
		return enc.EncodeFloat(vt)
	case float32:
		return enc.EncodeFloat32(vt)
	case *EmbeddedJSON:
		return enc.EncodeEmbeddedJSON(vt)
	default:
		return InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, reflect.TypeOf(vt).String()))
	}
}

// AddInterface adds an interface{} to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInterface(value interface{}) {
	switch vt := value.(type) {
	case string:
		enc.AddString(vt)
	case bool:
		enc.AddBool(vt)
	case MarshalerArray:
		enc.AddArray(vt)
	case MarshalerObject:
		enc.AddObject(vt)
	case int:
		enc.AddInt(vt)
	case int64:
		enc.AddInt(int(vt))
	case int32:
		enc.AddInt(int(vt))
	case int8:
		enc.AddInt(int(vt))
	case uint64:
		enc.AddInt(int(vt))
	case uint32:
		enc.AddInt(int(vt))
	case uint16:
		enc.AddInt(int(vt))
	case uint8:
		enc.AddInt(int(vt))
	case float64:
		enc.AddFloat(vt)
	case float32:
		enc.AddFloat32(vt)
	default:
		t := reflect.TypeOf(vt)
		if t != nil {
			enc.err = InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, t.String()))
			return
		}
		return
	}
}

// AddInterfaceKey adds an interface{} to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInterfaceKey(key string, value interface{}) {
	switch vt := value.(type) {
	case string:
		enc.AddStringKey(key, vt)
	case bool:
		enc.AddBoolKey(key, vt)
	case MarshalerArray:
		enc.AddArrayKey(key, vt)
	case MarshalerObject:
		enc.AddObjectKey(key, vt)
	case int:
		enc.AddIntKey(key, vt)
	case int64:
		enc.AddIntKey(key, int(vt))
	case int32:
		enc.AddIntKey(key, int(vt))
	case int16:
		enc.AddIntKey(key, int(vt))
	case int8:
		enc.AddIntKey(key, int(vt))
	case uint64:
		enc.AddIntKey(key, int(vt))
	case uint32:
		enc.AddIntKey(key, int(vt))
	case uint16:
		enc.AddIntKey(key, int(vt))
	case uint8:
		enc.AddIntKey(key, int(vt))
	case float64:
		enc.AddFloatKey(key, vt)
	case float32:
		enc.AddFloat32Key(key, vt)
	default:
		t := reflect.TypeOf(vt)
		if t != nil {
			enc.err = InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, t.String()))
			return
		}
		return
	}
}

// AddInterfaceKeyOmitEmpty adds an interface{} to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInterfaceKeyOmitEmpty(key string, v interface{}) {
	switch vt := v.(type) {
	case string:
		enc.AddStringKeyOmitEmpty(key, vt)
	case bool:
		enc.AddBoolKeyOmitEmpty(key, vt)
	case MarshalerArray:
		enc.AddArrayKeyOmitEmpty(key, vt)
	case MarshalerObject:
		enc.AddObjectKeyOmitEmpty(key, vt)
	case int:
		enc.AddIntKeyOmitEmpty(key, vt)
	case int64:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case int32:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case int16:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case int8:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case uint64:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case uint32:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case uint16:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case uint8:
		enc.AddIntKeyOmitEmpty(key, int(vt))
	case float64:
		enc.AddFloatKeyOmitEmpty(key, vt)
	case float32:
		enc.AddFloat32KeyOmitEmpty(key, vt)
	default:
		t := reflect.TypeOf(vt)
		if t != nil {
			enc.err = InvalidMarshalError(fmt.Sprintf(invalidMarshalErrorMsg, t.String()))
			return
		}
		return
	}
}
