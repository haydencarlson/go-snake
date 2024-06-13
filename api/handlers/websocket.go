package handlers

import (
	"encoding/json"
	"go-snake/game"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func upgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return nil, err
	}
	return conn, nil
}

func initializeGame(conn *websocket.Conn) {
	game := game.NewGame(conn, 1000, 10)
	go listenForWebsocketMessages(conn, game)
	game.Start()
}

func listenForWebsocketMessages(websocket *websocket.Conn, game *game.Game) {
	for {
		messageType, message, err := websocket.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Println("Received message type:", messageType)

		var msg WebsocketMessage
		err = json.Unmarshal(message, &msg)

		if err != nil {
			log.Println("Error unmarshalling message:", err)
			break
		}

		log.Printf("Received: %s", message)

		game.HandleWebsocketMessage(msg.Type, msg.Data)
	}
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeConnection(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	initializeGame(conn)
}
