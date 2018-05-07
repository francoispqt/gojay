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

// EncodeFloat encodes a float64 to JSON
func (enc *Encoder) EncodeFloat(n float64) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeFloat(n)
	_, err := enc.Write()
	if err != nil {
		return err
	}
	return nil
}

// encodeFloat encodes a float64 to JSON
func (enc *Encoder) encodeFloat(n float64) ([]byte, error) {
	enc.buf = strconv.AppendFloat(enc.buf, n, 'f', -1, 64)
	return enc.buf, nil
}

// EncodeFloat32 encodes a float32 to JSON
func (enc *Encoder) EncodeFloat32(n float32) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeFloat32(n)
	_, err := enc.Write()
	if err != nil {
		return err
	}
	return nil
}

func (enc *Encoder) encodeFloat32(n float32) ([]byte, error) {
	enc.buf = strconv.AppendFloat(enc.buf, float64(n), 'f', -1, 32)
	return enc.buf, nil
}

// AddInt adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInt(v int) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// AddIntOmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddIntOmitEmpty(v int) {
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

// AddInt64 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInt64(v int64) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

// AddIntOmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddInt64OmitEmpty(v int64) {
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

// AddFloat adds a float64 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddFloat(v float64) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 64)
}

// AddFloatOmitEmpty adds a float64 to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddFloatOmitEmpty(v float64) {
	if v == 0 {
		return
	}
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 64)
}

// AddFloat32 adds a float32 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddFloat32(v float32) {
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, float64(v), 'f', -1, 32)
}

// AddFloat32OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddFloat32OmitEmpty(v float32) {
	if v == 0 {
		return
	}
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, float64(v), 'f', -1, 32)
}

// AddIntKey adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddIntKey(key string, v int) {
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

// AddIntKeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddIntKeyOmitEmpty(key string, v int) {
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

// AddInt64Key adds an int64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInt64Key(key string, v int64) {
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

// AddInt64KeyOmitEmpty adds an int64 to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddInt64KeyOmitEmpty(key string, v int64) {
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

// AddFloatKey adds a float64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloatKey(key string, value float64) {
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.grow(10)
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendFloat(enc.buf, value, 'f', -1, 64)
}

// AddFloatKeyOmitEmpty adds a float64 to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key
func (enc *Encoder) AddFloatKeyOmitEmpty(key string, v float64) {
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
	enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 64)
}

// AddFloat32Key adds a float32 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloat32Key(key string, v float32) {
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeByte('"')
	enc.writeByte(':')
	enc.buf = strconv.AppendFloat(enc.buf, float64(v), 'f', -1, 32)
}

// AddFloat32KeyOmitEmpty adds a float64 to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key
func (enc *Encoder) AddFloat32KeyOmitEmpty(key string, v float32) {
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
	enc.buf = strconv.AppendFloat(enc.buf, float64(v), 'f', -1, 32)
}
