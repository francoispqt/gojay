package gojay

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeStreamBorrow1(t *testing.T) {
	// we override the pool chan
	streamEncPool = sync.Pool{
		New: func() interface{} {
			return Stream.NewEncoder(nil)
		},
	}
	// add one decoder to the channel
	enc := Stream.NewEncoder(nil)
	streamEncPool.Put(enc)
	// reset streamEncPool
	streamEncPool = sync.Pool{
		New: func() interface{} {
			return Stream.NewEncoder(nil)
		},
	}
	// borrow one decoder to the channel
	nEnc := Stream.BorrowEncoder(nil)
	// make sure they are the same
	assert.NotEqual(t, enc, nEnc, "encoder added to the pool and new decoder should be the same")
}
