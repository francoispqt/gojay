package gojay

// AddInterface adds an interface{} to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddInterface(value interface{}) error {
	switch value.(type) {
	case string:
		return enc.AddString(value.(string))
	case bool:
		return enc.AddBool(value.(bool))
	case MarshalerArray:
		return enc.AddArray(value.(MarshalerArray))
	case MarshalerObject:
		return enc.AddObject(value.(MarshalerObject))
	case int:
		return enc.AddInt(value.(int))
	case int64:
		return enc.AddInt(int(value.(int64)))
	case int32:
		return enc.AddInt(int(value.(int32)))
	case int8:
		return enc.AddInt(int(value.(int8)))
	case uint64:
		return enc.AddInt(int(value.(uint64)))
	case uint32:
		return enc.AddInt(int(value.(uint32)))
	case uint16:
		return enc.AddInt(int(value.(uint16)))
	case uint8:
		return enc.AddInt(int(value.(uint8)))
	case float64:
		return enc.AddFloat(value.(float64))
	case float32:
		return enc.AddFloat(float64(value.(float32)))
	}

	return nil
}

// AddInterfaceKey adds an interface{} to be encoded, must be used inside an object as it will encode a key
func (enc *Encoder) AddInterfaceKey(key string, value interface{}) error {
	switch value.(type) {
	case string:
		return enc.AddStringKey(key, value.(string))
	case bool:
		return enc.AddBoolKey(key, value.(bool))
	case MarshalerArray:
		return enc.AddArrayKey(key, value.(MarshalerArray))
	case MarshalerObject:
		return enc.AddObjectKey(key, value.(MarshalerObject))
	case int:
		return enc.AddIntKey(key, value.(int))
	case int64:
		return enc.AddIntKey(key, int(value.(int64)))
	case int32:
		return enc.AddIntKey(key, int(value.(int32)))
	case int16:
		return enc.AddIntKey(key, int(value.(int16)))
	case int8:
		return enc.AddIntKey(key, int(value.(int8)))
	case uint64:
		return enc.AddIntKey(key, int(value.(uint64)))
	case uint32:
		return enc.AddIntKey(key, int(value.(uint32)))
	case uint16:
		return enc.AddIntKey(key, int(value.(uint16)))
	case uint8:
		return enc.AddIntKey(key, int(value.(uint8)))
	case float64:
		return enc.AddFloatKey(key, value.(float64))
	case float32:
		return enc.AddFloat32Key(key, value.(float32))
	}

	return nil
}
