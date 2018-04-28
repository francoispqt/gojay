package gojay

import "io"

var streamDecPool = make(chan *StreamDecoder, 16)

// NewDecoder returns a new decoder.
// It takes an io.Reader implementation as data input.
// It initiates the done channel returned by Done().
func (s stream) NewDecoder(r io.Reader) *StreamDecoder {
	dec := NewDecoder(r)
	streamDec := &StreamDecoder{
		Decoder: dec,
		done:    make(chan struct{}, 1),
	}
	return streamDec
}

// BorrowDecoder borrows a StreamDecoder a decoder from the pool.
// It takes an io.Reader implementation as data input.
// It initiates the done channel returned by Done().
func (s stream) BorrowDecoder(r io.Reader, bufSize int) *StreamDecoder {
	select {
	case streamDec := <-streamDecPool:
		streamDec.called = 0
		streamDec.keysDone = 0
		streamDec.cursor = 0
		streamDec.err = nil
		streamDec.r = r
		streamDec.length = 0
		streamDec.isPooled = 0
		streamDec.done = make(chan struct{}, 1)
		if bufSize > 0 {
			streamDec.data = make([]byte, bufSize)
		}
		return streamDec
	default:
		dec := NewDecoder(r)
		if bufSize > 0 {
			dec.data = make([]byte, bufSize)
			dec.length = 0
		}
		streamDec := &StreamDecoder{
			Decoder: dec,
			done:    make(chan struct{}, 1),
		}
		return streamDec
	}
}
