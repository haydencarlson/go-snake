package handlers

import (
	"encoding/json"
	"go-rover/game/grid"
	"go-rover/game/rover"
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

	grid := grid.NewGrid(10)
	rover := rover.NewRover(grid)

	for {
		// Read message from WebSocket
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Println("Received message type:", messageType)

		var msg Message
		err = json.Unmarshal(message, &msg)

		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Log the received message
		log.Printf("Received: %s", message)

		switch msg.Type {
		case "initialize":
			data, _ := json.Marshal(grid.GameBoard)
			log.Println("Sending grid data:", string(data))
			conn.WriteMessage(messageType, data)
		case "move":
			rover.Move()
			data, _ := json.Marshal(grid.GameBoard)
			conn.WriteMessage(messageType, data)
		case "turn":
			direction := ""
			json.Unmarshal(msg.Data, &direction)
			rover.Turn(direction)
			data, _ := json.Marshal(grid.GameBoard)
			conn.WriteMessage(messageType, data)
		}
	}
}
