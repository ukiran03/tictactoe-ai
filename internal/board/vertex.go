package board

import "fmt"

type Vertex struct {
	X, Y int
}

func (v *Vertex) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}
