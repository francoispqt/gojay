package tests

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *A) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "str":
		return dec.String(&v.Str)
	case "bool":
		return dec.Bool(&v.Bool)
	case "int":
		return dec.Int(&v.Int)
	case "int64":
		return dec.Int64(&v.Int64)
	case "int32":
		return dec.Int32(&v.Int32)
	case "int16":
		return dec.Int16(&v.Int16)
	case "int8":
		return dec.Int8(&v.Int8)
	case "uint64":
		return dec.Uint64(&v.Uint64)
	case "uint32":
		return dec.Uint32(&v.Uint32)
	case "uint16":
		return dec.Uint16(&v.Uint16)
	case "uint8":
		return dec.Uint8(&v.Uint8)
	case "bval":
		if v.Bval == nil {
			v.Bval = &B{}
		}
		dec.Object(v.Bval)
	case "arrval":
		if v.Arrval == nil {
			arr := make(StrSlice, 0)
			v.Arrval = &arr
		}
		dec.Array(v.Arrval)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *A) NKeys() int { return 13 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *A) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("str", v.Str)
	enc.BoolKey("bool", v.Bool)
	enc.IntKey("int", v.Int)
	enc.Int64Key("int64", v.Int64)
	enc.Int32Key("int32", v.Int32)
	enc.Int16Key("int16", v.Int16)
	enc.Int8Key("int8", v.Int8)
	enc.Uint64Key("uint64", v.Uint64)
	enc.Uint32Key("uint32", v.Uint32)
	enc.Uint16Key("uint16", v.Uint16)
	enc.Uint8Key("uint8", v.Uint8)
	enc.ObjectKey("bval", v.Bval)
	enc.ArrayKey("arrval", v.Arrval)
}

// IsNil returns wether the structure is nil value or not
func (v *A) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *B) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "str":
		return dec.String(&v.Str)
	case "bool":
		return dec.Bool(&v.Bool)
	case "int":
		return dec.Int(&v.Int)
	case "int64":
		return dec.Int64(&v.Int64)
	case "int32":
		return dec.Int32(&v.Int32)
	case "int16":
		return dec.Int16(&v.Int16)
	case "int8":
		return dec.Int8(&v.Int8)
	case "uint64":
		return dec.Uint64(&v.Uint64)
	case "uint32":
		return dec.Uint32(&v.Uint32)
	case "uint16":
		return dec.Uint16(&v.Uint16)
	case "uint8":
		return dec.Uint8(&v.Uint8)
	case "strptr":
		return dec.String(v.StrPtr)
	case "boolptr":
		return dec.Bool(v.BoolPtr)
	case "intptr":
		return dec.Int(v.IntPtr)
	case "int64ptr":
		return dec.Int64(v.Int64Ptr)
	case "int32ptr":
		return dec.Int32(v.Int32Ptr)
	case "int16ptr":
		return dec.Int16(v.Int16Ptr)
	case "int8ptr":
		return dec.Int8(v.Int8Ptr)
	case "uint64ptr":
		return dec.Uint64(v.Uint64Ptr)
	case "uint32ptr":
		return dec.Uint32(v.Uint32Ptr)
	case "uint16ptr":
		return dec.Uint16(v.Uint16Ptr)
	case "uint8ptr":
		return dec.Uint8(v.Uint8PTr)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *B) NKeys() int { return 22 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *B) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("str", v.Str)
	enc.BoolKey("bool", v.Bool)
	enc.IntKey("int", v.Int)
	enc.Int64Key("int64", v.Int64)
	enc.Int32Key("int32", v.Int32)
	enc.Int16Key("int16", v.Int16)
	enc.Int8Key("int8", v.Int8)
	enc.Uint64Key("uint64", v.Uint64)
	enc.Uint32Key("uint32", v.Uint32)
	enc.Uint16Key("uint16", v.Uint16)
	enc.Uint8Key("uint8", v.Uint8)
	enc.StringKey("strptr", *v.StrPtr)
	enc.BoolKey("boolptr", *v.BoolPtr)
	enc.IntKey("intptr", *v.IntPtr)
	enc.Int64Key("int64ptr", *v.Int64Ptr)
	enc.Int32Key("int32ptr", *v.Int32Ptr)
	enc.Int16Key("int16ptr", *v.Int16Ptr)
	enc.Int8Key("int8ptr", *v.Int8Ptr)
	enc.Uint64Key("uint64ptr", *v.Uint64Ptr)
	enc.Uint32Key("uint32ptr", *v.Uint32Ptr)
	enc.Uint16Key("uint16ptr", *v.Uint16Ptr)
	enc.Uint8Key("uint8ptr", *v.Uint8PTr)
}

// IsNil returns wether the structure is nil value or not
func (v *B) IsNil() bool { return v == nil }
