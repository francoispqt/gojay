package gojay

import "fmt"

// DecodeBool reads the next JSON-encoded value from its input and stores it in the boolean pointed to by v.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeBool(v *bool) error {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			continue
		case 't':
			dec.cursor = dec.cursor + 4
			*v = true
			return nil
		case 'f':
			dec.cursor = dec.cursor + 5
			*v = false
			return nil
		case 'n':
			dec.cursor = dec.cursor + 4
			*v = false
			return nil
		default:
			dec.err = InvalidTypeError(
				fmt.Sprintf(
					"Cannot unmarshall to bool, wrong char '%s' found at pos %d",
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
