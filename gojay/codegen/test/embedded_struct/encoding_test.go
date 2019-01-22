package embedded_struct

import (
	"bytes"
	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/assert"
	"github.com/viant/assertly"
	"testing"
)

func TestMessage_Unmarshal(t *testing.T) {

	input := `{
  "Id": 1022,
  "Name": "name acc",
  "Description": "abcd",
  "Price": 13.3,
  "Ints": [
    1,
    2,
    5
  ],
  "Floats": [
    2.3,
    4.6,
    7.4
  ],
  "MessagesX": [
    {
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Description": "abcg"
    },
    {
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)
	assertly.AssertValues(t, input, message)
}

func TestMessage_Marshal(t *testing.T) {

	input := `{
  "Id": 1022,
  "Name": "name acc",
  "Description": "abcd",
  "Price": 13.3,
  "Ints": [
    1,
    2,
    5
  ],
  "Floats": [
    2.3,
    4.6,
    7.4
  ],
  "MessagesX": [
    {
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Description": "abcg"
    },
    {
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)
	assertly.AssertValues(t, input, message)
	var writer = new(bytes.Buffer)

	encoder := gojay.NewEncoder(writer)
	err = encoder.Encode(message)
	assert.Nil(t, err)
	var JSON = writer.String()
	assertly.AssertValues(t, input, JSON)
}
