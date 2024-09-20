package comp

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/venndev/VTetris/src/venndev/utils"
)

type Grid struct {
	numRows  int // number of rows
	numCols  int // number of columns
	cellSize int // size of each cell
	Grid     [][]int
}

func (g *Grid) Print() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			print(g.Grid[i][j], " ")
		}
		println()
	}
}

func (g *Grid) Draw() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			color := utils.GetCellColors()[g.Grid[i][j]]
			rl.DrawRectangle(
				int32(j*g.cellSize+11),
				int32(i*g.cellSize+11),
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

func (g *Grid) IsCellEmpty(row, col int) bool {
	return g.Grid[row][col] == 0
}

func (g *Grid) IsRowFull(row int) bool {
	for j := 0; j < g.numCols; j++ {
		if g.Grid[row][j] == 0 {
			return false
		}
	}
	return true
}

func (g *Grid) ClearRow(row int) {
	for i := 0; i < g.numCols; i++ {
		g.Grid[row][i] = 0
	}
}

func (g *Grid) MoveRowsDown(row, numRows int) {
	for i := 0; i < g.numCols; i++ {
		g.Grid[row+numRows][i] = g.Grid[row][i]
		g.Grid[row][i] = 0
	}
}

func (g *Grid) ClearFullRows() int {
	rowsCleared := 0
	for i := g.numRows - 1; i >= 0; i-- {
		if g.IsRowFull(i) {
			g.ClearRow(i)
			rowsCleared++
		} else if rowsCleared > 0 {
			g.MoveRowsDown(i, rowsCleared)
		}
	}
	return rowsCleared
}

func (g *Grid) ClearAll() {
	for i := 0; i < g.numRows; i++ {
		for j := 0; j < g.numCols; j++ {
			g.Grid[i][j] = 0
		}
	}
}

func NewGrid(numRows, numCols, cellSize int) *Grid {
	return &Grid{
		numRows:  numRows,
		numCols:  numCols,
		cellSize: cellSize,
		Grid:     getGrid(numRows, numCols),
	}
}

func getGrid(numRows, numCols int) [][]int {
	grid := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		grid[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			grid[i][j] = 0
		}
	}
	return grid
}
