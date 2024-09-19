package game

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/comp"
)

var (
	blocks       []*comp.Block
	currentBlock *comp.Block
	nextBlock    *comp.Block
	grid         *comp.Grid
)

func Init() {
	grid = comp.NewGrid(20, 10, 30)
	currentBlock = GetRandomBlock()
	nextBlock = GetRandomBlock()
}

func Draw() {
	grid.Draw()
	currentBlock.Draw()
}

func GetAllBlocks() []*comp.Block {
	return []*comp.Block{
		comp.LBlock(),
		comp.JBlock(),
		comp.ZBlock(),
		comp.SBlock(),
		comp.TBlock(),
		comp.OBlock(),
		comp.IBlock(),
	}
}

func GetRandomBlock() *comp.Block {
	if len(blocks) == 0 {
		blocks = GetAllBlocks()
	}
	index := rand.Intn(len(blocks))
	block := blocks[index]
	blocks = append(blocks[:index], blocks[index+1:]...) // remove block from list
	return block
}

func MoveLeft() {
	currentBlock.Move(0, -1)
	if IsBlockOutSide() {
		currentBlock.Move(0, 1)
	}
}

func MoveRight() {
	currentBlock.Move(0, 1)
	if IsBlockOutSide() {
		currentBlock.Move(0, -1)
	}
}

func MoveDown() {
	currentBlock.Move(1, 0)
	if IsBlockOutSide() {
		currentBlock.Move(-1, 0)
	}
}

func IsBlockOutSide() bool {
	for _, cell := range currentBlock.GetCellPositions() {
		if grid.IsCellOutSide(cell.Row, cell.Col) {
			return true
		}
		println(cell.Row, cell.Col)
	}
	return false
}

func HandleInput() {
	key := rl.GetKeyPressed()
	switch key {
	case rl.KeyLeft:
		MoveLeft()
	case rl.KeyRight:
		MoveRight()
	case rl.KeyDown:
		MoveDown()
	}
}
