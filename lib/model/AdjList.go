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
	alChan := make(chan map[uint64][]uint64)
	nlChan := make(chan map[uint64]coordinates.GeneralCoords)

	go csv.ReadNodeList(nlChan, nodesList)
	go csv.ReadAdjacencyList(alChan, adjacencyList)

	var al AdjList
	al.AdjacencyList = <-alChan
	al.Nodes = <-nlChan

	return al
}

func (al *AdjList) AddPoint(coords coordinates.GeneralCoords) {
	pdChan := make(chan pointDist, 1000)

	go func() {
		minimal := pointDist{point: 0, dist: math.MaxFloat64}

		for pd := range pdChan {
			if pd.dist < minimal.dist {
				minimal = pd
			}
		}

		al.Nodes[0] = coords
		al.AdjacencyList[0] = []uint64{minimal.point}
	}()

	var wg sync.WaitGroup
	for point, coord := range al.Nodes {
		wg.Add(1)
		go func(point uint64, coord coordinates.GeneralCoords) {
			pdChan <- pointDist{point: point, dist: coordinates.Distance(coord.Euclid, coords.Euclid)}
			wg.Done()
		}(point, coord)
	}
	wg.Wait()
	close(pdChan)
}
