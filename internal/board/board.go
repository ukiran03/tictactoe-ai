package board

import (
	"fmt"
	"io"
	"strings"
)

// CellState: represents the value of a single square
type CellState int

const (
	OState   CellState = iota // 0
	XState                    // 1
	NilState                  // 2

)

var cellSymbols = map[CellState]string{
	OState:   "O",
	XState:   "X",
	NilState: " ",
}

// Board: is a 3x3 grid of square states
type Board [3][3]CellState

// NewBoard: initializes a Board filled with NilState
func NewBoard() Board {
	var b Board
	for i := range 3 {
		for j := range 3 {
			b[i][j] = NilState
		}
	}
	return b
}

// String: provides a basic printing of the Board
func (b Board) String() string {
	symbols := map[CellState]string{
		OState:   "O",
		XState:   "X",
		NilState: " ",
	}
	var out strings.Builder
	for _, row := range b {
		for _, cs := range row {
			fmt.Fprintf(&out, " %s ", symbols[cs])
		}
		out.WriteString("\n")
	}
	return out.String()
}

// render: renders pretty version of the board
func (b Board) Render() string {
	var out strings.Builder

	// Boarders
	top := NewBorder(topBorder)
	mid := NewBorder(midBorder)
	bot := NewBorder(botBorder)

	// Vertical Num Row
	writeHIndex(&out, 3, 1, 3)

	out.WriteString(top.render(3)) // Top border
	for i, row := range b {
		writeCellRow(&out, fmt.Sprintf("%d ", i), 1, row)
		out.WriteString(mid.render(3))
	}
	out.WriteString(bot.render(3)) // Bot border

	return out.String()
}

// Usage: writeCellRow(&b, 1, state)
func writeCellRow(b io.Writer, prefix string, pad int, state [3]CellState) {
	fmt.Fprint(b, prefix)
	for _, cs := range state {
		// %s is the Vbar
		// %*s handles the padding + the symbol
		fmt.Fprintf(b, "%s %*s ", Vbar, pad, cellSymbols[cs])
	}
	// Add the closing bar for the far right
	fmt.Fprint(b, Vbar)
}

// writeHIndex: writes Horizontal Index row
// rng: 0, 1, 2,..
// pad: padding length
// prefix: prefix length for each integer
// Usage: writeHIndex(&b, 3, 1, 3)
func writeHIndex(b io.Writer, rng, pad, prefix int) {
	fmt.Fprintf(b, "%*s", pad, "")
	for i := range rng {
		fmt.Fprintf(b, "%*s%d", prefix, "", i)
	}
}
