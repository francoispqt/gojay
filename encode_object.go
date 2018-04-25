package gojay

var objKeyStr = []byte(`":"`)
var objKeyObj = []byte(`":{`)
var objKeyArr = []byte(`":[`)
var objKey = []byte(`":`)

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement Marshaler
func (enc *Encoder) AddObject(value MarshalerObject) error {
	if value.IsNil() {
		return nil
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	value.MarshalObject(enc)
	enc.writeByte('}')
	return nil
}

// AddObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement Marshaler
func (enc *Encoder) AddObjectKey(key string, value MarshalerObject) error {
	if value.IsNil() {
		return nil
	}
	r, ok := enc.getPreviousRune()
	if ok && r != '{' && r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeString(key)
	enc.write(objKeyObj)
	value.MarshalObject(enc)
	enc.writeByte('}')
	return nil
}
