package math

type Position struct {
	Row int
	Col int
}

func NewPosition(row int, col int) *Position {
	return &Position{Row: row, Col: col}
}
