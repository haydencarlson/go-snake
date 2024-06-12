package handlers

import (
	"encoding/json"
	"go-snake/game"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	game := game.NewGame(conn)
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Println("Received message type:", messageType)

		var msg Message
		json.Unmarshal(message, &msg)

		log.Printf("Received: %s", message)

		switch msg.Type {
		case "move":
			game.Snake.Move()
		case "turn":
			direction := ""
			json.Unmarshal(msg.Data, &direction)
			game.Snake.Turn(direction)
		}
		game.SendGridUpdate()
	}
}
