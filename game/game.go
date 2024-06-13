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

func NewGame(websocket *websocket.Conn, tickTime time.Duration, gridSize int) *Game {
	grid := grid.NewGrid(gridSize)
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
		game.tick()
	}
}

func (game *Game) HandleWebsocketMessage(actionType string, data json.RawMessage) {
	switch actionType {
	case "turn":
		var direction string
		json.Unmarshal(data, &direction)
		game.Snake.Turn(direction)
	}
}

func (game *Game) endGame() {
	grid := grid.NewGrid(game.grid.Size)
	snake := snake.NewSnake(grid)

	game.grid = grid
	game.Snake = snake
}

func (game *Game) tick() {
	game.Snake.Move()
	collidedWithWall := game.Snake.CheckWallCollision()

	if collidedWithWall {
		game.endGame()
	}

	collidedWithFood := game.Snake.CheckFoodCollision()
	if collidedWithFood {
		game.grid.AddRandomFoodToGameBoard()
	}

	game.Snake.UpdatePosition()

	game.sendGameBoardUpdate()
}

func (game *Game) sendGameBoardUpdate() {
	data, _ := json.Marshal(game.grid.GameBoard)
	game.Websocket.WriteMessage(websocket.TextMessage, data)
}
