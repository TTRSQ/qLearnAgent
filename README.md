# qLearnAgent
qLearnAgent for CircleCrossGame
https://github.com/TTRSQ/CircleCrossGame のAgentとして、QLearningを用いたものを実装してみました。

https://github.com/TTRSQ/qLearnAgent/blob/main/src/qFunc/qFunc.go
が本体になるかと思います。

## 評価値の保存先
実装内でのQ関数の実態はハッシュマップです。

### `map[int]map[int]float64{}`
利用側は返却値がnullの場合は未評価な座標として扱い、0.0を評価値と考えて利用しています。

## 実行
```
$ go run main.go | grep learn_data
learn_data: 	 game_count 	 win_rate 	 epsilon
learn_data: 	 100 	 0.595 	 1.00
learn_data: 	 200 	 0.685 	 1.00
learn_data: 	 300 	 0.675 	 0.99
learn_data: 	 400 	 0.645 	 0.99
learn_data: 	 500 	 0.625 	 0.99
learn_data: 	 600 	 0.64 	 0.99
learn_data: 	 700 	 0.605 	 0.99
...
learn_data: 	 49300 	 0.99 	 0.01
learn_data: 	 49400 	 1 	 0.01
learn_data: 	 49500 	 0.98 	 0.01
learn_data: 	 49600 	 0.99 	 0.01
learn_data: 	 49700 	 0.99 	 0.01
learn_data: 	 49800 	 1 	 0.00
learn_data: 	 49900 	 0.99 	 0.00
```
※ **epsilon** は手をランダムに打つ割合