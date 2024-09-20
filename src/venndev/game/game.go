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
	loopGame     *Loop
	GameOver     bool = false
)

func Init() {
	loopGame = NewLoop(500, func() {
		moveDown()
	})
	grid = comp.NewGrid(20, 10, 30)
	currentBlock = GetRandomBlock()
	nextBlock = GetRandomBlock()
}

func Update() {
	loopGame.Update()
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

func moveLeft() {
	if GameOver {
		return
	}
	currentBlock.Move(0, -1)
	if isBlockOutSide() || !blockFits() {
		currentBlock.Move(0, 1)
	}
}

func moveRight() {
	if GameOver {
		return
	}
	currentBlock.Move(0, 1)
	if isBlockOutSide() || !blockFits() {
		currentBlock.Move(0, -1)
	}
}

func moveDown() {
	if GameOver {
		return
	}
	currentBlock.Move(1, 0)
	if isBlockOutSide() || !blockFits() {
		currentBlock.Move(-1, 0)
		lockBlock()
	}
}

func rotate() {
	if GameOver {
		return
	}
	currentBlock.Rotate()
	if isBlockOutSide() || !blockFits() {
		currentBlock.Rotate()
	}
}

func unRotate() {
	if GameOver {
		return
	}
	currentBlock.UnRotate()
	if isBlockOutSide() || !blockFits() {
		currentBlock.UnRotate()
	}
}

func isBlockOutSide() bool {
	for _, cell := range currentBlock.GetCellPositions() {
		if grid.IsCellOutSide(cell.Row, cell.Col) {
			return true
		}
		println(cell.Row, cell.Col)
	}
	return false
}

func lockBlock() {
	for _, cell := range currentBlock.GetCellPositions() {
		grid.Grid[cell.Row][cell.Col] = currentBlock.GetId()
	}
	currentBlock = nextBlock
	if !blockFits() {
		GameOver = true
	}
	nextBlock = GetRandomBlock()
	grid.ClearFullRows()
}

func blockFits() bool {
	for _, cell := range currentBlock.GetCellPositions() {
		if !grid.IsCellEmpty(cell.Row, cell.Col) {
			return false
		}
	}
	return true
}

func Reset() {
	grid.ClearAll()
	currentBlock = GetRandomBlock()
	nextBlock = GetRandomBlock()
	GameOver = false
}

func HandleInput() {
	key := rl.GetKeyPressed()
	if GameOver && key != 0 {
		GameOver = false
		Reset()
	}
	switch key {
	case rl.KeyLeft:
		moveLeft()
	case rl.KeyRight:
		moveRight()
	case rl.KeyDown:
		moveDown()
	case rl.KeySpace:
		unRotate()
	case rl.KeyUp:
		rotate()
	}
}
