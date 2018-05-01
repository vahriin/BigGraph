package aco

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/algorithm"
	"github.com/vahriin/BigGraph/lib/model"
)

func AntColonyOptimization(out chan<- model.Path, points map[uint64]struct{}, start uint64, al model.AdjList) {
	defer close(out)

	pathMatrix := PreparePM(start, points, al)

	ants := MakeAnts(&pathMatrix, start, ANTS_AMOUNT)

	bestPath := model.Path{
		Points: make([]uint64, 0),
		Len:    math.MaxFloat64,
	}

	iterations := ITERATIONS

	pathCh := make(chan model.Path, len(ants))
	var wg sync.WaitGroup
	for iterations > 0 {
		for _, ant := range ants {
			wg.Add(1)
			ant := ant
			go ant.MakePath(pathCh, pathMatrix, &wg)
		}

		for j := 0; j < len(ants); j++ {
			path := <-pathCh

			if path.Len < bestPath.Len {
				bestPath = path
			}
		}

		VapePheromones(&pathMatrix)

		iterations--
		wg.Wait()
	}

	for i := 1; i < len(bestPath.Points); i++ {
		out <- pathMatrix.Path(bestPath.Points[i-1], bestPath.Points[i])
	}

}

func PreparePM(start uint64, travelPoints map[uint64]struct{}, al model.AdjList) PathMatrix {
	pm := NewPathMatrix(start, travelPoints)

	var wg sync.WaitGroup
	citiesCopy := make(map[uint64]struct{})
	for id := range pm.cities {
		citiesCopy[id] = struct{}{}
	}

	for start := range pm.cities {
		delete(citiesCopy, start)

		localTravelPoints := make(map[uint64]struct{})
		for id := range citiesCopy {
			localTravelPoints[id] = struct{}{}
		}

		if len(localTravelPoints) == 0 {
			break
		}

		wg.Add(1)
		go func(start uint64, travelPoints map[uint64]struct{}) {

			outCh := make(chan model.Path, len(travelPoints))
			algorithm.Levit(outCh, travelPoints, start, al)

			for path := range outCh {
				pm.SetPath(path.Start(), path.End(), path)
				pm.SetPheromone(path.Start(), path.End(), 1/float64(len(pm.cities)))
			}
			wg.Done()
		}(start, localTravelPoints)
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

			newPheromone := pm.Pheromone(i, j) * (1 - PH_VAPE)
			if newPheromone < 1/float64(len(pm.cities)) {
				newPheromone = 1 / float64(len(pm.cities))
			}
			pm.SetPheromone(i, j, newPheromone)
		}
	}
}
