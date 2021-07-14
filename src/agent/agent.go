package agent

import (
	"fmt"
	"math/rand"

	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/qLearnAgent/src/qFunc"
	"github.com/TTRSQ/qLearnAgent/src/status"
	"github.com/TTRSQ/qLearnAgent/src/util"
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

func (q *qLearning) NextAction(board board.Board) (*action.Item, error) {
	st := util.CalcStatus(board)
	acts := board.CanPutPoints()
	point := -100000.0
	act := []int{}
	selection := rand.Int() % len(acts)
	retAct, err := action.NewItem(acts[selection][0], acts[selection][1], q.symbol)
	// 1/5の確率で適当に打つ
	if rand.Int()%5 != 0 {
		for i := range acts {
			actVal := util.CalcAct(acts[i][0], acts[i][1])
			pp := q.qf.Value(st, actVal)
			p := 0.0
			if pp != nil {
				p = *pp
			}
			if p > point {
				point = p
				act = acts[i]
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
	q.hist.Append(status.NewNode(st, util.CalcAct(retAct.X(), retAct.Y())))

	return retAct, nil
}

func (q *qLearning) Symbol() constants.Symbol {
	return q.symbol
}

func (q *qLearning) Name() string {
	return q.name
}