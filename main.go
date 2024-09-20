package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/game"
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
		if !game.InGame {
			game.TitleScreen()
		} else {
			game.Load()
		}
		rl.EndDrawing()
	}

	// Close game
	game.Close()

	// De-Initialization
	rl.CloseWindow()
}
