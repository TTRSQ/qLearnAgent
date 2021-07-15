package qFunc

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
