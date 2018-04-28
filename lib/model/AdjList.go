package model

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/csv"
)

type AdjList struct {
	AdjacencyList map[uint64][]uint64
	Nodes         map[uint64]coordinates.GeneralCoords
}

func NewAdjList(adjacencyList, nodesList string) AdjList {
	return AdjList{AdjacencyList: csv.ReadAdjacencyList(adjacencyList), Nodes: csv.ReadNodeList(nodesList)}
}

func (al AdjList) AddPoint(coords coordinates.GeneralCoords) {
	pdChan := make(chan pointDist, 1000)

	var pSync sync.WaitGroup
	pSync.Add(1)
	go func() {
		minimal := pointDist{point: 0, dist: math.MaxFloat64}

		for pd := range pdChan {
			if pd.dist < minimal.dist {
				minimal = pd
			}
		}

		al.Nodes[0] = coords
		al.AdjacencyList[0] = []uint64{minimal.point}
		al.AdjacencyList[minimal.point] = append(al.AdjacencyList[minimal.point], 0)
		pSync.Done()
	}()

	var distSync sync.WaitGroup
	for point, coord := range al.Nodes {
		distSync.Add(1)
		go func(point uint64, coord coordinates.GeneralCoords) {
			pdChan <- pointDist{point: point, dist: coordinates.Distance(coord.Euclid, coords.Euclid)}
			distSync.Done()
		}(point, coord)
	}
	distSync.Wait()
	close(pdChan)
	pSync.Wait()
}
