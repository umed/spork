package main

import (
	"fmt"
	"log"
	"net/url"

	"spork/cmd/config"

	"github.com/gorilla/websocket"
)

func handleUserInput(c *websocket.Conn) {
	for {
		var outputMessage string
		fmt.Scanln(&outputMessage)
		c.WriteMessage(websocket.TextMessage, []byte(outputMessage))
	}
}

func main() {
	log.SetFlags(0)
	var addr = config.ParseAddr()
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	go handleUserInput(c)
	c.WriteMessage(websocket.TextMessage, []byte("client joined"))
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}
