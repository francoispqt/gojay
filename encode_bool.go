package gojay

import "strconv"

// EncodeBool encodes a bool to JSON
func (enc *Encoder) EncodeBool(v bool) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeBool(v)
	_, err := enc.write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

// encodeBool encodes a bool to JSON
func (enc *Encoder) encodeBool(v bool) ([]byte, error) {
	if v {
		enc.writeString("true")
	} else {
		enc.writeString("false")
	}
	return enc.buf, enc.err
}

// AddBool adds a bool to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddBool(v bool) {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	if v {
		enc.writeString("true")
	} else {
		enc.writeString("false")
	}
}

// AddBoolOmitEmpty adds a bool to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddBoolOmitEmpty(v bool) {
	if v == false {
		return
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeString("true")
}

// AddBoolKey adds a bool to be encoded, must be used inside an object as it will encode a key.
func (enc *Encoder) AddBoolKey(key string, value bool) {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendBool(enc.buf, value)
}

// AddBoolKeyOmitEmpty adds a bool to be encoded and skips it if it is zero value.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddBoolKeyOmitEmpty(key string, v bool) {
	if v == false {
		return
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendBool(enc.buf, v)
}
