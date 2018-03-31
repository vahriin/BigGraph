package algorithm

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/model"
)

type pointDist struct {
	point uint64
	dist  float64
}

func Search(al model.AdjList, coords coordinates.GeneralCoords) (ID uint64) {
	pdChan := make(chan pointDist, 1000)
	idChan := make(chan uint64)

	go func() {
		minimal := pointDist{point: 0, dist: math.MaxFloat64}

		for pd := range pdChan {
			if pd.dist < minimal.dist {
				minimal = pd
			}
		}

		idChan <- minimal.point
		close(idChan)
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

	return <-idChan
}
