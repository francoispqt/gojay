package unknown_struct

import (
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/require"
)

var msg = &Message{
	Id:   1022,
	Name: "name acc",
}

var jsonData = `{
  "Id": 1022,
  "Name": "name acc",
  "Price": 13.3
}`

func TestMessage_Unmarshal(t *testing.T) {
	var err error
	var data = []byte(jsonData)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	require.EqualError(t, err, gojay.MakeUnknownFieldErr(message, "Price").Error())
}
