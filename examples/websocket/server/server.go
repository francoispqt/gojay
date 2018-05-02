package server

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/francoispqt/gojay/examples/websocket/comm"
	"golang.org/x/net/websocket"
)

type server struct {
	clients []*Client
	mux     *sync.RWMutex
	handle  func(c *Client)
}

type Client struct {
	comm.SenderReceiver
	server *server
}

func NewClient(s *server, conn *websocket.Conn) *Client {
	sC := new(Client)
	sC.Conn = conn
	sC.server = s
	return sC
}

func NewServer() *server {
	s := new(server)
	s.mux = new(sync.RWMutex)
	s.clients = make([]*Client, 0, 100)
	return s
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (s *server) Handle(conn *websocket.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	c := NewClient(s, conn)
	// add our server client to the list of clients
	s.mux.Lock()
	s.clients = append(s.clients, c)
	s.mux.Unlock()
	// init Client's sender and receiver
	c.Init(10)
	s.handle(c)
	// block until reader is done
	<-c.Dec.Done()
}

func (s *server) Listen(port string, done chan error) {
	http.Handle("/ws", websocket.Handler(s.Handle))
	done <- http.ListenAndServe(port, nil)
}

func (s *server) OnConnection(h func(c *Client)) {
	s.handle = h
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (s *server) BroadCastRandom(sC *Client, m *comm.Message) {
	m.Message = "Random message"
	s.mux.RLock()
	r := random(0, len(s.clients))
	s.clients[r].SendMessage(m)
	s.mux.RUnlock()
}
