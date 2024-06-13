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
	score     int
}

func NewGame(websocket *websocket.Conn, tickTime time.Duration, gridSize int) *Game {
	grid := grid.NewGrid(gridSize)
	snake := snake.NewSnake(grid)

	return &Game{
		grid:      grid,
		tickTime:  tickTime,
		snake:     snake,
		websocket: websocket,
		score:     0,
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

func (game *Game) restartGame() {
	grid := grid.NewGrid(game.grid.Size)
	snake := snake.NewSnake(grid)

	game.score = 0
	game.grid = grid
	game.snake = snake
}

func (game *Game) tick() {
	game.snake.Move()

	collidedWithWall := game.snake.CheckWallCollision()
	collidedWithSelf := game.snake.CheckSelfCollision()
	if collidedWithWall || collidedWithSelf {
		game.restartGame()
	}

	collidedWithFood := game.snake.CheckFoodCollision()
	if collidedWithFood {
		game.score++
		game.snake.Grow()
	}

	game.snake.UpdateBoardPosition()

	if collidedWithFood {
		game.grid.AddRandomFoodToGameBoard()
	}

	game.sendWebsocketUpdate()
}

func (game *Game) sendWebsocketUpdate() {
	game.sendGameBoardUpdate()
	game.sendScoreUpdate()
}

func (game *Game) sendGameBoardUpdate() {
	data, _ := json.Marshal(struct {
		Type  string      `json:"type"`
		Board interface{} `json:"board"`
	}{
		Type:  "gameBoard",
		Board: game.grid.GameBoard,
	})
	game.websocket.WriteMessage(websocket.TextMessage, data)
}

func (game *Game) sendScoreUpdate() {
	data, _ := json.Marshal(struct {
		Type  string `json:"type"`
		Score int    `json:"score"`
	}{
		Type:  "score",
		Score: game.score,
	})
	game.websocket.WriteMessage(websocket.TextMessage, data)
}
