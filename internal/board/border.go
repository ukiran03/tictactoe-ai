package board

import (
	"fmt"
	"io"
)

type border struct {
	pre, middle, post, filler string
}

type borderType int

const (
	topBorder borderType = iota
	midBorder
	botBorder
)

// NewBorder returns the specific characters for table boundaries.
func NewBorder(bp borderType) border {
	switch bp {
	case topBorder:
		return border{"┌", "┬", "┐", Hbar}
	case midBorder:
		return border{"├", "┼", "┤", Hbar}
	case botBorder:
		return border{"└", "┴", "┘", Hbar}
	default:
		return border{" ", " ", " ", " "}
	}
}

func (br border) render(w io.Writer, pad, width int) {
	fmt.Fprintf(w, "%*s", pad, br.pre)
	for i := range width {
		for range width {
			fmt.Fprint(w, br.filler)
		}
		if i < width-1 {
			fmt.Fprint(w, br.middle)
		}
	}
	fmt.Fprintf(w, "%s\n", br.post)
}
