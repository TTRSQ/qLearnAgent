package util

import "github.com/TTRSQ/CircleCrossGame/domain/game/board"

// Q関数における状態値(空間内での座標)を計算して返す
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

// Q関数における手の値(空間内での座標)を計算して返す
func CalcAct(x, y int) int {
	stamp := [][]int{
		{256, 128, 64},
		{32, 16, 8},
		{4, 2, 1},
	}
	return stamp[y][x]
}
