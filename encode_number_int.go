package gojay

import "strconv"

// EncodeInt encodes an int to JSON
func (enc *Encoder) EncodeInt(n int) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeInt(n)
	_, err := enc.Write()
	if err != nil {
		return err
	}
	return nil
}

// encodeInt encodes an int to JSON
func (enc *Encoder) encodeInt(n int) ([]byte, error) {
	enc.buf = strconv.AppendInt(enc.buf, int64(n), 10)
	return enc.buf, nil
}

// EncodeInt64 encodes an int64 to JSON
func (enc *Encoder) EncodeInt64(n int64) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeInt64(n)
	_, err := enc.Write()
	if err != nil {
		return err
	}
	return nil
}

// encodeInt64 encodes an int to JSON
func (enc *Encoder) encodeInt64(n int64) ([]byte, error) {
	enc.buf = strconv.AppendInt(enc.buf, n, 10)
	return enc.buf, nil
}

// AddInt adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInt(v int) {
	enc.Int(v)
}

// AddIntOmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddIntOmitEmpty(v int) {
	enc.IntOmitEmpty(v)
}

// Int adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Int(v int) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// IntOmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) IntOmitEmpty(v int) {
	if v == 0 {
		return
	}
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// AddIntKey adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddIntKey(key string, v int) {
	enc.IntKey(key, v)
}

// AddIntKeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddIntKeyOmitEmpty(key string, v int) {
	enc.IntKeyOmitEmpty(key, v)
}

// IntKey adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) IntKey(key string, v int) {
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// IntKeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) IntKeyOmitEmpty(key string, v int) {
	if v == 0 {
		return
	}
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// AddInt64 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInt64(v int64) {
	enc.Int64(v)
}

// AddInt64OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddInt64OmitEmpty(v int64) {
	enc.Int64OmitEmpty(v)
}

// Int64 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Int64(v int64) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

// Int64OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) Int64OmitEmpty(v int64) {
	if v == 0 {
		return
	}
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

// AddInt64Key adds an int64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInt64Key(key string, v int64) {
	enc.Int64Key(key, v)
}

// AddInt64KeyOmitEmpty adds an int64 to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddInt64KeyOmitEmpty(key string, v int64) {
	enc.Int64KeyOmitEmpty(key, v)
}

// Int64Key adds an int64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) Int64Key(key string, v int64) {
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

// Int64KeyOmitEmpty adds an int64 to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) Int64KeyOmitEmpty(key string, v int64) {
	if v == 0 {
		return
	}
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}
