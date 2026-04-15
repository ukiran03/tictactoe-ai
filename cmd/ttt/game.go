package main

import (
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

func makeMove(b *board.Board, v *board.Vertex, p player) *board.Board {
	// This creates a brand new copy of the Board struct and its [3][3] grid
	newBoard := *b

	// Apply the move to the copy
	newBoard.Grid[v.X][v.Y] = playerStates[p]

	// Return the address of the new copy
	return &newBoard
}

func isValidMove(b *board.Board, v *board.Vertex) bool {
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
