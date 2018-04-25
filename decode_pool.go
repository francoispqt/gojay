package gojay

import "io"

var decPool = make(chan *Decoder, 16)

// NewDecoder returns a new decoder or borrows one from the pool
// it takes an io.Reader implementation as data input
func NewDecoder(r io.Reader) *Decoder {
	return newDecoder(r, 512)
}

func newDecoder(r io.Reader, bufSize int) *Decoder {
	select {
	case dec := <-decPool:
		dec.called = 0
		dec.keysDone = 0
		dec.cursor = 0
		dec.err = nil
		dec.r = r
		dec.length = 0
		if bufSize > 0 {
			dec.data = make([]byte, bufSize)
		}
		return dec
	default:
		dec := &Decoder{
			called:   0,
			cursor:   0,
			keysDone: 0,
			err:      nil,
			r:        r,
		}
		if bufSize > 0 {
			dec.data = make([]byte, bufSize)
			dec.length = 0
		}
		return dec
	}
}

func (dec *Decoder) addToPool() {
	select {
	case decPool <- dec:
	default:
	}
}
