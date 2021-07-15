package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TTRSQ/CircleCrossGame/domain/agent/randcp"
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/CircleCrossGame/domain/game/display/console"
	"github.com/TTRSQ/qLearnAgent/src/agent"
	"github.com/TTRSQ/twma"
)

func main() {
	p1 := agent.Get(constants.CIRCLE, "qLean")
	p2 := randcp.Get(constants.CROSS, "rand")

	// 適切ではないが100s窓のmaを1秒ずつ進めて勝率を出す
	ma := twma.NewTWMA(100 * time.Second)
	t := time.Now()

	for i := 0; i < 50000; i++ {
		t = t.Add(time.Second)
		// game毎に手順を初期化する
		p1.InitHist()
		manager, err := game.NewManager(
			&p1,
			p2,
			board.NewBoard(),
			console.Get(),
		)
		if err != nil {
			log.Fatalln(err)
		}
		err = manager.Play()

		if err == nil {
			winner := manager.Winner()
			isQLearnWin := false
			if winner == nil {
				fmt.Println("drow")
			} else {
				if winner.Name() == "qLean" {
					// 勝ったら報酬 +1
					p1.ApplyFromLast(1)
					isQLearnWin = true
				} else {
					// 負けたら報酬 -1
					p1.ApplyFromLast(-1)
				}
			}
			ma.Apply(twma.Item{
				// 勝ったときだけ1.0を加算
				Value: map[bool]float64{true: 1.0, false: 0.0}[isQLearnWin],
				Time:  t,
			})
		} else {
			log.Fatalln(err)
		}
		v, err := ma.Value()
		if err == nil && i%100 == 0 {
			fmt.Println("win_rate:", "\t", i, "\t", v)
		}
	}
}
