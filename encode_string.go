package gojay

// EncodeString encodes a string to
func (enc *Encoder) EncodeString(s string) ([]byte, error) {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	return enc.encodeString(s)
}

// encodeString encodes a string to
func (enc *Encoder) encodeString(s string) ([]byte, error) {
	enc.writeByte('"')
	enc.writeString(s)
	enc.writeByte('"')
	return enc.buf, nil
}

// AddString adds a string to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddString(value string) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(value)
	enc.writeByte('"')

	return nil
}

// AddStringKey adds a string to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddStringKey(key, value string) error {
	// grow to avoid allocs (length of key/value + quotes)
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.write(objKeyStr)
	enc.writeString(value)
	enc.writeByte('"')

	return nil
}
