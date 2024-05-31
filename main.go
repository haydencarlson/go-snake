package main

import (
	"fmt"
	"os"
	"syscall"

	"go-rover/grid"
	"go-rover/rover"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	var buf [1]byte

	fmt.Println("Move your rover with the arrow keys (press 'Ctrl+C' to quit):\r")

	grid := grid.NewGrid(10)
	rover := rover.NewRover(grid)

	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}

		if buf[0] == ' ' {
			rover.Move()
		} else if buf[0] == '\u001b' {
			// Read the next two bytes to determine the arrow key
			os.Stdin.Read(buf[:])
			os.Stdin.Read(buf[:])

			switch buf[0] {
			case 'A':
				rover.Turn(string('N'))
			case 'B':
				rover.Turn(string('S'))
			case 'C':
				rover.Turn(string('E'))
			case 'D':
				rover.Turn(string('W'))
			default:
				fmt.Printf("You pressed: %c\n", buf[0])
			}
		}

		if buf[0] == 3 {
			break
		}
	}
}
