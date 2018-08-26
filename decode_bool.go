package gojay

// DecodeBool reads the next JSON-encoded value from its input and stores it in the boolean pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeBool(v *bool) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeBool(v)
}
func (dec *Decoder) decodeBool(v *bool) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case 't':
			dec.cursor++
			err := dec.assertTrue()
			if err != nil {
				return err
			}
			*v = true
			dec.cursor++
			return nil
		case 'f':
			dec.cursor++
			err := dec.assertFalse()
			if err != nil {
				return err
			}
			*v = false
			dec.cursor++
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
			*v = false
			dec.cursor++
			return nil
		default:
			dec.err = dec.makeInvalidUnmarshalErr(v)
			err := dec.skipData()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
func (dec *Decoder) decodeBoolNull(v **bool) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case 't':
			dec.cursor++
			err := dec.assertTrue()
			if err != nil {
				return err
			}
			if *v == nil {
				*v = new(bool)
			}
			**v = true
			dec.cursor++
			return nil
		case 'f':
			dec.cursor++
			err := dec.assertFalse()
			if err != nil {
				return err
			}
			if *v == nil {
				*v = new(bool)
			}
			**v = false
			dec.cursor++
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
			dec.cursor++
			return nil
		default:
			dec.err = dec.makeInvalidUnmarshalErr(v)
			err := dec.skipData()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (dec *Decoder) assertTrue() error {
	i := 0
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.data[dec.cursor] != 'r' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.data[dec.cursor] != 'u' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.data[dec.cursor] != 'e' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
			switch dec.data[dec.cursor] {
			case ' ', '\t', '\n', ',', ']', '}':
				dec.cursor--
				return nil
			default:
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		}
		i++
	}
	if i == 3 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) assertNull() error {
	i := 0
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.data[dec.cursor] != 'u' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
			switch dec.data[dec.cursor] {
			case ' ', '\t', '\n', ',', ']', '}':
				dec.cursor--
				return nil
			default:
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		}
		i++
	}
	if i == 3 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) assertFalse() error {
	i := 0
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch i {
		case 0:
			if dec.data[dec.cursor] != 'a' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 1:
			if dec.data[dec.cursor] != 'l' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 2:
			if dec.data[dec.cursor] != 's' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 3:
			if dec.data[dec.cursor] != 'e' {
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		case 4:
			switch dec.data[dec.cursor] {
			case ' ', '\t', '\n', ',', ']', '}':
				dec.cursor--
				return nil
			default:
				return dec.raiseInvalidJSONErr(dec.cursor)
			}
		}
		i++
	}
	if i == 4 {
		return nil
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}
