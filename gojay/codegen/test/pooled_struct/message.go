package pooled_struct

type Message struct {
	Id          int
	Name        string
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
