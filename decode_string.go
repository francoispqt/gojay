package gojay

import (
	"fmt"
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
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to string, wrong char '%s' found at pos %d",
					string(dec.data[dec.cursor]),
					dec.cursor,
				),
			)
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
					return InvalidJSONError("Invalid JSON unescaped character")
				}
				diff := (nSlash - 1) >> 1
				dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff
				return nil
			case 'n', 'r', 't':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				var diff int
				if nSlash&1 == 1 {
					diff = (nSlash - 1) >> 1
					dec.data = append(dec.data[:start+diff], dec.data[dec.cursor-1:]...)
				} else {
					diff = nSlash >> 1
					dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
				}
				dec.length = len(dec.data)
				dec.cursor -= nSlash - diff
				return nil
			default:
				// nSlash must be even
				if nSlash&1 == 1 {
					return InvalidJSONError("Invalid JSON unescaped character")
				}
				diff := nSlash >> 1
				dec.data = append(dec.data[:start+diff-1], dec.data[dec.cursor-1:]...)
				dec.length = len(dec.data)
				dec.cursor -= (nSlash - diff)
				return nil
			}
		}
	}
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
	return 0, 0, InvalidJSONError("Invalid JSON while parsing string")
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
					return InvalidJSONError("Invalid JSON unescaped character")
				}
				return nil
			case 'n', 'r', 't':
				return nil
			default:
				// nSlash must be even
				if nSlash&1 == 1 {
					return InvalidJSONError("Invalid JSON unescaped character")
				}
				return nil
			}
		}
	}
	return InvalidJSONError("Invalid JSON")
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
	return InvalidJSONError("Invalid JSON while parsing string")
}
