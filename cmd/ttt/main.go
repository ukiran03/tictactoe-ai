package main

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
	"ukiran.com/tictactoe/internal/ui"
)

func main() {
	// Initialize the board
	gameBoard := board.NewBoard()

	// Simulate some moves
	gameBoard.Grid[0][0] = board.XState
	gameBoard.Grid[1][1] = board.OState
	gameBoard.Grid[2][2] = board.XState

	gameBoard.Render()

	var move *board.Vertex
	for {
		var err error
		move, err = ui.GetMove()
		if err != nil {
			fmt.Printf("Input error: %v. Please enter numbers only.\n", err)
			continue // Try again
		}

		if isValidMove(gameBoard, move) {
			break // Success! Exit the loop
		}

		fmt.Println("That square is already taken or out of bounds!")
	}
	gameBoard = makeMove(gameBoard, move, xPlayer)
	gameBoard.Render()
}
