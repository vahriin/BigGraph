package aco

import (
	"fmt"
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/algorithm"
	"github.com/vahriin/BigGraph/lib/model"
)

func AntColonyOptimization(out chan<- model.Path, points map[uint64]struct{}, start uint64, al model.AdjList) {
	defer close(out)

	pathMatrix := PreparePM(start, points, al)

	ants := MakeAnts(&pathMatrix, start, 1000)

	bestPath := model.Path{
		Points: make([]uint64, 0),
		Len:    math.MaxFloat64,
	}

	iterations := 100

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

	fmt.Println(bestPath)

	for i := 1; i < len(bestPath.Points); i++ {
		//fmt.Println(pathMatrix.Path(bestPath.Points[i-1], bestPath.Points[i]))
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
				pm.SetPath(start, path.Points[0], path)
				pm.SetPheromone(start, path.Points[0], phMin)
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

			newPheromone := pm.Pheromone(i, j) * (1 - phVape)
			if newPheromone < phMin {
				newPheromone = phMin
			}
			pm.SetPheromone(i, j, newPheromone)
		}
	}
}
