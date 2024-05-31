package rover

import (
	"fmt"
	"go-rover/grid"
)

type Rover struct {
	grid      *grid.Grid
	position  [2]int
	direction string
}

func NewRover(grid *grid.Grid) *Rover {
	return &Rover{
		grid:      grid,
		position:  [2]int{0, 0},
		direction: "E",
	}
}

func PrintPosition(position [2]int) {
	fmt.Println("New Rover Position (" + fmt.Sprint(position[0]) + ", " + fmt.Sprint(position[1]) + ")\r")
}

func (r *Rover) Turn(direction string) {
	r.direction = direction
	r.grid.GameBoard[r.position[0]][r.position[1]] = r.getArrowPosition()
	PrintPosition(r.position)
}

func (r *Rover) getArrowPosition() string {
	arrowPosition := ""

	if r.direction == "E" {
		arrowPosition = ">"
	} else if r.direction == "W" {
		arrowPosition = "<"
	} else if r.direction == "N" {
		arrowPosition = "^"
	} else if r.direction == "S" {
		arrowPosition = "v"
	}

	return arrowPosition
}

func (r *Rover) Move() {
	r.grid.GameBoard[r.position[0]][r.position[1]] = 0

	if r.direction == "E" {
		r.position[1] += 1
	} else if r.direction == "W" {
		r.position[1] -= 1
	} else if r.direction == "N" {
		r.position[0] -= 1
	} else if r.direction == "S" {
		r.position[0] += 1
	}

	if r.position[0] < 0 {
		r.position[0] = 0
	}

	if r.position[1] < 0 {
		r.position[1] = 0
	}

	if r.position[0] > r.grid.Size-1 {
		r.position[0] = r.grid.Size - 1
	}

	if r.position[1] > r.grid.Size-1 {
		r.position[1] = r.grid.Size - 1
	}

	r.grid.GameBoard[r.position[0]][r.position[1]] = r.getArrowPosition()
	PrintPosition(r.position)
}
