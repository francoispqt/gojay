package gojay

import (
	"fmt"
)

// DecodeArray reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
//
// v must implement UnmarshalerArray.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeArray(arr UnmarshalerArray) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	_, err := dec.decodeArray(arr)
	return err
}
func (dec *Decoder) decodeArray(arr UnmarshalerArray) (int, error) {
	// not an array not an error, but do not know what to do
	// do not check syntax
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case '[':
			n := 0
			dec.cursor = dec.cursor + 1
			// array is open, char is not space start readings
			for dec.nextChar() != 0 {
				// closing array
				if dec.data[dec.cursor] == ']' {
					dec.cursor = dec.cursor + 1
					return dec.cursor, nil
				}
				// calling unmarshall function for each element of the slice
				err := arr.UnmarshalArray(dec)
				if err != nil {
					return 0, err
				}
				n++
			}
			return dec.cursor, nil
		case 'n':
			// is null
			dec.cursor = dec.cursor + 4
			return dec.cursor, nil
		case '{', '"', 'f', 't', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// can't unmarshall to struct
			// we skip array and set Error
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to array, wrong char '%s' found at pos %d",
					string(dec.data[dec.cursor]),
					dec.cursor,
				),
			)
			err := dec.skipData()
			if err != nil {
				return 0, err
			}
			return dec.cursor, nil
		default:
			return 0, InvalidJSONError("Invalid JSON")
		}
	}
	return 0, InvalidJSONError("Invalid JSON")
}

func (dec *Decoder) skipArray() (int, error) {
	var arraysOpen = 1
	var arraysClosed = 0
	// var stringOpen byte = 0
	for j := dec.cursor; j < dec.length || dec.read(); j++ {
		switch dec.data[j] {
		case ']':
			arraysClosed++
			// everything is closed return
			if arraysOpen == arraysClosed {
				// add char to object data
				return j + 1, nil
			}
		case '[':
			arraysOpen++
		case '"':
			j++
			for ; j < dec.length; j++ {
				if dec.data[j] != '"' {
					continue
				}
				if dec.data[j-1] != '\\' {
					break
				}
				// loop backward and count how many anti slash found
				// to see if string is effectively escaped
				ct := 1
				for i := j - 2; i > 0; i-- {
					if dec.data[i] != '\\' {
						break
					}
					ct++
				}
				// is even number of slashes, quote is not escaped
				if ct&1 == 0 {
					break
				}
			}
		default:
			continue
		}
	}
	return 0, InvalidJSONError("Invalid JSON")
}
