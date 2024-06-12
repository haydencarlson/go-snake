package game

import (
	"encoding/json"
	"go-snake/game/grid"
	"go-snake/game/snake"

	"github.com/gorilla/websocket"
)

type Game struct {
	grid      *grid.Grid
	Snake     *snake.Snake
	websocket *websocket.Conn
}

func NewGame(websocket *websocket.Conn) *Game {
	grid := grid.NewGrid(10)
	snake := snake.NewSnake(grid)

	return &Game{
		grid:      grid,
		Snake:     snake,
		websocket: websocket,
	}
}

func (game *Game) SendGridUpdate() {
	data, _ := json.Marshal(game.grid.GameBoard)
	game.websocket.WriteMessage(websocket.TextMessage, data)
}
