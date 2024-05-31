package grid

type Grid struct {
	Size      int
	GameBoard [][]interface{}
}

func NewGrid(size int) *Grid {
	grid := &Grid{
		Size:      size,
		GameBoard: create(size),
	}
	return grid
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
