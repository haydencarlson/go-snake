package snake

import (
	"fmt"
	"go-snake/game/grid"
)

type Snake struct {
	grid      *grid.Grid
	head      [2]int
	body      [][2]int
	direction string
}

func NewSnake(grid *grid.Grid) *Snake {
	initialPosition := [2]int{0, 0}

	return &Snake{
		grid:      grid,
		head:      initialPosition,
		body:      [][2]int{initialPosition},
		direction: "E",
	}
}

func (s *Snake) Move() {
	for _, pos := range s.body {
		s.grid.GameBoard[pos[0]][pos[1]] = "empty"
	}

	newHead := s.head
	switch s.direction {
	case "E":
		newHead[1]++
	case "W":
		newHead[1]--
	case "N":
		newHead[0]--
	case "S":
		newHead[0]++
	}

	s.body = append([][2]int{newHead}, s.body[:len(s.body)-1]...)
	s.head = newHead

	printPosition(s.head)
}

func (s *Snake) Turn(direction string) {
	if direction == "N" && s.direction == "S" {
		return
	}

	if direction == "S" && s.direction == "N" {
		return
	}

	if direction == "E" && s.direction == "W" {
		return
	}

	if direction == "W" && s.direction == "E" {
		return
	}

	s.direction = direction
	printDirection(s.direction)
}

func (s *Snake) UpdatePosition() {
	for _, pos := range s.body {
		s.grid.GameBoard[pos[0]][pos[1]] = "snake"
	}
}

func (s *Snake) CheckWallCollision() bool {
	if s.head[0] < 0 || s.head[0] >= len(s.grid.GameBoard) {
		return true
	}
	if s.head[1] < 0 || s.head[1] >= len(s.grid.GameBoard[0]) {
		return true
	}

	return false
}

func (s *Snake) CheckFoodCollision() bool {
	return s.grid.GameBoard[s.head[0]][s.head[1]] == "food"
}

func (s *Snake) CheckSelfCollision() bool {
	for _, pos := range s.body[1:] {
		if pos == s.head {
			return true
		}
	}

	return false
}

func (s *Snake) Grow() {
	s.body = append(s.body, s.body[len(s.body)-1])
}

func printPosition(head [2]int) {
	fmt.Println("New Snake Position (" + fmt.Sprint(head[0]) + ", " + fmt.Sprint(head[1]) + ")\r")
}

func printDirection(direction string) {
	fmt.Println("New Snake Direction: " + direction + "\r")
}
