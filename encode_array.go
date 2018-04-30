package gojay

// EncodeArray encodes an implementation of MarshalerArray to JSON
func (enc *Encoder) EncodeArray(v MarshalerArray) ([]byte, error) {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	return enc.encodeArray(v)
}
func (enc *Encoder) encodeArray(v MarshalerArray) ([]byte, error) {
	enc.grow(200)
	enc.writeByte('[')
	v.MarshalArray(enc)
	enc.writeByte(']')
	return enc.buf, nil
}

// AddArray adds an array or slice to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement Marshaler
func (enc *Encoder) AddArray(value MarshalerArray) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('[')
	value.MarshalArray(enc)
	enc.writeByte(']')
	return nil
}

// AddArrayKey adds an array or slice to be encoded, must be used inside an object as it will encode a key
// value must implement Marshaler
func (enc *Encoder) AddArrayKey(key string, value MarshalerArray) error {
	// grow to avoid allocs (length of key/value + quotes)
	r, ok := enc.getPreviousRune()
	if ok && r != '[' && r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKeyArr)
	value.MarshalArray(enc)
	enc.writeByte(']')
	return nil
}
