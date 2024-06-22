package handlers

import (
	"encoding/json"
	"go-snake/game"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
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

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeConnection(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	gameVars, err := godotenv.Read()

	if err != nil {
		log.Println("Error reading game variables:", err)
		return
	}

	tickTime, err := strconv.Atoi(gameVars["TICK_TIME"])
	if err != nil {
		log.Println("Error converting TICK_TIME to int:", err)
		return
	}

	gridSize, err := strconv.Atoi(gameVars["GRID_SIZE"])
	if err != nil {
		log.Println("Error converting grid size to int:", err)
		return
	}

	newGame(conn, tickTime, gridSize)
}

func upgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return nil, err
	}
	return conn, nil
}

func handleWebsocketMessage(game *game.Game, actionType string, data json.RawMessage) {
	switch actionType {
	case "turn":
		var direction string
		json.Unmarshal(data, &direction)
		game.Snake.Turn(direction)
	}
}

func newGame(conn *websocket.Conn, tickTime int, gridSize int) {
	log.Printf("Starting game with grid size %s and tick time %s", strconv.Itoa(gridSize), strconv.Itoa(tickTime))

	game := game.NewGame(conn, tickTime, gridSize)
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

		handleWebsocketMessage(game, msg.Type, msg.Data)
	}
}
