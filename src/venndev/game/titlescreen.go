package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/utils"
)

func TitleScreen() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		InGame = true
	}

	rl.BeginDrawing()
	rl.ClearBackground(utils.DarkBlue)
	rl.DrawTextEx(
		Font,
		"VTetris",
		rl.Vector2{X: 150, Y: 150},
		64,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	rl.DrawTextEx(
		Font,
		"Press Enter to start",
		rl.Vector2{X: 50, Y: 300},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	rl.DrawTextEx(
		Font,
		"Press ESC to exit",
		rl.Vector2{X: 50, Y: 350},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
}
