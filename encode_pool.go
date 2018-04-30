package gojay

import "io"

var encPool = make(chan *Encoder, 16)
var streamEncPool = make(chan *StreamEncoder, 16)

// NewEncoder returns a new encoder or borrows one from the pool
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}
func newEncoder() *Encoder {
	return &Encoder{}
}

// BorrowEncoder borrows an Encoder from the pool.
func BorrowEncoder(w io.Writer) *Encoder {
	select {
	case enc := <-encPool:
		enc.isPooled = 0
		enc.w = w
		enc.buf = make([]byte, 0)
		return enc
	default:
		return &Encoder{w: w}
	}
}

// Release sends back a Encoder to the pool.
func (enc *Encoder) Release() {
	enc.buf = nil
	select {
	case encPool <- enc:
		enc.isPooled = 1
	default:
	}
}
