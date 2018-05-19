package gojay

var objKeyStr = []byte(`":"`)
var objKeyObj = []byte(`":{`)
var objKeyArr = []byte(`":[`)
var objKey = []byte(`":`)

// EncodeObject encodes an object to JSON
func (enc *Encoder) EncodeObject(v MarshalerJSONObject) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, err := enc.encodeObject(v)
	if err != nil {
		enc.err = err
		return err
	}
	_, err = enc.Write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

func (enc *Encoder) encodeObject(v MarshalerJSONObject) ([]byte, error) {
	enc.grow(500)
	enc.writeByte('{')
	if !v.IsNil() {
		v.MarshalJSONObject(enc)
	}
	enc.writeByte('}')
	return enc.buf, enc.err
}

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObject(v MarshalerJSONObject) {
	enc.Object(v)
}

// AddObjectOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectOmitEmpty(v MarshalerJSONObject) {
	enc.ObjectOmitEmpty(v)
}

// AddObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectKey(key string, v MarshalerJSONObject) {
	enc.ObjectKey(key, v)
}

// AddObjectKeyOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectKeyOmitEmpty(key string, v MarshalerJSONObject) {
	enc.ObjectKeyOmitEmpty(key, v)
}

// Object adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) Object(v MarshalerJSONObject) {
	if v.IsNil() {
		enc.grow(2)
		r := enc.getPreviousRune()
		if r != '{' && r != '[' {
			enc.writeByte(',')
		}
		enc.writeByte('{')
		enc.writeByte('}')
		return
	}
	enc.grow(4)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	v.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectOmitEmpty(v MarshalerJSONObject) {
	if v.IsNil() {
		return
	}
	enc.grow(2)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	v.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectKey(key string, value MarshalerJSONObject) {
	if value.IsNil() {
		enc.grow(2 + len(key))
		r := enc.getPreviousRune()
		if r != '{' {
			enc.writeByte(',')
		}
		enc.writeByte('"')
		enc.writeStringEscape(key)
		enc.writeBytes(objKeyObj)
		enc.writeByte('}')
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectKeyOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectKeyOmitEmpty(key string, value MarshalerJSONObject) {
	if value.IsNil() {
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// EncodeObjectFunc is a custom func type implementing MarshaleObject.
// Use it to cast a func(*Encoder) to Marshal an object.
//
//	enc := gojay.NewEncoder(io.Writer)
//	enc.EncodeObject(gojay.EncodeObjectFunc(func(enc *gojay.Encoder) {
//		enc.AddStringKey("hello", "world")
//	}))
type EncodeObjectFunc func(*Encoder)

// MarshalJSONObject implements MarshalerJSONObject.
func (f EncodeObjectFunc) MarshalJSONObject(enc *Encoder) {
	f(enc)
}

// IsNil implements MarshalerJSONObject.
func (f EncodeObjectFunc) IsNil() bool {
	return f == nil
}
