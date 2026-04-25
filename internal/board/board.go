package board

// CellState: represents the value of a single square
type CellState int

const (
	OState   CellState = iota // 0
	XState                    // 1
	NilState                  // 2

)

func (c CellState) Player() string {
	return c.String()
}

func (c CellState) String() string {
	switch c {
	case XState:
		return "X"
	case OState:
		return "O"
	default:
		return "-"
	}
}

// Next tells you who follows this state
func (c CellState) Next() CellState {
	if c == XState {
		return OState
	}
	return XState
}

type Board struct {
	Grid  [3][3]CellState
	Width int
}

func NewBoard() Board {
	var grid [3][3]CellState
	for i := range 3 {
		for j := range 3 {
			grid[i][j] = NilState
		}
	}
	return Board{
		Grid:  grid,
		Width: 3,
	}
}
