package qFunc

// 状態 * 行動 の空間すべての値を保存している。
type Func struct {
	// 前者が状態、後者が空間
	space map[int]map[int]float64
	eta   float64
}

func NewQFunc() Func {
	return Func{
		space: map[int]map[int]float64{},
		eta:   0.1,
	}
}

func (f *Func) Apply(st, act int, r float64) error {
	_, exists := f.space[st]
	if !exists {
		f.space[st] = map[int]float64{
			act: r,
		}
		return nil
	}
	_, exists = f.space[st][act]
	if !exists {
		f.space[st][act] = r
		return nil
	}

	f.space[st][act] += f.eta * r

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
