package graph

import "errors"

type AdjacencyMatrix struct {
	nodes map[uint]uint
	matrix []float64
}

func (am AdjacencyMatrix) Distance(from, to uint) (float64, error) {
	indexfrom, ok1 := am.nodes[from]
	indexto, ok2 := am.nodes[to]
	if ok1 && ok2 {
		return am.matrix[uint(len(am.nodes)) * (indexfrom) + indexto], nil
	}
	return 0.0, errors.New("no id in matrix")
}
