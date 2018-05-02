package client

import (
	"github.com/francoispqt/gojay/examples/websocket/comm"
	"golang.org/x/net/websocket"
)

type client struct {
	comm.SenderReceiver
	id int
}

func NewClient(id int) *client {
	c := new(client)
	c.id = id
	return c
}

func (c *client) Dial(url, origin string) error {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	c.Conn = conn
	c.Init(10)
	return nil
}
