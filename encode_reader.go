package gojay

import (
	"encoding/base64"
	"io"
)

type writerEncoder Encoder // we can't simply implement io.Writer on Encoder because Write on Encoder already present

func (enc *writerEncoder) Write(p []byte) (int, error) {
	(*Encoder)(enc).writeBytes(p)
	return len(p), nil
}

type escapingWriterEncoder Encoder

func (enc *escapingWriterEncoder) Write(p []byte) (int, error) {
	(*Encoder)(enc).writeBytesEscape(p)
	return len(p), nil
}

func (enc *Encoder) writeBase64(r io.Reader, encoding *base64.Encoding) {
	b64enc := base64.NewEncoder(encoding, (*writerEncoder)(enc))
	_, err := io.Copy(b64enc, r)
	if err != nil {
		enc.err = err
		return
	}
	if err := b64enc.Close(); err != nil {
		enc.err = err
		return
	}
}

// AddReaderToBase64 adds a data to be base64-encoded (read from provided reader)
// must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddReaderToBase64(r io.Reader, encoding *base64.Encoding) {
	enc.ReaderToBase64(r, encoding)
}

// ReaderToBase64 adds a data to be base64-encoded (read from provided reader)
// must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) ReaderToBase64(r io.Reader, encoding *base64.Encoding) {
	prevRune := enc.getPreviousRune()
	if prevRune != '[' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeBase64(r, encoding)
	enc.writeByte('"')
}

// AddReaderToBase64 adds a data to be base64-encoded (read from provided reader)
// must be used inside an object as it will encode a key
func (enc *Encoder) AddReaderToBase64Key(key string, r io.Reader, encoding *base64.Encoding) {
	enc.ReaderToBase64Key(key, r, encoding)
}

// ReaderToBase64Key adds a data to be base64-encoded (read from provided reader)
// must be used inside an object as it will encode a key
func (enc *Encoder) ReaderToBase64Key(key string, r io.Reader, encoding *base64.Encoding) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	enc.grow(2 + len(key))
	prevRune := enc.getPreviousRune()
	if prevRune != '{' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.writeByte('"')
	enc.writeBase64(r, encoding)
	enc.writeByte('"')
}

// AddReaderToEscaped adds a string to be encoded (read from provided reader)
// must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) AddReaderToEscaped(r io.Reader) {
	enc.ReaderToEscaped(r)
}

// ReaderToEscaped adds a string to be encoded (read from provided reader)
// must be used inside a slice or array encoding (does not encode a key)
func (enc *Encoder) ReaderToEscaped(r io.Reader) {
	prevRune := enc.getPreviousRune()
	if prevRune != '[' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	_, err := io.Copy((*escapingWriterEncoder)(enc), r)
	if err != nil {
		enc.err = err
	}
	enc.writeByte('"')
}

// AddReaderToEscapedKey adds a string to be encoded (read from provided reader)
// must be used inside an object as it will encode a key
func (enc *Encoder) AddReaderToEscapedKey(key string, r io.Reader) {
	enc.ReaderToEscapedKey(key, r)
}

// ReaderToEscapedKey adds a string to be encoded (read from provided reader)
// must be used inside an object as it will encode a key
func (enc *Encoder) ReaderToEscapedKey(key string, r io.Reader) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	enc.grow(2 + len(key))
	prevRune := enc.getPreviousRune()
	if prevRune != '{' {
		enc.writeTwoBytes(',', '"')
	} else {
		enc.writeByte('"')
	}
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	enc.writeByte('"')
	_, enc.err = io.Copy((*escapingWriterEncoder)(enc), r)
	enc.writeByte('"')
}
