package model

import "github.com/vahriin/BigGraph/lib/coordinates"
import "github.com/vahriin/BigGraph/lib/csv"

type AdjList struct {
	AdjacencyList map[uint64][]uint64
	Nodes         map[uint64]coordinates.GeneralCoords
}

func NewAdjList(adjacencyList, nodesList string) AdjList {
	alChan := make(chan map[uint64][]uint64)
	nlChan := make(chan map[uint64]coordinates.GeneralCoords)

	go csv.ReadNodeList(nlChan, nodesList)
	go csv.ReadAdjacencyList(alChan, adjacencyList)

	var al AdjList
	al.AdjacencyList = <-alChan
	al.Nodes = <-nlChan

	return al
}
