package model

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Server stores clients
type Server struct {
	upgrader websocket.Upgrader
	clients  []*Client
	addr     *string
	receive  chan []byte
}

// NewServer creates new chat server
func NewServer(addr *string) *Server {
	var server = Server{upgrader: websocket.Upgrader{}, addr: addr, receive: make(chan []byte)}
	http.HandleFunc("/", server.register)
	return &server
}

// Register registers new clients
func (server *Server) register(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to establish connection with client: ", err)
		return
	}
	var client = Client{conn: conn, receive: server.receive, send: make(chan []byte)}
	go client.Handle()
	server.clients = append(server.clients, &client)
}

func (server *Server) broadcast(message []byte) {
	for _, client := range server.clients {
		client.send <- message
	}
}

func (server *Server) run() {
	for {
		select {
		case message := <-server.receive:
			server.broadcast(message)
		}
	}
}

// Run runs chat's server
func (server *Server) Run() {
	defer server.close()
	go server.run()
	log.Fatal(http.ListenAndServe(*server.addr, nil))
}

// Close closes all connections of clients
func (server *Server) close() {
	for len(server.clients) > 0 {
		server.clients[0].close()
		server.clients = server.clients[1:]
	}
}
