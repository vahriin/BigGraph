package graph

import "github.com/vahriin/BigGraph/xmlparse"

type AdjacencyList struct {
	List map[xmlparse.Node][]xmlparse.Node
}

/*func NewAdjacencyList (area xmlparse.Area) AdjacencyList {
	var al AdjacencyList

	for _, edge := range area.Edges {

	}
}*/