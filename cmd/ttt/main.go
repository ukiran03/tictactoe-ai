package main

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

func main() {
	// Initialize the board
	gameBoard := board.NewBoard()

	// Simulate some moves
	gameBoard.Grid[0][0] = board.XState
	gameBoard.Grid[1][1] = board.OState
	gameBoard.Grid[2][2] = board.XState

	fmt.Println(gameBoard.Render())
}
