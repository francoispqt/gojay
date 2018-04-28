package gojay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeStreamBorrow(t *testing.T) {
	// we override the pool chan
	streamDecPool = make(chan *StreamDecoder, 1)
	// add one decoder to the channel
	dec := Stream.NewDecoder(nil)
	streamDecPool <- dec
	// borrow one decoder to the channel
	nDec := Stream.BorrowDecoder(nil, 0)
	// make sure they are the same
	assert.Equal(t, dec, nDec, "decoder added to the pool and new decoder should be the same")
}

func TestDecodeStreamBorrow1(t *testing.T) {
	// we override the pool chan
	streamDecPool = make(chan *StreamDecoder, 1)
	// add one decoder to the channel
	dec := Stream.NewDecoder(nil)
	streamDecPool <- dec
	// reset streamDecPool
	streamDecPool = make(chan *StreamDecoder, 1)
	// borrow one decoder to the channel
	nDec := Stream.BorrowDecoder(nil, 0)
	// make sure they are the same
	assert.NotEqual(t, dec, nDec, "decoder added to the pool and new decoder should be the same")
}
func TestDecodeStreamBorrow2(t *testing.T) {
	// we override the pool chan
	streamDecPool = make(chan *StreamDecoder, 1)
	// add one decoder to the channel
	dec := Stream.NewDecoder(nil)
	dec.data = make([]byte, 128)
	streamDecPool <- dec
	// borrow one decoder to the channel
	nDec := Stream.BorrowDecoder(nil, 512)
	// make sure they are the same
	assert.Equal(t, dec, nDec, "decoder added to the pool and new decoder should be the same")
	assert.Equal(t, 512, len(nDec.data), "len of dec.data should be 512")
}
func TestDecodeStreamBorrow3(t *testing.T) {
	// we override the pool chan
	streamDecPool = make(chan *StreamDecoder, 16)
	// borrow one decoder to the channel
	nDec := Stream.BorrowDecoder(nil, 512)
	// make sure they are the same
	assert.Equal(t, 512, len(nDec.data), "len of dec.data should be 512")
}

func TestDecodeStreamDecodePooledDecoderError(t *testing.T) {
	// we override the pool chan
	dec := Stream.NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	var v = 0
	dec.Decode(&v)
	// make sure it fails if this is called
	assert.True(t, false, "should not be called as decoder should have panicked")
}

func TestDecodeStreamDecodePooledDecoderError1(t *testing.T) {
	// we override the pool chan
	dec := Stream.NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
	}()
	var v = testSliceStrings{}
	dec.DecodeArray(&v)
	// make sure they are the same
	assert.True(t, false, "should not be called as decoder should have panicked")
}

func TestDecodeStreamDecodePooledDecoderError2(t *testing.T) {
	// we override the pool chan
	dec := Stream.NewDecoder(nil)
	dec.Release()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err shouldnot be nil")
		assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
		assert.Equal(t, "Invalid usage of pooled decoder", err.(InvalidUsagePooledDecoderError).Error(), "err should be of type InvalidUsagePooledDecoderError")
	}()
	var v = TestObj{}
	dec.DecodeObject(&v)
	// make sure they are the same
	assert.True(t, false, "should not be called as decoder should have panicked")
}
