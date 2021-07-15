package status

type node struct {
	st  int
	act int
}

func (n *node) Status() int {
	return n.st
}

func (n *node) Action() int {
	return n.act
}

func NewNode(st, act int) node {
	return node{
		st:  st,
		act: act,
	}
}

// History こういう状態に対してこのような行動をしたよという記録を積む
type History struct {
	Nodes []node
}

func (h *History) Append(n node) {
	h.Nodes = append(h.Nodes, n)
}

func NewHistory() History {
	return History{
		Nodes: []node{},
	}
}
