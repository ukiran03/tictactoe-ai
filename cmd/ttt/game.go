package main

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

type player int

const (
	xPlayer player = iota
	oPlayer
	nilPlayer
)

func makeMove(b board.Board, v *board.Vertex, st board.CellState) (board.Board, error) {
	if !isValidMove(b, v) {
		return b,
			fmt.Errorf("Invalid move %v: square already taken or out of bounds", v)
	}

	// Creates a full copy of the struct and the fixed array
	newBoard := b

	// Apply the move to the copy
	newBoard.Grid[v.X][v.Y] = st

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

func getWinner(b board.Board) (bool, string) {
	for _, line := range allLineTriplets {
		if found, state := getWinningPlayer(b, line); found {
			return true, state.Player()
		}
	}
	return false, board.NilState.Player()
}

func getWinningPlayer(b board.Board, line []board.Vertex) (bool, board.CellState) {
	// initials
	initVertex := line[0]
	initState := b.Grid[initVertex.X][initVertex.Y]

	if initState == board.NilState {
		return false, board.NilState
	}

	for i := 1; i < len(line); i++ {
		v := line[i]
		if b.Grid[v.X][v.Y] != initState {
			return false, board.NilState
		}
	}
	return true, initState
}

func isBoardFull(b board.Board) bool {
	for _, line := range b.Grid {
		for _, cell := range line {
			if cell == board.NilState {
				return false
			}
		}
	}
	return true
}
