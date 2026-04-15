package main

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

type player int

const (
	xPlayer player = iota
	oPlayer
)

var playerStates = map[player]board.CellState{
	xPlayer: board.XState,
	oPlayer: board.OState,
}

func makeMove(b board.Board, v *board.Vertex, p player) (board.Board, error) {
	if !isValidMove(b, v) {
		return b,
			fmt.Errorf("Invalid move %v: square already taken or out of bounds", v)
	}

	// Creates a full copy of the struct and the fixed array
	newBoard := b

	// Apply the move to the copy
	newBoard.Grid[v.X][v.Y] = playerStates[p]

	return newBoard, nil
}

func isValidMove(b board.Board, v *board.Vertex) bool {
	if v == nil {
		return false
	}
	// Bounds check
	if v.X < 0 || v.X >= len(b.Grid) ||
		v.Y < 0 || v.Y >= len(b.Grid[0]) {
		return false
	}
	// Check if spot is empty
	return b.Grid[v.X][v.Y] == board.NilState
}
