package comp

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/utils"
)

type Grid struct {
	numRows  int // number of rows
	numCols  int // number of columns
	cellSize int // size of each cell
	grids    [][]int
}

func (g *Grid) Print() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			print(g.grids[i][j], " ")
		}
		println()
	}
}

func (g *Grid) Draw() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			color := utils.GetCellColors()[g.grids[i][j]]
			rl.DrawRectangle(
				int32(j*g.cellSize+1),
				int32(i*g.cellSize+1),
				int32(g.cellSize-1),
				int32(g.cellSize-1),
				color,
			)
		}
	}
}

func (g *Grid) IsCellOutSide(row, col int) bool {
	return row < 0 || row >= g.numRows || col < 0 || col >= g.numCols
}

func NewGrid(numRows, numCols, cellSize int) *Grid {
	return &Grid{
		numRows:  numRows,
		numCols:  numCols,
		cellSize: cellSize,
		grids:    getGrids(numRows, numCols),
	}
}

func getGrids(numRows, numCols int) [][]int {
	grid := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			grid[i][j] = 0
		}
	}
	return grid
}
