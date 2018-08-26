package gojay

// DecodeArray reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
//
// v must implement UnmarshalerJSONArray.
//
// See the documentation for Unmarshal for details about the conversion of JSON into a Go value.
func (dec *Decoder) DecodeArray(arr UnmarshalerJSONArray) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	_, err := dec.decodeArray(arr)
	return err
}
func (dec *Decoder) decodeArray(arr UnmarshalerJSONArray) (int, error) {
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
				err := arr.UnmarshalJSONArray(dec)
				if err != nil {
					return 0, err
				}
				n++
			}
			return 0, dec.raiseInvalidJSONErr(dec.cursor)
		case 'n':
			// is null
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return 0, err
			}
			dec.cursor++
			return dec.cursor, nil
		case '{', '"', 'f', 't', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// can't unmarshall to struct
			// we skip array and set Error
			dec.err = dec.makeInvalidUnmarshalErr(arr)
			err := dec.skipData()
			if err != nil {
				return 0, err
			}
			return dec.cursor, nil
		default:
			return 0, dec.raiseInvalidJSONErr(dec.cursor)
		}
	}
	return 0, dec.raiseInvalidJSONErr(dec.cursor)
}

// func (dec *Decoder) decodeArrayNull(factory func() UnmarshalerJSONArray) (int, error) {
// 	// not an array not an error, but do not know what to do
// 	// do not check syntax
// 	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
// 		switch dec.data[dec.cursor] {
// 		case ' ', '\n', '\t', '\r', ',':
// 			continue
// 		case '[':
// 			n := 0
// 			dec.cursor = dec.cursor + 1
// 			// array is open, char is not space start readings
// 			for dec.nextChar() != 0 {
// 				// closing array
// 				if dec.data[dec.cursor] == ']' {
// 					dec.cursor = dec.cursor + 1
// 					return dec.cursor, nil
// 				}
// 				// calling unmarshall function for each element of the slice
// 				err := arr.UnmarshalJSONArray(dec)
// 				if err != nil {
// 					return 0, err
// 				}
// 				n++
// 			}
// 			return 0, dec.raiseInvalidJSONErr(dec.cursor)
// 		case 'n':
// 			// is null
// 			dec.cursor++
// 			err := dec.assertNull()
// 			if err != nil {
// 				return 0, err
// 			}
// 			dec.cursor++
// 			return dec.cursor, nil
// 		case '{', '"', 'f', 't', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
// 			// can't unmarshall to struct
// 			// we skip array and set Error
// 			dec.err = dec.makeInvalidUnmarshalErr(arr)
// 			err := dec.skipData()
// 			if err != nil {
// 				return 0, err
// 			}
// 			return dec.cursor, nil
// 		default:
// 			return 0, dec.raiseInvalidJSONErr(dec.cursor)
// 		}
// 	}
// 	return 0, dec.raiseInvalidJSONErr(dec.cursor)
// }

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
			var isInEscapeSeq bool
			var isFirstQuote = true
			for ; j < dec.length || dec.read(); j++ {
				if dec.data[j] != '"' {
					continue
				}
				if dec.data[j-1] != '\\' || (!isInEscapeSeq && !isFirstQuote) {
					break
				} else {
					isInEscapeSeq = false
				}
				if isFirstQuote {
					isFirstQuote = false
				}
				// loop backward and count how many anti slash found
				// to see if string is effectively escaped
				ct := 0
				for i := j - 1; i > 0; i-- {
					if dec.data[i] != '\\' {
						break
					}
					ct++
				}
				// is pair number of slashes, quote is not escaped
				if ct&1 == 0 {
					break
				}
				isInEscapeSeq = true
			}
		default:
			continue
		}
	}
	return 0, dec.raiseInvalidJSONErr(dec.cursor)
}

// DecodeArrayFunc is a custom func type implementing UnarshaleArray.
// Use it to cast a func(*Decoder) to Unmarshal an object.
//
//	str := ""
//	dec := gojay.NewDecoder(io.Reader)
//	dec.DecodeArray(gojay.DecodeArrayFunc(func(dec *gojay.Decoder, k string) error {
//		return dec.AddString(&str)
//	}))
type DecodeArrayFunc func(*Decoder) error

// UnmarshalJSONArray implements UnarshalerArray.
func (f DecodeArrayFunc) UnmarshalJSONArray(dec *Decoder) error {
	return f(dec)
}

// IsNil implements UnarshalerArray.
func (f DecodeArrayFunc) IsNil() bool {
	return f == nil
}
