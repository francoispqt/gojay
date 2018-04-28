package gojay

import (
	"fmt"
)

var digits []int8

const maxUint32 = uint32(0xffffffff)
const maxUint64 = uint64(0xffffffffffffffff)
const maxInt32 = int32(0x7fffffff)
const maxInt64 = int64(0x7fffffffffffffff)
const maxInt64toMultiply = int64(0x7fffffffffffffff) / 10
const maxInt32toMultiply = int32(0x7fffffff) / 10
const maxUint32toMultiply = uint32(0xffffffff) / 10
const maxUint64toMultiply = uint64(0xffffffffffffffff) / 10
const maxUint32Length = 10
const maxUint64Length = 20
const maxInt32Length = 10
const maxInt64Length = 19
const invalidNumber = int8(-1)

var pow10uint64 = [20]uint64{
	0,
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
}

func init() {
	digits = make([]int8, 256)
	for i := 0; i < len(digits); i++ {
		digits[i] = invalidNumber
	}
	for i := int8('0'); i <= int8('9'); i++ {
		digits[i] = i - int8('0')
	}
}

// DecodeInt reads the next JSON-encoded value from its input and stores it in the int pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeInt(v *int) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeInt(v)
}
func (dec *Decoder) decodeInt(v *int) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getInt64(c)
			if err != nil {
				return err
			}
			*v = int(val)
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getInt64(dec.data[dec.cursor])
			if err != nil {
				return err
			}
			*v = -int(val)
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing int")
}

// DecodeInt32 reads the next JSON-encoded value from its input and stores it in the int32 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeInt32(v *int32) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeInt32(v)
}
func (dec *Decoder) decodeInt32(v *int32) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getInt32(c)
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getInt32(dec.data[dec.cursor])
			if err != nil {
				return err
			}
			*v = -val
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing int")
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
			val, err := dec.getUint32(c)
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getUint32(dec.data[dec.cursor])
			if err != nil {
				return err
			}
			// unsigned int so we don't bother with the sign
			*v = val
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing int")
}

// DecodeInt64 reads the next JSON-encoded value from its input and stores it in the int64 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeInt64(v *int64) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeInt64(v)
}

func (dec *Decoder) decodeInt64(v *int64) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getInt64(c)
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getInt64(dec.data[dec.cursor])
			if err != nil {
				return err
			}
			*v = -val
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing int")
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
			val, err := dec.getUint64(c)
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getUint64(dec.data[dec.cursor])
			if err != nil {
				return err
			}
			// unsigned int so we don't bother with the sign
			*v = val
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing int")
}

// DecodeFloat64 reads the next JSON-encoded value from its input and stores it in the float64 pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeFloat64(v *float64) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeFloat64(v)
}
func (dec *Decoder) decodeFloat64(v *float64) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch c := dec.data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getFloat(c)
			if err != nil {
				return err
			}
			*v = val
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getFloat(c)
			if err != nil {
				return err
			}
			*v = -val
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to float, wrong char '%s' found at pos %d",
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
	return InvalidJSONError("Invalid JSON while parsing float")
}

func (dec *Decoder) skipNumber() (int, error) {
	end := dec.cursor + 1
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j + 1
			continue
		case '.':
			end = j + 1
			continue
		case ',', '}', ']':
			return end, nil
		case ' ', '\n', '\t', '\r':
			continue
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return end, InvalidJSONError("Invalid JSON while parsing number")
	}
	return end, nil
}

func (dec *Decoder) getInt64(b byte) (int64, error) {
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
			return dec.atoi64(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, InvalidJSONError("Invalid JSON while parsing number")
	}
	return dec.atoi64(start, end), nil
}

func (dec *Decoder) getUint64(b byte) (uint64, error) {
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
			return dec.atoui64(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, InvalidJSONError("Invalid JSON while parsing number")
	}
	return dec.atoui64(start, end), nil
}

