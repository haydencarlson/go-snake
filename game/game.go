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
	tickTime  int
	websocket *websocket.Conn
	Snake     *snake.Snake
	score     int
}

func NewGame(websocket *websocket.Conn, tickTime int, gridSize int) *Game {
	grid := grid.NewGrid(gridSize)
	snake := snake.NewSnake(grid)

	return &Game{
		grid:      grid,
		tickTime:  tickTime,
		Snake:     snake,
		websocket: websocket,
		score:     0,
	}
}

func (game *Game) Start() {
	ticker := time.NewTicker(time.Duration(game.tickTime) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		game.tick()
	}
}

func (game *Game) restartGame() {
	grid := grid.NewGrid(game.grid.Size)
	snake := snake.NewSnake(grid)

	game.score = 0
	game.grid = grid
	game.Snake = snake
}

func (game *Game) tick() {
	game.Snake.Move()

	collidedWithWall := game.Snake.CheckWallCollision()
	collidedWithSelf := game.Snake.CheckSelfCollision()
	if collidedWithWall || collidedWithSelf {
		game.restartGame()
	}

	collidedWithFood := game.Snake.CheckFoodCollision()
	if collidedWithFood {
		game.score++
		game.Snake.Grow()
	}

	game.Snake.UpdateBoardPosition()

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
