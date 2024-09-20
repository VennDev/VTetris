package comp

import "github.com/venndev/VTetris/src/venndev/utils/math"

const (
	LBLOCK = 1
	JBLOCK = 2
	ZBLOCK = 3
	SBLOCK = 4
	TBLOCK = 5
	OBLOCK = 6
	IBLOCK = 7
)

func LBlock() *Block {
	return NewBlock(LBLOCK, map[int][]math.Position{
		0: {
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
		},
		1: {
			{0, 1},
			{1, 1},
			{1, 0},
			{1, -1},
		},
		2: {
			{0, 0},
			{0, 1},
			{1, 1},
			{2, 1},
		},
		3: {
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
		},
	}, 30, 0)
}

func JBlock() *Block {
	return NewBlock(JBLOCK, map[int][]math.Position{
		0: {
			{0, 1},
			{1, 1},
			{2, 1},
			{2, 0},
		},
		1: {
			{0, 0},
			{1, 0},
			{1, 1},
			{1, 2},
		},
		2: {
			{0, 1},
			{0, 0},
			{1, 0},
			{2, 0},
		},
		3: {
			{0, 2},
			{1, 2},
			{1, 1},
			{1, 0},
		},
	}, 30, 0)
}

func ZBlock() *Block {
	return NewBlock(ZBLOCK, map[int][]math.Position{
		0: {
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
		1: {
			{0, 1},
			{1, 1},
			{1, 0},
			{2, 0},
		},
		2: {
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
		3: {
			{0, 1},
			{1, 1},
			{1, 0},
			{2, 0},
		},
	}, 30, 0)
}

func SBlock() *Block {
	return NewBlock(SBLOCK, map[int][]math.Position{
		0: {
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 1},
		},
		1: {
			{0, 0},
			{1, 0},
			{1, 1},
			{2, 1},
		},
		2: {
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 1},
		},
		3: {
			{0, 0},
			{1, 0},
			{1, 1},
			{2, 1},
		},
	}, 30, 0)
}

func TBlock() *Block {
	return NewBlock(TBLOCK, map[int][]math.Position{
		0: {
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 1},
		},
		1: {
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 0},
		},
		2: {
			{0, 1},
			{1, 0},
			{1, 1},
			{1, 2},
		},
		3: {
			{0, 1},
			{1, 0},
			{1, 1},
			{2, 1},
		},
	}, 30, 0)
}

func OBlock() *Block {
	return NewBlock(OBLOCK, map[int][]math.Position{
		0: {
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
		1: {
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
		2: {
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
		3: {
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
	}, 30, 0)
}

func IBlock() *Block {
	return NewBlock(IBLOCK, map[int][]math.Position{
		0: {
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		1: {
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
		2: {
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		3: {
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
	}, 30, 0)
}
