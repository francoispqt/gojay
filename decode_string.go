package gojay

import (
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
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
				dec.cursor = dec.cursor - diff
				return nil
			case 'b':
				// number of slash must be even
				// if is odd number of slashes
				// divide nSlash - 1 by 2 and leave last one
				// else divide nSlash by 2 and leave the letter
				if nSlash&1 != 0 {
					return InvalidJSONError("Invalid JSON unescaped character")
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
					return InvalidJSONError("Invalid JSON unescaped character")
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
					return InvalidJSONError("Invalid JSON unescaped character")
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
					return InvalidJSONError("Invalid JSON unescaped character")
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
					return InvalidJSONError("Invalid JSON unescaped character")
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

func (dec *Decoder) getUnicode() (rune, error) {
	i := 0
	r := rune(0)
	for ; (dec.cursor < dec.length || dec.read()) && i < 4; dec.cursor++ {
		c := dec.data[dec.cursor]
		if c >= '0' && c <= '9' {
			r = r*16 + rune(c-'0')
		} else if c >= 'a' && c <= 'f' {
			r = r*16 + rune(c-'a'+10)
		} else if c >= 'A' && c <= 'F' {
			r = r*16 + rune(c-'A'+10)
		} else {
			return 0, InvalidJSONError("Invalid unicode code point")
		}
		i++
	}
	return r, nil
}

func (dec *Decoder) appendEscapeChar(str []byte, c byte) ([]byte, error) {
	switch c {
	case 't':
		str = append(str, '\t')
	case 'n':
		str = append(str, '\n')
	case 'r':
		str = append(str, '\r')
	case 'b':
		str = append(str, '\b')
	case 'f':
		str = append(str, '\f')
	case '\\':
		str = append(str, '\\')
	default:
		return nil, InvalidJSONError("Invalid JSON")
	}
	return str, nil
}

func (dec *Decoder) parseUnicode() ([]byte, error) {
	// get unicode after u
	r, err := dec.getUnicode()
	if err != nil {
		return nil, err
	}
	// no error start making new string
	str := make([]byte, 16, 16)
	i := 0
	if utf16.IsSurrogate(r) {
		if dec.cursor < dec.length || dec.read() {
			c := dec.data[dec.cursor]
			if c != '\\' {
				i += utf8.EncodeRune(str, r)
				return str[:i], nil
			}
			dec.cursor++
			if dec.cursor >= dec.length && !dec.read() {
				return nil, InvalidJSONError("Invalid JSON")
			}
			c = dec.data[dec.cursor]
			if c != 'u' {
				i += utf8.EncodeRune(str, r)
				str, err = dec.appendEscapeChar(str[:i], c)
				if err != nil {
					dec.err = err
					return nil, err
				}
				i++
				dec.cursor++
				return str[:i], nil
			}
			dec.cursor++
			r2, err := dec.getUnicode()
			if err != nil {
				return nil, err
			}
			combined := utf16.DecodeRune(r, r2)
			if combined == '\uFFFD' {
				i += utf8.EncodeRune(str, r)
				i += utf8.EncodeRune(str, r2)
			} else {
				i += utf8.EncodeRune(str, combined)
			}
		}
		return str[:i], nil
	}
	i += utf8.EncodeRune(str, r)
	return str[:i], nil
}
