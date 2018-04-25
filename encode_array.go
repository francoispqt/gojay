package gojay

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
	enc.write(objKeyArr)
	value.MarshalArray(enc)
	enc.writeByte(']')
	return nil
}
