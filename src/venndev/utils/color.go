package utils

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	White       = rl.Color{1, 1, 1, 1}
	Black       = rl.Color{0, 0, 0, 1}
	Red         = rl.Color{1, 0, 0, 1}
	Blue        = rl.Color{0, 0, 1, 1}
	Yellow      = rl.Color{1, 1, 0, 1}
	Cyan        = rl.Color{0, 1, 1, 1}
	Magenta     = rl.Color{1, 0, 1, 1}
	Transparent = rl.Color{0, 0, 0, 0}
	DarkBlue    = rl.Color{44, 44, 127, 255}
)

func GetCellColors() []rl.Color {
	return []rl.Color{
		rl.White,
		rl.Red,
		rl.Blue,
		rl.Green,
		rl.Yellow,
		rl.Purple,
		rl.Orange,
		rl.Brown,
	}
}
