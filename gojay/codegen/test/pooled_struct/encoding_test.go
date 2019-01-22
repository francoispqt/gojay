package pooled_struct

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
  "SubMessageX": {
    "Id": 102,
    "Description": "abcd"
  },
  "MessagesX": [
    {
      "Id": 2102,
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Id": 3102,
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Id": 5102,
      "Description": "abcg"
    },
    {
      "Id": 5106,
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var data = []byte(input)
	message := MessagePool.Get().(*Message)
	err := gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)
	message.Reset()
	MessagePool.Put(message)

}

func TestMessage_Marshal(t *testing.T) {

	input := `{
  "Id": 1022,
  "Name": "name acc",
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
  "SubMessageX": {
    "Id": 102,
    "Description": "abcd"
  },
  "MessagesX": [
    {
      "Id": 2102,
      "Description": "abce"
    }
  ],
  "SubMessageY": {
    "Id": 3102,
    "Description": "abcf"
  },
  "MessagesY": [
    {
      "Id": 5102,
      "Description": "abcg"
    },
    {
      "Id": 5106,
      "Description": "abcgg"
    }
  ],
  "IsTrue": true,
  "Payload": ""
}`

	var data = []byte(input)
	message := MessagePool.Get().(*Message)
	err := gojay.UnmarshalJSONObject(data, message)
	assert.Nil(t, err)
	defer func() {
		message.Reset()
		MessagePool.Put(message)

	}()
	var writer = new(bytes.Buffer)
	encoder := gojay.NewEncoder(writer)
	err = encoder.Encode(message)
	assert.Nil(t, err)
	var JSON = writer.String()
	assertly.AssertValues(t, input, JSON)
}
