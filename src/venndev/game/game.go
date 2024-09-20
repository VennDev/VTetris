package game

import (
	"image/color"
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/comp"
	"github.com/venndev/VTetris/src/venndev/utils"
)

var (
	blocks       []*comp.Block
	currentBlock *comp.Block
	nextBlock    *comp.Block
	grid         *comp.Grid
	loopGame     *Loop
	InGame       bool = false
	GameOver     bool = false
	Score        int  = 0
	Music        rl.Music
	RotateSound  rl.Sound
	ClearSound   rl.Sound
	Font         rl.Font
)

func Init() {
	Font = rl.LoadFontEx("resources/font/monogram.ttf", 64, nil, 0)
	loopGame = NewLoop(500, func() {
		moveDown()
	})
	grid = comp.NewGrid(20, 10, 30)
	currentBlock = GetRandomBlock()
	nextBlock = GetRandomBlock()
	rl.InitAudioDevice()
	Music = rl.LoadMusicStream("resources/sounds/music.mp3")
	RotateSound = rl.LoadSound("resources/sounds/rotate.mp3")
	ClearSound = rl.LoadSound("resources/sounds/clear.mp3")
	rl.PlayMusicStream(Music)
	rl.SetMusicVolume(Music, 0.2)
}

func Close() {
	rl.UnloadMusicStream(Music)
	rl.UnloadSound(RotateSound)
	rl.UnloadSound(ClearSound)
	rl.CloseAudioDevice()
}

func Update() {
	loopGame.Update()
}

func Load() {
	rl.BeginDrawing()
	if GameOver {
		GameOverLoad()
		return
	}
	rl.UpdateMusicStream(Music)
	rl.ClearBackground(utils.DarkBlue)
	rl.DrawTextEx(
		Font,
		"Score:",
		rl.Vector2{X: 360, Y: 15},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	rl.DrawTextEx(
		Font,
		"Next:",
		rl.Vector2{X: 360, Y: 150},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	rl.DrawRectangleRounded(
		rl.Rectangle{X: 320, Y: 55, Width: 170, Height: 60},
		0.3,
		6,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	scoreText := []rune(strconv.Itoa(Score))
	textSize := rl.MeasureTextEx(Font, string(scoreText), 38, 2)
	rl.DrawTextEx(
		Font,
		string(scoreText),
		rl.Vector2{X: 320 + (170-textSize.X)/2, Y: 55 + (60-textSize.Y)/2},
		38,
		2,
		color.RGBA{R: 0, G: 0, B: 0, A: 255},
	)
	rl.DrawRectangleRounded(
		rl.Rectangle{X: 320, Y: 190, Width: 170, Height: 360},
		0.3,
		6,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	Draw()
	HandleInput()
	Update()
}

func Draw() {
	grid.Draw()
	currentBlock.Draw(11, 11)
	switch nextBlock.GetId() {
	case comp.ZBLOCK:
		nextBlock.Draw(265, 295)
	case comp.SBLOCK:
		nextBlock.Draw(265, 285)
	default:
		nextBlock.Draw(265, 275)
	}
}

func UpdateScore(lineCleared, moveDownPoints int) {
	switch lineCleared {
	case 1:
		Score += 40
	case 2:
		Score += 100
	case 3:
		Score += 300
	default:
		break
	}

	Score += moveDownPoints
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
	} else {
		rl.PlaySound(RotateSound)
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
	rowsCleared := grid.ClearFullRows()
	if rowsCleared > 0 {
		rl.PlaySound(ClearSound)
	}
	UpdateScore(rowsCleared, 0)
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
	Score = 0
}

func HandleInput() {
	key := rl.GetKeyPressed()
	switch key {
	case rl.KeyLeft:
		moveLeft()
	case rl.KeyRight:
		moveRight()
	case rl.KeyDown:
		moveDown()
		UpdateScore(0, 1)
	case rl.KeySpace:
		unRotate()
	case rl.KeyUp:
		rotate()
	}
}
