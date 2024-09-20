package game

import (
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/utils"
)

func GameOverLoad() {
	rl.BeginDrawing()
	rl.ClearBackground(utils.DarkBlue)
	rl.DrawTextEx(
		Font,
		"Game Over! :(",
		rl.Vector2{X: 140, Y: 30},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	rl.DrawTextEx(
		Font,
		"Score: "+strconv.Itoa(Score),
		rl.Vector2{X: 150, Y: 150},
		38,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
	if rl.IsKeyPressed(rl.KeyEnter) {
		InGame = false
		GameOver = false
		Score = 0
		Reset()
	}
}
