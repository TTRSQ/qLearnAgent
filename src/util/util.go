package util

import "github.com/TTRSQ/CircleCrossGame/domain/game/board"

func CalcStatus(board board.Board) int {
	stamp := [][]int{
		{6561, 2187, 729},
		{243, 81, 27},
		{9, 3, 1},
	}
	sum := 0
	for i := range board.Status() {
		for j := range board.Status()[i] {
			symbolValue := 0
			if board.Status()[i][j] != nil {
				symbolValue = int(*(board.Status()[i][j]))
				sum += symbolValue * stamp[i][j]
			}
		}
	}
	return sum
}

func CalcAct(x, y int) int {
	stamp := [][]int{
		{256, 128, 64},
		{32, 16, 8},
		{4, 2, 1},
	}
	return stamp[y][x]
}
