package gojay

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncoderBuilderError(t *testing.T) {
	enc := NewEncoder()
	defer func() {
		err := recover()
		assert.NotNil(t, err, "err is not nil as we pass an invalid number to grow")
	}()
	enc.grow(-1)
	assert.True(t, false, "should not be called")
}