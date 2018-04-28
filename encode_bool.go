package gojay

import "strconv"

// EncodeBool encodes a bool to JSON
func (enc *Encoder) EncodeBool(v bool) ([]byte, error) {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	return enc.encodeBool(v)
}

// encodeBool encodes a bool to JSON
func (enc *Encoder) encodeBool(v bool) ([]byte, error) {
	if v {
		enc.writeString("true")
	} else {
		enc.writeString("false")
	}
	return enc.buf, nil
}

// AddBool adds a bool to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddBool(value bool) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	if value {
		enc.writeString("true")
	} else {
		enc.writeString("false")
	}
	return nil
}

// AddBoolKey adds a bool to be encoded, must be used inside an object as it will encode a key.
func (enc *Encoder) AddBoolKey(key string, value bool) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.write(objKey)
	enc.buf = strconv.AppendBool(enc.buf, value)
	return nil
}
