package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderNewFromPool(t *testing.T) {
	// reset pool
	encPool = make(chan *Encoder, 16)
	// get new Encoder
	enc := NewEncoder(nil)
	// add to pool
	enc.Release()
	// borrow encoder
	nEnc := BorrowEncoder(nil)
	// make sure it's the same
	assert.Equal(t, enc, nEnc, "enc and nEnc from pool should be the same")
}
