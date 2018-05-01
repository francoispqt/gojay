package gojay

import "strconv"

// EncodeInt encodes an int to JSON
func (enc *Encoder) EncodeInt(n int64) ([]byte, error) {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	return enc.encodeInt(n)
}

// encodeInt encodes an int to JSON
func (enc *Encoder) encodeInt(n int64) ([]byte, error) {
	s := strconv.Itoa(int(n))
	enc.writeString(s)
	return enc.buf, nil
}

// EncodeFloat encodes a float64 to JSON
func (enc *Encoder) EncodeFloat(n float64) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, err := enc.encodeFloat(n)
	if err != nil {
		return err
	}
	_, err = enc.write()
	if err != nil {
		return err
	}
	return nil
}

// encodeFloat encodes a float64 to JSON
func (enc *Encoder) encodeFloat(n float64) ([]byte, error) {
	enc.buf = strconv.AppendFloat(enc.buf, float64(n), 'f', -1, 64)
	return enc.buf, nil
}

// EncodeFloat encodes a float32 to JSON
func (enc *Encoder) EncodeFloat32(n float32) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, err := enc.encodeFloat32(n)
	if err != nil {
		return err
	}
	_, err = enc.write()
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
func (enc *Encoder) AddInt(value int) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendInt(enc.buf, int64(value), 10)
	return nil
}

// AddFloat adds a float64 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddFloat(value float64) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendFloat(enc.buf, value, 'f', -1, 64)

	return nil
}

// AddIntKey adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddIntKey(key string, value int) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendInt(enc.buf, int64(value), 10)

	return nil
}

// AddFloatKey adds a float64 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloatKey(key string, value float64) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendFloat(enc.buf, value, 'f', -1, 64)

	return nil
}

// AddFloat32Key adds a float32 to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddFloat32Key(key string, value float32) error {
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeByte('"')
	enc.writeByte(':')
	enc.buf = strconv.AppendFloat(enc.buf, float64(value), 'f', -1, 32)

	return nil
}
