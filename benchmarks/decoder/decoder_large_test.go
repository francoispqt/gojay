package benchmarks

import (
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	"github.com/stretchr/testify/assert"
)

func TestGoJayDecodeObjLarge(t *testing.T) {
	result := benchmarks.LargePayload{}
	err := gojay.UnmarshalJSONObject(benchmarks.LargeFixture, &result)
	assert.Nil(t, err, "err should be nil")
	assert.Len(t, result.Users, 32, "Len of users should be 32")
	for _, u := range result.Users {
		assert.True(t, len(u.Username) > 0, "User should have username")
	}
	assert.Len(t, result.Topics.Topics, 30, "Len of topics should be 30")
	for _, top := range result.Topics.Topics {
		assert.True(t, top.Id > 0, "Topic should have Id")
	}
}
