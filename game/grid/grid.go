package grid

import (
	"math/rand"
)

type Grid struct {
	Size      int
	GameBoard [][]interface{}
}

func NewGrid(size int) *Grid {
	grid := &Grid{
		Size:      size,
		GameBoard: create(size),
	}
	grid.AddRandomObstacleToGrid()
	return grid
}

func (g *Grid) AddRandomObstacleToGrid() {
	x, y := g.generateRandomObstaclePosition()
	g.GameBoard[x][y] = "O"
}

func (g *Grid) generateRandomObstaclePosition() (int, int) {
	x := rand.Intn(g.Size)
	y := rand.Intn(g.Size)
	for g.GameBoard[x][y] != 0 {
		x = rand.Intn(g.Size)
		y = rand.Intn(g.Size)
	}
	return x, y
}

func create(size int) [][]interface{} {
	grid := make([][]interface{}, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]interface{}, size)
		for j := 0; j < size; j++ {
			grid[i][j] = 0
		}
	}
	grid[0][0] = ">"
	return grid
}
