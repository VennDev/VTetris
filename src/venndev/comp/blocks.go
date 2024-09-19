package comp

import "github.com/venndev/VTetris/src/venndev/utils/math"

func LBlock() *Block {
	return NewBlock(1, map[int]math.Position{
		0: {0, 0},
		1: {0, 1},
		2: {0, 2},
		3: {1, 2},
	}, 30, 4)
}

func JBlock() *Block {
	return NewBlock(2, map[int]math.Position{
		0: {0, 2},
		1: {1, 2},
		2: {2, 2},
		3: {2, 1},
	}, 30, 4)
}

func IBlock() *Block {
	return NewBlock(3, map[int]math.Position{
		0: {0, 0},
		1: {0, 1},
		2: {0, 2},
		3: {0, 3},
	}, 30, 2)
}

func OBlock() *Block {
	return NewBlock(4, map[int]math.Position{
		0: {0, 0},
		1: {0, 1},
		2: {1, 0},
		3: {1, 1},
	}, 30, 1)
}

func SBlock() *Block {
	return NewBlock(5, map[int]math.Position{
		0: {0, 1},
		1: {0, 2},
		2: {1, 0},
		3: {1, 1},
	}, 30, 2)
}

func ZBlock() *Block {
	return NewBlock(6, map[int]math.Position{
		0: {0, 0},
		1: {0, 1},
		2: {1, 1},
		3: {1, 2},
	}, 30, 2)
}

func TBlock() *Block {
	return NewBlock(7, map[int]math.Position{
		0: {0, 1},
		1: {1, 0},
		2: {1, 1},
		3: {1, 2},
	}, 30, 4)
}
