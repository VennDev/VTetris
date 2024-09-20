package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/game"
	"github.com/venndev/VTetris/src/venndev/utils"
)

const (
	Title        = "VTetris"
	ScreenWidth  = 500
	ScreenHeight = 650
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, Title)

	// Set our game to run at 60 frames-per-second
	rl.SetTargetFPS(60)

	game.Init()

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(utils.DarkBlue)
		game.Draw()
		game.HandleInput()
		game.Update()
		rl.EndDrawing()
	}

	// De-Initialization
	rl.CloseWindow()
}
