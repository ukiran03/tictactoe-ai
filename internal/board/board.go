package board

// CellState: represents the value of a single square
type CellState int

const (
	OState   CellState = iota // 0
	XState                    // 1
	NilState                  // 2

)

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
