package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: ":8000", Path: "/ws"}
	fmt.Println(u)
	conn, res, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println(conn)
	defer conn.Close()

	var buffer []byte = make([]byte, 4)
	for i := 0; i < 4; i++ {
		buffer[i] = byte(i)
	}
	conn.WriteMessage(websocket.BinaryMessage, buffer)
}
