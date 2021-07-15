package qFunc

import "github.com/TTRSQ/CircleCrossGame/domain/game/board"

// 状態 * 行動 の空間すべての値を保存している。
type Func struct {
	space map[int]map[int]float64 // 前者が状態、後者が行動
	eta   float64                 // 学習率
}

func NewQFunc() Func {
	return Func{
		space: map[int]map[int]float64{},
		eta:   0.1,
	}
}

func (f *Func) Apply(st, act int, r float64) error {
	delta := f.eta * r
	_, exists := f.space[st]
	if !exists {
		f.space[st] = map[int]float64{
			act: delta,
		}
		return nil
	}
	_, exists = f.space[st][act]
	if !exists {
		f.space[st][act] = delta
		return nil
	}

	f.space[st][act] += delta

	return nil
}

func (f *Func) Value(st, act int) *float64 {
	acts, exists := f.space[st]
	if !exists {
		return nil
	}
	val, exists := acts[act]
	if !exists {
		return nil
	}

	return &val
}

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
