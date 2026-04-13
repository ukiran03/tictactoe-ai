package ui

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

func GetMove() (*board.Vertex, error) {
	var x, y int
	fmt.Print("What is your move's X co-ordinate?: ")
	if _, err := fmt.Scan(&x); err != nil {
		return nil, err
	}
	fmt.Print("What is your move's Y co-ordinate?: ")
	if _, err := fmt.Scan(&y); err != nil {
		return nil, err
	}
	return &board.Vertex{X: x, Y: y}, nil
}
