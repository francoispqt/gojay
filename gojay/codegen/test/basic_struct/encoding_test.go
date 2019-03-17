package basic_struct

import (
	"bytes"
	"database/sql"
	"testing"

	"github.com/francoispqt/gojay"
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
    "Description": "abcd",
	"StartTime": "0001-01-01T00:00:00Z"
  },
  "MessagesX": [
    {
      "Id": 2102,
      "Description": "abce", 
	  "StartTime": "0001-01-01T00:00:00Z"
	}
  ],
  "SubMessageY": {
    "Id": 3102,
    "Description": "abcf",
    "StartTime": "0001-01-01T00:00:00Z"
  },
  "MessagesY": [
    {
      "Id": 5102,
      "Description": "abcg", 
	  "StartTime": "0001-01-01T00:00:00Z"
	},
    {
      "Id": 5106,
      "Description": "abcgg",
	  "StartTime": "0001-01-01T00:00:00Z"
    }
  ],
  "IsTrue": true,
  "Payload": "123",
  "SQLNullString": "test"
}`

func TestMessage_Unmarshal(t *testing.T) {
	var err error
	var data = []byte(jsonData)
	message := &Message{}
	err = gojay.UnmarshalJSONObject(data, message)
	require.Nil(t, err)
	require.Equal(t, msg, message)
}

func TestMessage_Marshal(t *testing.T) {
	var writer = new(bytes.Buffer)

	encoder := gojay.NewEncoder(writer)
	var err = encoder.Encode(msg)

	require.Nil(t, err)
	var JSON = writer.String()

	require.JSONEq(t, jsonData, JSON)
}
