package game

import (
	"encoding/json"
	"go-snake/game/grid"
	"go-snake/game/snake"
	"time"

	"github.com/gorilla/websocket"
)

type Game struct {
	grid      *grid.Grid
	tickTime  time.Duration
	Snake     *snake.Snake
	Websocket *websocket.Conn
}

func NewGame(websocket *websocket.Conn, tickTime time.Duration) *Game {
	if tickTime <= 0 {
		panic("tickTime must be greater than zero")
	}

	grid := grid.NewGrid(10)
	snake := snake.NewSnake(grid)

	return &Game{
		grid:      grid,
		tickTime:  tickTime,
		Snake:     snake,
		Websocket: websocket,
	}
}

func (game *Game) Start() {
	ticker := time.NewTicker(game.tickTime * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		game.Tick()
	}
}

func (game *Game) Tick() {
	game.Snake.Move()
	game.SendGridUpdate()
}

func (game *Game) SendGridUpdate() {
	data, _ := json.Marshal(game.grid.GameBoard)
	game.Websocket.WriteMessage(websocket.TextMessage, data)
}

func (game *Game) HandleWebsocketMessage(actionType string, data json.RawMessage) {
	switch actionType {
	case "move":
		game.Snake.Move()
	case "turn":
		var direction string
		json.Unmarshal(data, &direction)
		game.Snake.Turn(direction)
	}
	game.SendGridUpdate()
}
