package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("got conn in handleWS", conn)
	defer conn.Close()

	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("read this type of message", messageType)
	fmt.Println("read this:", p)
}

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(w, r)
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
