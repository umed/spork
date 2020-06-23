package model

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client stores clients connection
type Client struct {
	conn    *websocket.Conn
	send    chan []byte
	receive chan []byte
}

// Handle handles messages from clients
func (c *Client) Handle() {
	go c.read()
	go c.write()
}

func (c *Client) read() {
	defer c.close()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from client: ", err)
			break
		}
		c.receive <- message
	}
}

func (c *Client) write() {
	defer c.close()
	for {
		select {
		case message := <-c.send:
			err := c.conn.WriteMessage(websocket.TextMessage, HandleMessageSpaces(message))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

// Close closes clinet's connection
func (c *Client) close() {
	c.conn.Close()
}
