package benchmarks

import (
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	"github.com/stretchr/testify/assert"
)

func TestGoJayDecodeObjSmall(t *testing.T) {
	result := benchmarks.SmallPayload{}
	err := gojay.Unmarshal(benchmarks.SmallFixture, &result)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, result.St, 1, "result.St should be 1")
	assert.Equal(t, result.Sid, 486, "result.Sid should be 486")
	assert.Equal(t, result.Tt, "active", "result.Sid should be 'active'")
	assert.Equal(t, result.Gr, 0, "result.Gr should be 0")
	assert.Equal(
		t,
		result.Uuid,
		"de305d54-75b4-431b-adb2-eb6b9e546014",
		"result.Gr should be 'de305d54-75b4-431b-adb2-eb6b9e546014'",
	)
	assert.Equal(t, result.Ip, "127.0.0.1", "result.Ip should be '127.0.0.1'")
	assert.Equal(t, result.Ua, "user_agent", "result.Ua should be 'user_agent'")
	assert.Equal(t, result.Tz, -6, "result.Tz should be 6")
	assert.Equal(t, result.V, 1, "result.V should be 1")
}
