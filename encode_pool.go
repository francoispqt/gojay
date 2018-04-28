package gojay

var encObjPool = make(chan *Encoder, 16)

// NewEncoder returns a new encoder or borrows one from the pool
func NewEncoder() *Encoder {
	return &Encoder{}
}
func newEncoder() *Encoder {
	return &Encoder{}
}

// BorrowEncoder borrows an Encoder from the pool.
func BorrowEncoder() *Encoder {
	select {
	case enc := <-encObjPool:
		enc.isPooled = 0
		return enc
	default:
		return &Encoder{}
	}
}

// Release sends back a Encoder to the pool.
func (enc *Encoder) Release() {
	enc.buf = nil
	select {
	case encObjPool <- enc:
		enc.isPooled = 1
	default:
	}
}
