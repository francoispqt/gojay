package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeStreamBorrow1(t *testing.T) {
	// we override the pool chan
	streamEncPool = make(chan *StreamEncoder, 1)
	// add one decoder to the channel
	enc := Stream.NewEncoder(nil)
	streamEncPool <- enc
	// reset streamEncPool
	streamEncPool = make(chan *StreamEncoder, 1)
	// borrow one decoder to the channel
	nEnc := Stream.BorrowEncoder(nil)
	// make sure they are the same
	assert.NotEqual(t, enc, nEnc, "encoder added to the pool and new decoder should be the same")
}
