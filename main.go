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

	qLeanWins := 0

	ma := twma.NewTWMA(100 * time.Second)
	t := time.Now()

	for i := 0; i < 50000; i++ {
		t = t.Add(time.Second)
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
			if winner == nil {
				fmt.Println("drow")
				ma.Apply(twma.Item{
					Value: 0.0,
					Time:  t,
				})
			} else {
				if winner.Name() == "qLean" {
					p1.ApplyFromLast(1)
					ma.Apply(twma.Item{
						Value: 1.0,
						Time:  t,
					})
					qLeanWins++
				} else {
					p1.ApplyFromLast(-1)
					ma.Apply(twma.Item{
						Value: 0.0,
						Time:  t,
					})
				}
			}
		} else {
			log.Fatalln(err)
		}
		v, err := ma.Value()
		if err == nil && i%100 == 0 {
			fmt.Println("log:", "\t", i, "\t", v)
		}
	}
	fmt.Println(qLeanWins)
}
