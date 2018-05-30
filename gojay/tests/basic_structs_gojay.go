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
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *A) NKeys() int { return 12 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *A) MarshalJSONOject(enc *gojay.Encoder) {
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
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *B) NKeys() int { return 11 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *B) MarshalJSONOject(enc *gojay.Encoder) {
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
}

// IsNil returns wether the structure is nil value or not
func (v *B) IsNil() bool { return v == nil }
