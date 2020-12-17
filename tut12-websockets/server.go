package main

import (
	"fmt"
	"net/http"

	// websocket library
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//var chatMessages []string

func main() {

	// For template rendering, the data passed can be any kind of Go data structure
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))
			/*/ add it to the chat messages
			chatMessages = append(chatMessages, string(msg))*/

			// Write the message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
		// Send existing chat messages
		/*for i := 0; i < len(chatMessages); i = i + 1 {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return
			}

			// Write the message back to browser
			if err = conn.WriteMessage(1, []byte(chatMessages[i])); err != nil {
				return
			}
		}*/
	})

	http.ListenAndServe(":80", nil)
}
