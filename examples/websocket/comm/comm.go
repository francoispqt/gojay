package comm

import (
	"errors"
	"log"

	"github.com/francoispqt/gojay"
	"golang.org/x/net/websocket"
)

// A basic message for our WebSocket app

type Message struct {
	Message  string
	UserName string
}

func (m *Message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "message":
		return dec.AddString(&m.Message)
	case "userName":
		return dec.AddString(&m.UserName)
	}
	return nil
}
func (m *Message) NKeys() int {
	return 2
}

func (m *Message) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("message", m.Message)
	enc.AddStringKey("userName", m.UserName)
}
func (u *Message) IsNil() bool {
	return u == nil
}

// Here are defined our communication types
type Sender chan gojay.MarshalerJSONObject

func (s Sender) MarshalStream(enc *gojay.StreamEncoder) {
	select {
	case <-enc.Done():
		return
	case m := <-s:
		enc.AddObject(m)
	}
}

type Receiver chan *Message

func (s Receiver) UnmarshalStream(dec *gojay.StreamDecoder) error {
	m := &Message{}
	if err := dec.AddObject(m); err != nil {
		return err
	}
	s <- m
	return nil
}

type SenderReceiver struct {
	Send    Sender
	Receive Receiver
	Dec     *gojay.StreamDecoder
	Enc     *gojay.StreamEncoder
	Conn    *websocket.Conn
}

func (sc *SenderReceiver) SetReceiver() {
	sc.Receive = Receiver(make(chan *Message))
	sc.Dec = gojay.Stream.BorrowDecoder(sc.Conn)
	go sc.Dec.DecodeStream(sc.Receive)
}

func (sc *SenderReceiver) SetSender(nCons int) {
	sc.Send = Sender(make(chan gojay.MarshalerJSONObject))
	sc.Enc = gojay.Stream.BorrowEncoder(sc.Conn).NConsumer(nCons).LineDelimited()
	go sc.Enc.EncodeStream(sc.Send)
}

func (sc *SenderReceiver) SendMessage(m gojay.MarshalerJSONObject) error {
	select {
	case <-sc.Enc.Done():
		return errors.New("sender closed")
	case sc.Send <- m:
		log.Print("message sent by client: ", m)
		return nil
	}
}

func (c *SenderReceiver) OnMessage(f func(*Message)) error {
	for {
		select {
		case <-c.Dec.Done():
			return errors.New("receiver closed")
		case m := <-c.Receive:
			f(m)
		}
	}
}

func (sc *SenderReceiver) Init(sender int) *SenderReceiver {
	sc.SetSender(sender)
	sc.SetReceiver()
	return sc
}
