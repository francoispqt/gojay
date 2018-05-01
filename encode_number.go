package gojay

import "strconv"

// EncodeInt encodes an int to JSON
func (enc *Encoder) EncodeInt(n int) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeInt(n)
	_, err := enc.write()
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
	_, err := enc.write()
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
	_, err := enc.write()
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
	_, err := enc.write()
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
func (enc *Encoder) AddInt(value int) {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, int64(value), 10)
}

// AddFloat adds a float64 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddFloat(value float64) {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, value, 'f', -1, 64)
}

// AddFloat32 adds a float32 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddFloat32(value float32) {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, float64(value), 'f', -1, 32)
}

// AddIntKey adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddIntKey(key string, value int) {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, int64(value), 10)
}

// AddFloatKey adds a float64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloatKey(key string, value float64) {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendFloat(enc.buf, value, 'f', -1, 64)
}

// AddFloat32Key adds a float32 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloat32Key(key string, value float32) {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeByte('"')
	enc.writeByte(':')
	enc.buf = strconv.AppendFloat(enc.buf, float64(value), 'f', -1, 32)
}
