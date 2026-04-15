package ui

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

func GetMove() (*board.Vertex, error) {
	var x, y int

	fmt.Print("Enter X coordinate: ")
	if _, err := fmt.Scan(&x); err != nil {
		return nil, fmt.Errorf("invalid X input: %w", err)
	}

	fmt.Print("Enter Y coordinate: ")
	if _, err := fmt.Scan(&y); err != nil {
		return nil, fmt.Errorf("invalid Y input: %w", err)
	}

	// Basic validation logic
	if x < 0 || y < 0 || x >= 3 || y >= 3 {
		return nil, fmt.Errorf("move (%d, %d) is out of bounds", x, y)
	}

	return &board.Vertex{X: x, Y: y}, nil
}

// [13-04-2026] TODO: start here
// https://robertheaton.com/2018/10/09/programming-projects-for-advanced-beginners-3-a/#:~:text=square%20already%20taken!%22-,5,-.%20Combining%20our%20work
