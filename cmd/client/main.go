package main

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"spork/cmd/config"
	"spork/internal/model"

	"github.com/gorilla/websocket"
)

func handleUserInput(c *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			outputMessage := scanner.Bytes()
			c.WriteMessage(websocket.TextMessage, outputMessage)
		}
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
	// c.SetReadLimit(1024)
	c.WriteMessage(websocket.TextMessage, []byte("client joined"))
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", model.HandleMessageSpaces(message))
	}
}
