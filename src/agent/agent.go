package agent

import (
	"fmt"
	"math/rand"

	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/qLearnAgent/src/qFunc"
	"github.com/TTRSQ/qLearnAgent/src/status"
)

// ユーザー(console)
type qLearning struct {
	symbol constants.Symbol
	hist   status.History
	name   string
	qf     qFunc.Func
}

func Get(symbol constants.Symbol, name string) qLearning {
	return qLearning{
		symbol: symbol,
		name:   name,
		hist:   status.NewHistory(),
		qf:     qFunc.NewQFunc(),
	}
}

func (q *qLearning) InitHist() {
	q.hist = status.NewHistory()
}

// 後ろから割引しつつ報酬を伝搬する
func (q *qLearning) ApplyFromLast(r float64) {
	lastIdx := len(q.hist.Nodes) - 1
	rGamma := r
	for i := lastIdx; i >= 0; i -= 1 {
		q.qf.Apply(
			q.hist.Nodes[i].Status(),
			q.hist.Nodes[i].Action(),
			rGamma,
		)
		rGamma *= 0.1
	}
}

func (q *qLearning) useGreedy() bool {
	// 1/5の確率で適当に打つ
	return rand.Int()%5 == 0
}

func (q *qLearning) NextAction(board board.Board) (*action.Item, error) {
	canPutPoints := board.CanPutPoints()
	selection := rand.Int() % len(canPutPoints)
	// useGreedy == true の時はrandで選んだ手をそのまま返却する
	retAct, err := action.NewItem(canPutPoints[selection][0], canPutPoints[selection][1], q.symbol)

	act := []int{}
	point := -100000.0
	stPos := qFunc.CalcStatus(board)
	if !q.useGreedy() {
		for i := range canPutPoints {
			actPos := qFunc.CalcAct(canPutPoints[i][0], canPutPoints[i][1])
			pp := q.qf.Value(stPos, actPos)
			p := 0.0
			if pp != nil {
				p = *pp
			}
			if p > point {
				point = p
				act = canPutPoints[i]
			}
		}
		retAct, err = action.NewItem(act[0], act[1], q.symbol)
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(q.name + " placed " + fmt.Sprintf(
		"(%d %d)",
		retAct.X(),
		retAct.Y(),
	))
	q.hist.Append(status.NewNode(stPos, qFunc.CalcAct(retAct.X(), retAct.Y())))

	return retAct, nil
}

func (q *qLearning) Symbol() constants.Symbol {
	return q.symbol
}

func (q *qLearning) Name() string {
	return q.name
}
