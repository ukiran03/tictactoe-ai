package main

import (
	"ukiran.com/tictactoe/internal/board"
)

var allLineTriplets = getAllLineTriplets()

var (
	BOARD_WIDTH  = 3
	BOARD_HEIGHT = 3
)

func getAllLineTriplets() [][]board.Vertex {
	//  3 (cols) + 3 (rows) + 2 (diags) = 8 triplets
	triplets := make([][]board.Vertex, 0, 8)

	// Columns
	for x := 0; x < BOARD_WIDTH; x++ {
		col := make([]board.Vertex, 0, BOARD_HEIGHT)
		for y := 0; y < BOARD_HEIGHT; y++ {
			col = append(col, board.Vertex{X: x, Y: y})
		}
		triplets = append(triplets, col)
	}

	// Rows
	for y := 0; y < BOARD_HEIGHT; y++ {
		row := make([]board.Vertex, 0, BOARD_WIDTH)
		for x := 0; x < BOARD_WIDTH; x++ {
			row = append(row, board.Vertex{X: x, Y: y})
		}
		triplets = append(triplets, row)
	}

	// Diagonals
	diagonals := [][]board.Vertex{
		{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}},
		{{X: 0, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 0}},
	}

	return append(triplets, diagonals...)
}
