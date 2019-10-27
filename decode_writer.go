package gojay

import (
	"encoding/base64"
	"io"
)

type stringDecodeReader Decoder

func (dec *stringDecodeReader) copyAndShift(target []byte, start int) int {
	n := copy(target, dec.data[start:dec.cursor])
	// shift buffer left to reuse it between reads
	copy(dec.data[start:], dec.data[start+n:])
	dec.length -= n
	dec.cursor = start
	return n
}

func (dec *stringDecodeReader) Read(b []byte) (int, error) {
	start := dec.cursor
	for dec.cursor-start < len(b) && (dec.cursor < dec.length || (*Decoder)(dec).read()) {
		switch dec.data[dec.cursor] {
		// string end
		case '"':
			// making copy before exit because parseEscapedString may change
			// previous (before cursor) bytes of data buffer
			n := dec.copyAndShift(b, start)
			dec.cursor++
			return n, io.EOF
		// escape sequence
		case '\\':
			dec.cursor++
			err := (*Decoder)(dec).parseEscapedString()
			if err != nil {
				return 0, err
			}
		default:
			dec.cursor++
		}
	}

	if dec.cursor-start < len(b) {
		// input buffer not filled and exited before EOF
		// that means json is invalid
		return 0, (*Decoder)(dec).raiseInvalidJSONErr(dec.cursor)
	}

	return dec.copyAndShift(b, start), nil
}

func (dec *Decoder) decodeStringStream() (*stringDecodeReader, error) {
	for ; dec.cursor < dec.length || dec.read(); dec.cursor++ {
		switch dec.data[dec.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			// is string
			continue
		case '"':
			dec.cursor++
			return (*stringDecodeReader)(dec), nil
		// is nil
		case 'n':
			dec.cursor++
			return nil, dec.assertNull()
		default:
			dec.err = dec.makeInvalidUnmarshalErr((*stringDecodeReader)(nil))
			return nil, dec.skipData()
		}
	}
	return nil, nil
}

// AddWriterFromEscaped decodes the JSON value within an object or an array to a provided writer.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) AddWriterFromEscaped(w io.Writer) error {
	return dec.WriterFromEscaped(w)
}

// WriterFromEscaped decodes the JSON value within an object or an array to a provided writer.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) WriterFromEscaped(w io.Writer) error {
	reader, err := dec.decodeStringStream()
	if err != nil {
		return err
	}
	dec.called |= 1
	if reader == nil {
		return nil
	}
	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}
	return nil
}

// AddWriterFromBase64 decodes the JSON value (base64-encoded data) within an object or an array to a provided writer.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) AddWriterFromBase64(w io.Writer, encoding *base64.Encoding) error {
	return dec.WriterFromBase64(w, encoding)
}

// WriterFromEscaped decodes the JSON value (base64-encoded data) within an object or an array to a provided writer.
// If next key is not a JSON string nor null, InvalidUnmarshalError will be returned.
func (dec *Decoder) WriterFromBase64(w io.Writer, encoding *base64.Encoding) error {
	reader, err := dec.decodeStringStream()
	if err != nil {
		return err
	}
	dec.called |= 1
	if reader == nil {
		return nil
	}
	_, err = io.Copy(w, base64.NewDecoder(encoding, reader))
	if err != nil {
		return err
	}
	return nil
}
