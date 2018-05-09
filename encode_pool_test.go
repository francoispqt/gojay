package gojay

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoderNewFromPool(t *testing.T) {
	// reset pool
	encPool = sync.Pool{
		New: func() interface{} {
			return NewEncoder(nil)
		},
	}

	// get new Encoder
	enc := encPool.New().(*Encoder)
	// add to pool
	enc.Release()
	// borrow encoder
	nEnc := BorrowEncoder(nil)
	// make sure it's the same
	assert.Equal(t, enc, nEnc, "enc and nEnc from pool should be the same")
}
