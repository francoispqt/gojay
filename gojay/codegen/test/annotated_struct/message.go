package annotated_struct

type Paylod []byte

type Message struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Price       float64       `json:"price"`
	Ints        []int         `json:"ints"`
	Floats      []*float32    `json:"floats"`
	SubMessageX *SubMessage   `json:"subMessageX"`
	MessagesX   []*SubMessage `json:"messagesX"`
	SubMessageY SubMessage
	MessagesY   []SubMessage
	IsTrue      *bool  `json:"enabled"`
	Payload     Paylod `json:"data"`
	Ignore      string `json:"-"`
}
