package gojay

import "strconv"

// EncodeUint32 encodes an int64 to JSON
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

// AddUint32 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddUint32(v uint32) {
	enc.Uint64(uint64(v))
}

// AddUint32OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddUint32OmitEmpty(v uint32) {
	enc.Uint64OmitEmpty(uint64(v))
}

// Uint32 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Uint32(v uint32) {
	enc.Uint64(uint64(v))
}

// Uint32OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) Uint32OmitEmpty(v uint32) {
	enc.Uint64OmitEmpty(uint64(v))
}

// AddUint32Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddUint32Key(key string, v uint32) {
	enc.Uint64Key(key, uint64(v))
}

// AddUint32KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddUint32KeyOmitEmpty(key string, v uint32) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}

// Uint32Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) Uint32Key(key string, v uint32) {
	enc.Uint64Key(key, uint64(v))
}

// Uint32KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) Uint32KeyOmitEmpty(key string, v uint32) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}

// AddUint16 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddUint16(v uint16) {
	enc.Uint64(uint64(v))
}

// AddUint16OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddUint16OmitEmpty(v uint16) {
	enc.Uint64OmitEmpty(uint64(v))
}

// Uint16 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Uint16(v uint16) {
	enc.Uint64(uint64(v))
}

// Uint16OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) Uint16OmitEmpty(v uint16) {
	enc.Uint64OmitEmpty(uint64(v))
}

// AddUint16Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddUint16Key(key string, v uint16) {
	enc.Uint64Key(key, uint64(v))
}

// AddUint16KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddUint16KeyOmitEmpty(key string, v uint16) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}

// Uint16Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) Uint16Key(key string, v uint16) {
	enc.Uint64Key(key, uint64(v))
}

// Uint16KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) Uint16KeyOmitEmpty(key string, v uint16) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}

// AddUint8 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddUint8(v uint8) {
	enc.Uint64(uint64(v))
}

// AddUint8OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) AddUint8OmitEmpty(v uint8) {
	enc.Uint64OmitEmpty(uint64(v))
}

// Uint8 adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) Uint8(v uint8) {
	enc.Uint64(uint64(v))
}

// Uint8OmitEmpty adds an int to be encoded and skips it if its value is 0,
// must be used inside a slice or array encoding (does not encode a key).
func (enc *Encoder) Uint8OmitEmpty(v uint8) {
	enc.Uint64OmitEmpty(uint64(v))
}

// AddUint8Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddUint8Key(key string, v uint8) {
	enc.Uint64Key(key, uint64(v))
}

// AddUint8KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) AddUint8KeyOmitEmpty(key string, v uint8) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}

// Uint8Key adds an int to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) Uint8Key(key string, v uint8) {
	enc.Uint64Key(key, uint64(v))
}

// Uint8KeyOmitEmpty adds an int to be encoded and skips it if its value is 0.
// Must be used inside an object as it will encode a key.
func (enc *Encoder) Uint8KeyOmitEmpty(key string, v uint8) {
	enc.Uint64KeyOmitEmpty(key, uint64(v))
}
