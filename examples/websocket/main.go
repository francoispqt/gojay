// package main simulates a conversation between
// a given set of websocket clients and a server.
//
// It uses gojay's streaming feature to abstract JSON communication
// and only having to handle go values.
package main

import (
	"log"
	"strconv"

	"github.com/francoispqt/gojay/examples/websocket/client"
	"github.com/francoispqt/gojay/examples/websocket/comm"
	"github.com/francoispqt/gojay/examples/websocket/server"
)

func createServer(done chan error) {
	// create our server, with a done signal
	s := server.NewServer()
	// set our connection handler
	s.OnConnection(func(c *server.Client) {
		// send welcome message to initiate the conversation
		c.SendMessage(&comm.Message{
			UserName: "server",
			Message:  "Welcome !",
		})
		// start handling messages
		c.OnMessage(func(m *comm.Message) {
			log.Print("message received from client: ", m)
			s.BroadCastRandom(c, m)
		})
	})
	go s.Listen(":8070", done)
}

func createClient(url, origin string, i int) {
	// create our client
	c := client.NewClient(i)
	// Dial connection to the WS server
	err := c.Dial(url, origin)
	if err != nil {
		panic(err)
	}
	str := strconv.Itoa(i)
	// Init client's sender and receiver
	// Set the OnMessage handler
	c.Init(10).OnMessage(func(m *comm.Message) {
		log.Print("client "+str+" received from "+m.UserName+" message: ", m)
		c.SendMessage(&comm.Message{
			UserName: str,
			Message:  "Responding to: " + m.UserName + " | old message: " + m.Message,
		})
	})
}

// Our main function
func main() {
	done := make(chan error)
	createServer(done)
	// add our clients connection
	for i := 0; i < 100; i++ {
		i := i
		go createClient("ws://localhost:8070/ws", "http://localhost/", i)
	}
	// handle server's termination
	log.Fatal(<-done)
}
