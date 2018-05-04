package gojay

var objKeyStr = []byte(`":"`)
var objKeyObj = []byte(`":{`)
var objKeyArr = []byte(`":[`)
var objKey = []byte(`":`)

// EncodeObject encodes an object to JSON
func (enc *Encoder) EncodeObject(v MarshalerObject) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, err := enc.encodeObject(v)
	if err != nil {
		enc.err = err
		return err
	}
	_, err = enc.write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

func (enc *Encoder) encodeObject(v MarshalerObject) ([]byte, error) {
	enc.grow(500)
	enc.writeByte('{')
	if !v.IsNil() {
		v.MarshalObject(enc)
	}
	enc.writeByte('}')
	return enc.buf, enc.err
}

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerObject
func (enc *Encoder) AddObject(v MarshalerObject) {
	if v.IsNil() {
		enc.grow(2)
		r, ok := enc.getPreviousRune()
		if ok && r != '{' && r != '[' {
			enc.writeByte(',')
		}
		enc.writeByte('{')
		enc.writeByte('}')
		return
	}
	enc.grow(4)
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	v.MarshalObject(enc)
	enc.writeByte('}')
}

// AddObjectOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerObject
func (enc *Encoder) AddObjectOmitEmpty(v MarshalerObject) {
	if v.IsNil() {
		return
	}
	enc.grow(2)
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	v.MarshalObject(enc)
	enc.writeByte('}')
}

// AddObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerObject
func (enc *Encoder) AddObjectKey(key string, value MarshalerObject) {
	if value.IsNil() {
		enc.grow(2 + len(key))
		r, ok := enc.getPreviousRune()
		if ok && r != '{' {
			enc.writeByte(',')
		}
		enc.writeByte('"')
		enc.writeStringEscape(key)
		enc.writeBytes(objKeyObj)
		enc.writeByte('}')
		return
	}
	enc.grow(5 + len(key))
	r, ok := enc.getPreviousRune()
	if ok && r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalObject(enc)
	enc.writeByte('}')
}

// AddObjectKeyOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerObject
func (enc *Encoder) AddObjectKeyOmitEmpty(key string, value MarshalerObject) {
	if value.IsNil() {
		return
	}
	enc.grow(5 + len(key))
	r, ok := enc.getPreviousRune()
	if ok && r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalObject(enc)
	enc.writeByte('}')
}

// EncodeObjectFunc is a custom func type implementating MarshaleObject.
// Use it to cast a func(*Encoder) to Marshal and object.
//
//	enc := gojay.NewEncoder(io.Writer)
//	enc.EncoderObject(gojay.EncodeObjectFunc(func(enc *gojay.Encoder) {
//		enc.AddStringKey("hello", "world")
//	}))
type EncodeObjectFunc func(*Encoder)

// MarshalObject implements MarshalerObject.
func (f EncodeObjectFunc) MarshalObject(enc *Encoder) {
	f(enc)
}

// IsNil implements MarshalerObject.
func (f EncodeObjectFunc) IsNil() bool {
	return f == nil
}
