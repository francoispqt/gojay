package benchmarks

import (
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	"github.com/stretchr/testify/assert"
)

func TestGoJayDecodeObjMedium(t *testing.T) {
	result := benchmarks.MediumPayload{}
	err := gojay.Unmarshal(benchmarks.MediumFixture, &result)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, "Leonid Bugaev", result.Person.Name.FullName, "result.Person.Name.FullName should be Leonid Bugaev")
	assert.Equal(t, 95, result.Person.Github.Followers, "result.Person.Github.Followers should be 95")
	assert.Len(t, result.Person.Gravatar.Avatars, 1, "result.Person.Gravatar.Avatars should have 1 item")
}
