package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderNewFromPool(t *testing.T) {
	// reset pool
	encObjPool = make(chan *Encoder, 16)
	// get new Encoder
	enc := NewEncoder()
	// add to pool
	enc.Release()
	// borrow encoder
	nEnc := BorrowEncoder()
	// make sure it's the same
	assert.Equal(t, enc, nEnc, "enc and nEnc from pool should be the same")
}