func (dec *Decoder) getInt32(b byte) (int32, error) {
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
			return dec.atoi32(start, end), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, InvalidJSONError("Invalid JSON while parsing number")
	}
	return dec.atoi32(start, end), nil
}

func (dec *Decoder) getUint32(b byte) (uint32, error) {
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
		return 0, InvalidJSONError("Invalid JSON while parsing number")
	}
	return dec.atoui32(start, end), nil
}

func (dec *Decoder) getFloat(b byte) (float64, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case '.':
			// we get part before decimal as integer
			beforeDecimal := dec.atoi64(start, end)
			// then we get part after decimal as integer
			start = j + 1
			// get number after the decimal point
			// multiple the before decimal point portion by 10 using bitwise
			for i := j + 1; i < dec.length || dec.read(); i++ {
				c := dec.data[i]
				if isDigit(c) {
					end = i
					beforeDecimal = (beforeDecimal << 3) + (beforeDecimal << 1)
					continue
				}
				dec.cursor = i
				break
			}
			// then we add both integers
			// then we divide the number by the power found
			afterDecimal := dec.atoi64(start, end)
			pow := pow10uint64[end-start+2]
			return float64(beforeDecimal+afterDecimal) / float64(pow), nil
		case ' ', '\n', '\t', '\r':
			continue
		case ',', '}', ']': // does not have decimal
			dec.cursor = j
			return float64(dec.atoi64(start, end)), nil
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, InvalidJSONError("Invalid JSON while parsing number")
	}
	return float64(dec.atoi64(start, end)), nil
}

func (dec *Decoder) atoi64(start, end int) int64 {
	var ll = end + 1 - start
	var val = int64(digits[dec.data[start]])
	end = end + 1
	if ll < maxInt64Length {
		for i := start + 1; i < end; i++ {
			intv := int64(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + intv
		}
		return val
	} else if ll == maxInt64Length {
		for i := start + 1; i < end; i++ {
			intv := int64(digits[dec.data[i]])
			if val > maxInt64toMultiply {
				dec.err = InvalidTypeError("Overflows int64")
				return 0
			}
			val = (val << 3) + (val << 1)
			if maxInt64-val < intv {
				dec.err = InvalidTypeError("Overflows int64")
				return 0
			}
			val += intv
		}
	} else {
		dec.err = InvalidTypeError("Overflows int64")
		return 0
	}
	return val
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
				dec.err = InvalidTypeError("Overflows uint64")
				return 0
			}
			val = (val << 3) + (val << 1)
			if maxUint64-val < uintv {
				dec.err = InvalidTypeError("Overflows uint64")
				return 0
			}
			val += uintv
		}
	} else {
		dec.err = InvalidTypeError("Overflows uint64")
		return 0
	}
	return val
}

func (dec *Decoder) atoi32(start, end int) int32 {
	var ll = end + 1 - start
	var val = int32(digits[dec.data[start]])
	end = end + 1
	// overflowing
	if ll < maxInt32Length {
		for i := start + 1; i < end; i++ {
			intv := int32(digits[dec.data[i]])
			val = (val << 3) + (val << 1) + intv
		}
	} else if ll == maxInt32Length {
		for i := start + 1; i < end; i++ {
			intv := int32(digits[dec.data[i]])
			if val > maxInt32toMultiply {
				dec.err = InvalidTypeError("Overflows int32")
				return 0
			}
			val = (val << 3) + (val << 1)
			if maxInt32-val < intv {
				dec.err = InvalidTypeError("Overflows int322")
				return 0
			}
			val += intv
		}
	} else {
		dec.err = InvalidTypeError("Overflows int32")
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
				dec.err = InvalidTypeError("Overflows uint32")
				return 0
			}
			val = (val << 3) + (val << 1)
			if maxUint32-val < uintv {
				dec.err = InvalidTypeError("Overflows int32")
				return 0
			}
			val += uintv
		}
	} else if ll > maxUint32Length {
		dec.err = InvalidTypeError("Overflows uint32")
		val = 0
	}
	return val
}
