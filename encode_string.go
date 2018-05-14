package gojay

// EncodeString encodes a string to
func (enc *Encoder) EncodeString(s string) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeString(s)
	_, err := enc.Write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

// encodeString encodes a string to
func (enc *Encoder) encodeString(v string) ([]byte, error) {
	enc.writeByte('"')
	enc.writeStringEscape(v)
	enc.writeByte('"')
	return enc.buf, nil
}

// AppendString appends a string to the buffer
func (enc *Encoder) AppendString(v string) {
	enc.grow(len(v) + 2)
	enc.writeByte('"')
	enc.writeStringEscape(v)
	enc.writeByte('"')
}

// AddString adds a string to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddString(v string) {
	enc.grow(len(v) + 4)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(v)
	enc.writeByte('"')
}

// AddStringOmitEmpty adds a string to be encoded or skips it if it is zero value.
// Must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddStringOmitEmpty(v string) {
	if v == "" {
		return
	}
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(v)
	enc.writeByte('"')
}

// AddStringKey adds a string to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddStringKey(key, v string) {
	enc.grow(len(key) + len(v) + 5)
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyStr)
	enc.writeStringEscape(v)
	enc.writeByte('"')
}

// AddStringKeyOmitEmpty adds a string to be encoded or skips it if it is zero value.
// Must be used inside an object as it will encode a key
func (enc *Encoder) AddStringKeyOmitEmpty(key, v string) {
	if v == "" {
		return
	}
	enc.grow(len(key) + len(v) + 5)
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyStr)
	enc.writeStringEscape(v)
	enc.writeByte('"')
}
