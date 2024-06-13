package grid

import (
	"math/rand"
)

type Grid struct {
	Size      int
	GameBoard [][]string
}

func NewGrid(size int) *Grid {
	grid := &Grid{
		Size:      size,
		GameBoard: create(size),
	}
	grid.AddRandomFoodToGameBoard()
	return grid
}

func (g *Grid) AddRandomFoodToGameBoard() {
	x, y := g.generateRandomFoodPosition()
	g.GameBoard[x][y] = "food"
}

func (g *Grid) generateRandomFoodPosition() (int, int) {
	x := rand.Intn(g.Size)
	y := rand.Intn(g.Size)
	for g.GameBoard[x][y] == "snake" {
		x = rand.Intn(g.Size)
		y = rand.Intn(g.Size)
	}
	return x, y
}

func create(size int) [][]string {
	grid := make([][]string, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = "empty"
		}
	}
	grid[0][0] = "snake"

	return grid
}
