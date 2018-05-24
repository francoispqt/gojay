package gojay

import (
	"math"
)

// DecodeUint8 reads the next JSON-encoded value from its input and stores it in the uint8 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeUint8(v *uint8) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeUint8(v)
}

func (dec *Decoder) decodeUint8(v *uint8) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getUint8()
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-': // if negative, we just set it to 0
			err := dec.skipData()
			if err != nil {
				return err
			}
			*v = 0
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
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
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) getUint8() (uint8, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case ' ', '\n', '\t', '\r':
			continue
		case '.', ',', '}', ']':
			dec.cursor = j
			return dec.atoui8(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoui8(start, end), nil
}

// DecodeUint16 reads the next JSON-encoded value from its input and stores it in the uint16 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeUint16(v *uint16) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeUint16(v)
}

func (dec *Decoder) decodeUint16(v *uint16) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getUint16()
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			err := dec.skipData()
			if err != nil {
				return err
			}
			*v = 0
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
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
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) getUint16() (uint16, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case ' ', '\n', '\t', '\r':
			continue
		case '.', ',', '}', ']':
			dec.cursor = j
			return dec.atoui16(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoui16(start, end), nil
}

// DecodeUint32 reads the next JSON-encoded value from its input and stores it in the uint32 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeUint32(v *uint32) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeUint32(v)
}

func (dec *Decoder) decodeUint32(v *uint32) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getUint32()
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			err := dec.skipData()
			if err != nil {
				return err
			}
			*v = 0
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
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
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) getUint32() (uint32, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case ' ', '\n', '\t', '\r':
			continue
		case '.', ',', '}', ']':
			dec.cursor = j
			return dec.atoui32(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoui32(start, end), nil
}

// DecodeUint64 reads the next JSON-encoded value from its input and stores it in the uint64 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeUint64(v *uint64) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeUint64(v)
}
func (dec *Decoder) decodeUint64(v *uint64) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getUint64()
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			err := dec.skipData()
			if err != nil {
				return err
			}
			*v = 0
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
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
	return dec.raiseInvalidJSONErr(dec.cursor)
}

func (dec *Decoder) getUint64() (uint64, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case ' ', '\n', '\t', '\r', '.', ',', '}', ']':
			dec.cursor = j
			return dec.atoui64(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoui64(start, end), nil
}

func (dec *Decoder) atoui64(start, end int) uint64 {
	var ll = end + 1 - start
	var val = uint64(digits[dec.data[start]])
	end = end + 1
	if ll < maxUint64Length {
		for i := start + 1; i < end; i++ {
			uintv := uint64(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + uintv
		}
	} else if ll == maxUint64Length {
		for i := start + 1; i < end; i++ {
			uintv := uint64(digits[dec.data[i]])
			if val > maxUint64toMultiply {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val = (val << 3) + (val << 1)
			if math.MaxUint64-val < uintv {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val += uintv
		}
	} else {
		dec.err = dec.makeInvalidUnmarshalErr(val)
		return 0
	}
	return val
}

func (dec *Decoder) atoui32(start, end int) uint32 {
	var ll = end + 1 - start
	var val uint32
	val = uint32(digits[dec.data[start]])
	end = end + 1
	if ll < maxUint32Length {
		for i := start + 1; i < end; i++ {
			uintv := uint32(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + uintv
		}
	} else if ll == maxUint32Length {
		for i := start + 1; i < end; i++ {
			uintv := uint32(digits[dec.data[i]])
			if val > maxUint32toMultiply {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val = (val << 3) + (val << 1)
			if math.MaxUint32-val < uintv {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val += uintv
		}
	} else if ll > maxUint32Length {
		dec.err = dec.makeInvalidUnmarshalErr(val)
		val = 0
	}
	return val
}

func (dec *Decoder) atoui16(start, end int) uint16 {
	var ll = end + 1 - start
	var val uint16
	val = uint16(digits[dec.data[start]])
	end = end + 1
	if ll < maxUint16Length {
		for i := start + 1; i < end; i++ {
			uintv := uint16(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + uintv
		}
	} else if ll == maxUint16Length {
		for i := start + 1; i < end; i++ {
			uintv := uint16(digits[dec.data[i]])
			if val > maxUint16toMultiply {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val = (val << 3) + (val << 1)
			if math.MaxUint16-val < uintv {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val += uintv
		}
	} else if ll > maxUint16Length {
		dec.err = dec.makeInvalidUnmarshalErr(val)
		val = 0
	}
	return val
}

func (dec *Decoder) atoui8(start, end int) uint8 {
	var ll = end + 1 - start
	var val uint8
	val = uint8(digits[dec.data[start]])
	end = end + 1
	if ll < maxUint8Length {
		for i := start + 1; i < end; i++ {
			uintv := uint8(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + uintv
		}
	} else if ll == maxUint8Length {
		for i := start + 1; i < end; i++ {
			uintv := uint8(digits[dec.data[i]])
			if val > maxUint8toMultiply {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val = (val << 3) + (val << 1)
			if math.MaxUint8-val < uintv {
				dec.err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val += uintv
		}
	} else if ll > maxUint8Length {
		dec.err = dec.makeInvalidUnmarshalErr(val)
		val = 0
	}
	return val
}
