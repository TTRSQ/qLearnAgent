# qLearnAgent
qLearnAgent for CircleCrossGame
https://github.com/TTRSQ/CircleCrossGame のAgentとして、QLearningを用いたものを実装してみました。

https://github.com/TTRSQ/qLearnAgent/blob/main/src/qFunc/qFunc.go
が本体になるかと思います。
実装内でのQ関数の実態はハッシュマップです。
`map[int]map[int]float64{}`
利用側は返却値がnullの場合は未評価な座標として扱い、0.0を評価値と考えて利用しています。
