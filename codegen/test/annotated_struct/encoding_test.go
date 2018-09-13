package annotated_struct

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/francoispqt/gojay"
	"github.com/viant/assertly"
	"log"
	"bytes"
)

func TestMessage_Unmarshal(t *testing.T) {



	input := `{
  "id": 1022,
  "name": "name acc",
  "price": 13.3,
  "ints": [
    1,
    2,
    5
  ],
  "floats": [
    2.3,
    4.6,
    7.4
  ],
  "subMessageX": {
    "id": 102,
    "description": "abcd"
  },
  "messagesX": [
    {
      "id": 2102,
      "description": "abce"
    }
  ],
  "SubMessageY": {
    "id": 3102,
    "description": "abcf"
  },
  "MessagesY": [
    {
      "id": 5102,
      "description": "abcg"
    },
    {
      "id": 5106,
      "description": "abcgg"
    }
  ],
  "enabled": true,
  "data": "123"
}`


	expacted := `{
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
  "Payload": "\"123\""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	if !assert.Nil(t, err) {
		log.Fatal(err)
	}
	assertly.AssertValues(t, expacted, message)
}




func TestMessage_Marshal(t *testing.T) {

	input := `{
  "id": 1022,
  "name": "name acc",
  "price": 13.3,
  "ints": [
    1,
    2,
    5
  ],
  "floats": [
    2.3,
    4.6,
    7.4
  ],
  "subMessageX": {
    "id": 102,
    "description": "abcd"
  },
  "messagesX": [
    {
      "id": 2102,
      "description": "abce"
    }
  ],
  "SubMessageY": {
    "id": 3102,
    "description": "abcf"
  },
  "MessagesY": [
    {
      "id": 5102,
      "description": "abcg"
    },
    {
      "id": 5106,
      "description": "abcgg"
    }
  ],
  "enabled": true,
  "data": "123"
}`


	expacted := `{
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
  "Payload": "\"123\""
}`

	var err error
	var data = []byte(input)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	if !assert.Nil(t, err) {
		log.Fatal(err)
	}
	assertly.AssertValues(t, expacted, message)
	var writer = new(bytes.Buffer)
	encoder :=  gojay.NewEncoder(writer)
	err = encoder.Encode(message)
	assert.Nil(t, err)
	var JSON = writer.String()
	assertly.AssertValues(t, input,  JSON)

}