package gojay

import "strconv"

// EncodeUint64 encodes an int64 to JSON
func (enc *Encoder) EncodeUint64(n uint64) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeUint64(n)
	_, err := enc.Write()
	if err != nil {
		return err
	}
	return nil
}

// encodeUint64 encodes an int to JSON
func (enc *Encoder) encodeUint64(n uint64) ([]byte, error) {
	enc.buf = strconv.AppendUint(enc.buf, n, 10)
	return enc.buf, nil
}

// AddUint64 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddUint64(v uint64) {
	enc.Uint64(v)
}

// AddUint64OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddUint64OmitEmpty(v uint64) {
	enc.Uint64OmitEmpty(v)
}

// Uint64 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Uint64(v uint64) {
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}

// Uint64OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) Uint64OmitEmpty(v uint64) {
	if v == 0 {
		return
	}
	enc.grow(10)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}

// AddUint64Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddUint64Key(key string, v uint64) {
	enc.Uint64Key(key, v)
}

// AddUint64KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddUint64KeyOmitEmpty(key string, v uint64) {
	enc.Uint64KeyOmitEmpty(key, v)
}

// Uint64Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) Uint64Key(key string, v uint64) {
	enc.grow(10 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}

// Uint64KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) Uint64KeyOmitEmpty(key string, v uint64) {
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
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}
