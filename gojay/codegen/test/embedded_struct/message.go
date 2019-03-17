package embedded_struct

type BaseId struct {
	Id   int
	Name string
}

type Message struct {
	*BaseId
	SubMessage
	Price       float64
	Ints        []int
	Floats      []float64
	SubMessageX *SubMessage
	MessagesX   []*SubMessage
	SubMessageY SubMessage
	MessagesY   []SubMessage
	IsTrue      *bool
	Payload     []byte
}
