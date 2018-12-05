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
	// know where to stop slash
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
				diff := (nSlash - 1) >> 1
				dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff
				return nil
			case 'u':
				if nSlash&1 == 0 {
					diff := nSlash >> 1
					dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
					dec.length = len(dec.data)
					dec.cursor -= nSlash - diff
					return nil
				}
				start := dec.cursor - 2 - ((nSlash - 1) >> 1)
				str, err := dec.parseUnicode()
				if err != nil {
					dec.err = err
					return err
				}
				diff := dec.cursor - start
				dec.data = append(append(dec.data[:start], str...), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor = dec.cursor - diff + len(str)
				return nil
			case 'b':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				var diff int
				diff = nSlash >> 1
				dec.data = append(append(dec.data[:start+diff-2], '\b'), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff + 1
				return nil
			case 'f':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				var diff int
				diff = nSlash >> 1
				dec.data = append(append(dec.data[:start+diff-2], '\f'), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff + 1
				return nil
			case 'n':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				var diff int
				diff = nSlash >> 1
				dec.data = append(append(dec.data[:start+diff-2], '\n'), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff + 1
				return nil
			case 'r':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				var diff int
				diff = nSlash >> 1
				dec.data = append(append(dec.data[:start+diff-2], '\r'), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff + 1
				return nil
			case 't':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				var diff int
				diff = nSlash >> 1
				dec.data = append(append(dec.data[:start+diff-2], '\t'), dec.data[dec.cursor:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff + 1
				return nil
			default:
				// nSlash must be even
				if nSlash&1 == 1 {
					return dec.raiseInvalidJSONErr(dec.cursor)
				}
				diff := nSlash >> 1
				dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
				dec.length = len(dec.data)
				dec.cursor -= (nSlash - diff)
				return nil
			}
		}
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
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
			case 'u': // is unicode, we skip the following characters and place the cursor one one byte backward to avoid it breaking when returning to skipString
				if err := dec.skipString(); err != nil {
					return err
				}
				dec.cursor--
				return nil
			case 'n', 'r', 't', '/', 'f', 'b':
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
		// found the closing quote
		// let's return
		case '"':
			dec.cursor = dec.cursor + 1
			return nil
		// solidus found start parsing an escaped string
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
