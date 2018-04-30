package aco

import (
	"sync"

	"github.com/vahriin/BigGraph/lib/algorithm"
	"github.com/vahriin/BigGraph/lib/model"
)

func PreparePM(start uint64, travelPoints map[uint64]struct{}, al model.AdjList) PathMatrix {
	pm := NewPathMatrix(start, travelPoints)

	var wg sync.WaitGroup
	for start := range pm.cities {
		wg.Add(1)

		go func(start uint64) {
			localTravelPoints := make(map[uint64]struct{}, len(travelPoints)-1)
			for id := range travelPoints {
				if id == start {
					continue
				}
				localTravelPoints[id] = struct{}{}
			}

			outCh := make(chan model.Path, len(localTravelPoints))
			go algorithm.Levit(outCh, localTravelPoints, start, al)

			for path := range outCh {
				// TODO: catch unreachable points here

				pm.SetPath(start, path.Points[0], path)
			}

			wg.Done()
		}(start)
	}

	for i := range pm.cities {
		for j := range pm.cities {
			if i == j {
				continue
			}

			pm.SetPheromone(i, j, phMin)
		}
	}

	wg.Wait()
	return pm
}

func VapePheromones(pm *PathMatrix) {
	for i := range pm.cities {
		for j := range pm.cities {
			if i == j {
				continue
			}

			newPheromone := pm.Pheromone(i, j) * (1 - phVape)
			if newPheromone < phMin {
				newPheromone = phMin
			}
			pm.SetPheromone(i, j, newPheromone)
		}
	}
}
