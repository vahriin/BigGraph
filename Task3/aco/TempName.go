package aco

import (
	"sync"

	"github.com/vahriin/BigGraph/lib/algorithm"
	"github.com/vahriin/BigGraph/lib/model"
)

func PreparePM(start uint64, travelPoints map[uint64]struct{}, al model.AdjList) PathMatrix {
	pm := NewPathMatrix(start, travelPoints)

	var wg sync.WaitGroup
	for start := range pm.index {
		wg.Add(1)
		go func(start uint64) {
			localTravelPoints := make(map[uint64]struct{}, len(travelPoints)-1)
			for id := range localTravelPoints {
				if id == start {
					continue
				}
				localTravelPoints[id] = struct{}{}
			}

			outCh := make(chan model.Path, len(localTravelPoints))
			go algorithm.Levit(outCh, localTravelPoints, start, al)

			for path := range outCh {
				pm.SetPath(start, path.Points[0], path)
			}

			wg.Done()
		}(start)
	}
	wg.Wait()
	return pm
}
