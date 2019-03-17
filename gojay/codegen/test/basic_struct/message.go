package basic_struct

import "database/sql"

type Message struct {
	Id            int
	Name          string
	Price         float64
	Ints          []int
	Floats        []float32
	SubMessageX   *SubMessage
	MessagesX     []*SubMessage
	SubMessageY   SubMessage
	MessagesY     []SubMessage
	IsTrue        *bool
	Payload       []byte
	SQLNullString *sql.NullString
}
