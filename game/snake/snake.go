package snake

import (
	"fmt"
	"go-snake/game/grid"
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

func (s *Snake) Move() {
	s.grid.GameBoard[s.position[0]][s.position[1]] = "empty"

	switch s.direction {
	case "E":
		s.position[1]++
	case "W":
		s.position[1]--
	case "N":
		s.position[0]--
	case "S":
		s.position[0]++
	}

	printPosition(s.position)
}

func (s *Snake) UpdatePosition() {
	s.grid.GameBoard[s.position[0]][s.position[1]] = "snake"
}

func (r *Snake) CheckWallCollision() bool {
	if r.position[0] < 0 || r.position[0] >= len(r.grid.GameBoard) {
		return true
	}
	if r.position[1] < 0 || r.position[1] >= len(r.grid.GameBoard[0]) {
		return true
	}

	return false
}

func (r *Snake) CheckFoodCollision() bool {
	if r.grid.GameBoard[r.position[0]][r.position[1]] == "food" {
		return true
	}

	return false
}

func (r *Snake) Turn(direction string) {
	r.direction = direction
	printDirection(r.direction)
}

func printPosition(position [2]int) {
	fmt.Println("New Snake Position (" + fmt.Sprint(position[0]) + ", " + fmt.Sprint(position[1]) + ")\r")
}

func printDirection(direction string) {
	fmt.Println("New Snake Direction: " + direction + "\r")
}
