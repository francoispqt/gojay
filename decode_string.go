package gojay

import (
	"unsafe"
)

// DecodeString reads the next JSON-encoded value from its input and stores it in the string pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeString(v *string) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeString(v)
}
func (dec *Decoder) decodeString(v *string) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			// is string
			continue
		case '"':
			dec.cursor = dec.cursor + 1
			start, end, err := dec.getString()
			if err != nil {
				return err
			}
			// we do minus one to remove the last quote
			d := dec.data[start : end-1]
			*v = *(*string)(unsafe.Pointer(&d))
			dec.cursor = end
			return nil
		// is nil
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

func (dec *Decoder) parseEscapedString() error {
	if dec.cursor >= dec.length && !dec.read() {
		return dec.raiseInvalidJSONErr(dec.cursor)
	}

	switch dec.data[dec.cursor] {
	case '"':
		dec.data[dec.cursor] = '"'
	case '\\':
		dec.data[dec.cursor] = '\\'
	case '/':
		dec.data[dec.cursor] = '/'
	case 'b':
		dec.data[dec.cursor] = '\b'
	case 'f':
		dec.data[dec.cursor] = '\f'
	case 'n':
		dec.data[dec.cursor] = '\n'
	case 'r':
		dec.data[dec.cursor] = '\r'
	case 't':
		dec.data[dec.cursor] = '\t'
	case 'u':
		start := dec.cursor
		dec.cursor++
		str, err := dec.parseUnicode()
		if err != nil {
			return err
		}

		diff := dec.cursor - start
		dec.data = append(append(dec.data[:start-1], str...), dec.data[dec.cursor:]...)
		dec.length = len(dec.data)
		dec.cursor += len(str) - diff - 1

		return nil
	default:
		return dec.raiseInvalidJSONErr(dec.cursor)
	}

	// Truncate the previous backslash character, and the
	dec.data = append(dec.data[:dec.cursor-1], dec.data[dec.cursor:]...)
	dec.length--

	// Since we've lost a character, our dec.cursor offset is now
	// 1 past the escaped character which is precisely where we
	// want it.

	return nil
}

func (dec *Decoder) getString() (int, int, error) {
	// extract key
	var keyStart = dec.cursor
	// var str *Builder
	for dec.cursor < dec.length || dec.read() {
		switch dec.data[dec.cursor] {
		// string found
		case '"':
			dec.cursor = dec.cursor + 1
			return keyStart, dec.cursor, nil
		// slash found
		case '\\':
			dec.cursor = dec.cursor + 1
			err := dec.parseEscapedString()
			if err != nil {
				return 0, 0, err
			}
		default:
			dec.cursor = dec.cursor + 1
			continue
		}
	}
	return 0, 0, dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) skipEscapedString() error {
	start := dec.cursor
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		if dec.data[dec.cursor] != '\\' {
			d := dec.data[dec.cursor]
			dec.cursor = dec.cursor + 1
			nSlash := dec.cursor - start
			switch d {
			case '"':
				// nSlash must be odd
				if nSlash&1 != 1 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				return nil
			case 'n', 'r', 't':
				return nil
			default:
				// nSlash must be even
				if nSlash&1 == 1 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				return nil
			}
		}
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) skipString() error {
	for dec.cursor < dec.length || dec.read() {
		switch dec.data[dec.cursor] {
		// string found
		case '"':
			dec.cursor = dec.cursor + 1
			return nil
		// slash found
		case '\\':
			dec.cursor = dec.cursor + 1
			err := dec.skipEscapedString()
			if err != nil {
				return err
			}
		default:
			dec.cursor = dec.cursor + 1
			continue
		}
	}
	return dec.raiseInvalidJSONErr(len(dec.data) - 1)
}
