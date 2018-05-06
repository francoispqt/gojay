package gojay

import "io"

var encPool = make(chan *Encoder, 16)
var streamEncPool = make(chan *StreamEncoder, 16)
var bufPool = make(chan []byte, 16)

func init() {
initStreamEncPool:
	for {
		select {
		case streamEncPool <- Stream.NewEncoder(nil):
		default:
			break initStreamEncPool
		}
	}
initEncPool:
	for {
		select {
		case encPool <- NewEncoder(nil):
		default:
			break initEncPool
		}
	}
	for {
		select {
		case bufPool <- make([]byte, 0, 512):
		default:
			return
		}
	}
}

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
		enc.err = nil
		return enc
	default:
		return &Encoder{w: w, buf: borrowBuf()}
	}
}

// Release sends back a Encoder to the pool.
func (enc *Encoder) Release() {
	enc.buf = enc.buf[:0]
	enc.isPooled = 1
	select {
	case encPool <- enc:
	default:
	}
}

func borrowBuf() []byte {
	select {
	case b := <-bufPool:
		return b
	default:
		return make([]byte, 0, 512)
	}
}
