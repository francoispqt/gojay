package annotated_struct

import "database/sql"

type Payload []byte

type Message struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Price         float64       `json:"price"`
	Ints          []int         `json:"ints"`
	Floats        []float32     `json:"floats"`
	SubMessageX   *SubMessage   `json:"subMessageX"`
	MessagesX     []*SubMessage `json:"messagesX"`
	SubMessageY   SubMessage
	MessagesY     []SubMessage
	IsTrue        *bool           `json:"enabled"`
	Payload       Payload         `json:"data"`
	Ignore        string          `json:"-"`
	SQLNullString *sql.NullString `json:"sqlNullString"`
}
