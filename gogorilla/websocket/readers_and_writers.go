package main

import (
	"io"
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
				messageType, reader, _ := conn.NextReader() // receive as much data as it has

				writer, _ := conn.NextWriter(messageType)
				io.Copy(writer, reader) // can write as much data as you want
				writer.Close()          // the websocket flush the buffer and send back the message
			}
		}()
	})

	http.ListenAndServe(":3000", nil)
}
