package annotated_struct

import (
	"bytes"
	"database/sql"
	"log"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var isTrue = true
var msg = &Message{
	Id:     1022,
	Name:   "name acc",
	Price:  13.3,
	Ints:   []int{1, 2, 5},
	Floats: []float32{2.3, 4.6, 7.4},
	SubMessageX: &SubMessage{
		Id:          102,
		Description: "abcd",
	},
	MessagesX: []*SubMessage{
		&SubMessage{
			Id:          2102,
			Description: "abce",
		},
	},
	SubMessageY: SubMessage{
		Id:          3102,
		Description: "abcf",
	},
	MessagesY: []SubMessage{
		SubMessage{
			Id:          5102,
			Description: "abcg",
		},
		SubMessage{
			Id:          5106,
			Description: "abcgg",
		},
	},
	IsTrue:  &isTrue,
	Payload: []byte(`"123"`),
	SQLNullString: &sql.NullString{
		String: "test",
		Valid:  true,
	},
}

var jsonData = `{
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
    "description": "abcd",
	"startDate": "0001-01-01 00:00:00"
  },
  "messagesX": [
    {
      "id": 2102,
      "description": "abce",
	  "startDate": "0001-01-01 00:00:00"
    }
  ],
  "SubMessageY": {
    "id": 3102,
    "description": "abcf",
	"startDate": "0001-01-01 00:00:00"
  },
  "MessagesY": [
    {
      "id": 5102,
      "description": "abcg",
	  "startDate": "0001-01-01 00:00:00"
	},
    {
      "id": 5106,
      "description": "abcgg",
	  "startDate": "0001-01-01 00:00:00"
    }
  ],
  "enabled": true,
  "data": "123",
  "sqlNullString": "test"
}`

func TestMessage_Unmarshal(t *testing.T) {

	var err error
	var data = []byte(jsonData)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	if !assert.Nil(t, err) {
		log.Fatal(err)
	}
	require.Equal(
		t,
		msg,
		message,
	)
}

func TestMessage_Marshal(t *testing.T) {
	var err error
	var writer = new(bytes.Buffer)
	encoder := gojay.NewEncoder(writer)
	err = encoder.Encode(msg)
	assert.Nil(t, err)
	var JSON = writer.String()
	require.JSONEq(t, jsonData, JSON)

}
