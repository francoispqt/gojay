package gojay

import "strconv"

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

// AddBoolKey adds a bool to be encoded, must be used inside an object as it will encode a key
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
