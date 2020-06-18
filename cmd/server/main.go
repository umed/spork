package main

import (
	"log"
	"net/http"
	"spork/cmd/config"

	"github.com/gorilla/websocket"
)

func run(upgrader *websocket.Upgrader, w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.SetFlags(0)
	var addr = config.ParseAddr()
	var upgrader = websocket.Upgrader{} // use default options
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		run(&upgrader, w, r)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
