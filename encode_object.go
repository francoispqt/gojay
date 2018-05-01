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
	enc.grow(200)
	enc.writeByte('{')
	v.MarshalObject(enc)
	enc.writeByte('}')
	return enc.buf, enc.err
}

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerObject
func (enc *Encoder) AddObject(value MarshalerObject) {
	if value.IsNil() {
		return
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	value.MarshalObject(enc)
	enc.writeByte('}')
}

// AddObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerObject
func (enc *Encoder) AddObjectKey(key string, value MarshalerObject) {
	if value.IsNil() {
		return
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.writeBytes(objKeyObj)
	value.MarshalObject(enc)
	enc.writeByte('}')
}
