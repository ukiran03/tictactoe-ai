package board

import (
	"fmt"
	"strings"
)

const (
	Vbar = "│"
	Hbar = "─"
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

func (br border) render(fillWid int) string {
	var out strings.Builder

	out.WriteString(br.pre)
	for i := range fillWid {
		for range fillWid {
			out.WriteString(br.filler)
		}
		if i < fillWid-1 {
			out.WriteString(br.middle)
		}
	}
	out.WriteString(br.post)
	return out.String()
}

// Helper to create a horizontal line (e.g., ┌───┬───┐)
func renderBorder(b border) {
	line := b.pre
	for i := 0; i < len(gb.Cells[0]); i++ {
		for j := 0; j < gb.Width; j++ {
			line += "─"
		}
		if i < len(gb.Cells[0])-1 {
			line += b.middle
		}
	}
	fmt.Println(line + b.post)
}
