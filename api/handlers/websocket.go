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

func sendGridUpdate(conn *websocket.Conn, grid *grid.Grid) {
	data, _ := json.Marshal(grid.GameBoard)
	conn.WriteMessage(websocket.TextMessage, data)
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

		log.Printf("Received: %s", message)

		switch msg.Type {
		case "initialize":
		case "move":
			rover.Move()
		case "turn":
			direction := ""
			json.Unmarshal(msg.Data, &direction)
			rover.Turn(direction)
		}
		sendGridUpdate(conn, grid)
	}
}
