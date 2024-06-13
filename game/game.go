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
	websocket *websocket.Conn
	snake     *snake.Snake
}

func NewGame(websocket *websocket.Conn, tickTime time.Duration, gridSize int) *Game {
	grid := grid.NewGrid(gridSize)
	snake := snake.NewSnake(grid)

	return &Game{
		grid:      grid,
		tickTime:  tickTime,
		snake:     snake,
		websocket: websocket,
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
		game.snake.Turn(direction)
	}
}

func (game *Game) endGame() {
	grid := grid.NewGrid(game.grid.Size)
	snake := snake.NewSnake(grid)

	game.grid = grid
	game.snake = snake
}

func (game *Game) tick() {
	game.snake.Move()

	collidedWithWall := game.snake.CheckWallCollision()
	collidedWithSelf := game.snake.CheckSelfCollision()
	if collidedWithWall || collidedWithSelf {
		game.endGame()
	}

	collidedWithFood := game.snake.CheckFoodCollision()
	if collidedWithFood {
		game.snake.Grow()
		game.grid.AddRandomFoodToGameBoard()
	}

	game.snake.UpdatePosition()

	game.sendGameBoardUpdate()
}

func (game *Game) sendGameBoardUpdate() {
	data, _ := json.Marshal(game.grid.GameBoard)
	game.websocket.WriteMessage(websocket.TextMessage, data)
}
