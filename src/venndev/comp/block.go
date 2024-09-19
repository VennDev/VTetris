package comp

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/utils"
	"github.com/venndev/VTetris/src/venndev/utils/math"
)

type Block struct {
	id            int
	cells         map[int]math.Position
	cellSize      int
	rotationState int
	color         []rl.Color
	rowOffset     int
	colOffset     int
}

func (b *Block) GetId() int {
	return b.id
}

func (b *Block) GetCells() map[int]math.Position {
	return b.cells
}

// This function is used to draw the block
func (b *Block) Draw() {
	for _, cell := range b.GetCellPositions() {
		rl.DrawRectangle(
			int32(cell.Col*b.cellSize+1),
			int32(cell.Row*b.cellSize+1),
			int32(b.cellSize-1),
			int32(b.cellSize-1),
			b.color[b.id],
		)
	}
}

// This function is used to rotate the block
func (b *Block) Move(rows int, cols int) {
	b.rowOffset += rows
	b.colOffset += cols
}

func (b *Block) Rotate() {
	b.rotationState = (b.rotationState + 1) % 4
}

// This function is used to get the positions of the cells in the block
func (b *Block) GetCellPositions() []math.Position {
	cellPositions := make([]math.Position, len(b.cells))
	for i, cell := range b.cells {
		cellPositions[i] = math.Position{cell.Row + b.rowOffset, cell.Col + b.colOffset}
	}
	return cellPositions
}

func NewBlock(id int, cells map[int]math.Position, cellSize int, rotationState int) *Block {
	block := &Block{id, cells, cellSize, rotationState, utils.GetCellColors(), 0, 0}
	block.Move(0, 3)
	return block
}
