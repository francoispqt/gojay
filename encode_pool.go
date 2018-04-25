package gojay

var encObjPool = make(chan *Encoder, 16)

// NewEncoder returns a new encoder or borrows one from the pool
func NewEncoder() *Encoder {
	select {
	case enc := <-encObjPool:
		return enc
	default:
		return &Encoder{}
	}
}

func (enc *Encoder) addToPool() {
	enc.buf = nil
	select {
	case encObjPool <- enc:
	default:
	}
}
