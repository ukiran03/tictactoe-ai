package board

import (
	"fmt"
	"io"
	"strings"
)

var cellSymbols = map[CellState]string{
	OState:   "O",
	XState:   "X",
	NilState: " ",
}

const (
	Vbar = "│"
	Hbar = "─"
)

func (b Board) Render() {
	var out strings.Builder

	// Boarders
	top := NewBorder(topBorder)
	mid := NewBorder(midBorder)
	bot := NewBorder(botBorder)

	// Vertical Num Row
	writeHIndex(&out, 3, 1, 3)

	top.render(&out, 3, 3) // Top border
	for i, row := range b.Grid {
		writeCellRow(&out, fmt.Sprintf("%d ", i), 1, row)

		if i < len(b.Grid)-1 {
			mid.render(&out, 3, 3) // Mid border
		}
	}
	bot.render(&out, 3, 3) // Bot border
	fmt.Println(out.String())
}

// Usage: writeCellRow(&w, 1, state)
func writeCellRow(w io.Writer, prefix string, pad int, state [3]CellState) {
	fmt.Fprint(w, prefix)
	for _, cs := range state {
		// %s is the Vbar
		// %*s handles the padding + the symbol
		fmt.Fprintf(w, "%s %*s ", Vbar, pad, cellSymbols[cs])
	}
	// Add the closing bar for the far right
	fmt.Fprintf(w, "%s\n", Vbar)
}

// writeHIndex: writes Horizontal Index row
// rng: 0, 1, 2,..
// pad: padding length
// prefix: prefix length for each integer
// Usage: writeHIndex(&w, 3, 1, 3)
func writeHIndex(w io.Writer, rng, pad, prefix int) {
	fmt.Fprintf(w, "%*s", pad, "")
	for i := range rng {
		fmt.Fprintf(w, "%*s%d", prefix, "", i)
	}
	fmt.Fprint(w, "\n")
}
