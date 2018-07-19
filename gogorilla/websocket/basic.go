package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		go func() {
			for {
				messageType, p, _ := conn.ReadMessage()

				conn.WriteMessage(messageType, p)
			}
		}()
	})

	http.ListenAndServe(":3000", nil)
}
