package snake

import (
	"fmt"
	"go-snake/game/grid"
	"log"
)

type Snake struct {
	grid      *grid.Grid
	position  [2]int
	direction string
}

func NewSnake(grid *grid.Grid) *Snake {
	return &Snake{
		grid:      grid,
		position:  [2]int{0, 0},
		direction: "E",
	}
}

func (r *Snake) Move() {
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
	printPosition(r.position)
}

func printPosition(position [2]int) {
	fmt.Println("New Snake Position (" + fmt.Sprint(position[0]) + ", " + fmt.Sprint(position[1]) + ")\r")
}

func printDirection(direction string) {
	fmt.Println("New Snake Direction: " + direction + "\r")
}

func (r *Snake) Turn(direction string) {
	r.direction = direction
	r.grid.GameBoard[r.position[0]][r.position[1]] = r.getArrowPosition()
	printDirection(r.direction)
}

func (r *Snake) getArrowPosition() string {
	arrowPosition := ""

	if r.direction == "E" {
		log.Println("E turning E")
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
