package aco

import (
	"fmt"
	"math"

	"github.com/vahriin/BigGraph/lib/model"
)

func AntColonyOptimization(out chan<- model.Path, points map[uint64]struct{}, start uint64, al model.AdjList) {
	defer close(out)

	pathMatrix := PreparePM(start, points, al)

	ants := MakeAnts(&pathMatrix, start, 20)

	bestPath := model.Path{
		Points: make([]uint64, 0),
		Len:    math.MaxFloat64,
	}

	iterations := 1

	pathCh := make(chan model.Path, len(ants))
	for iterations > 0 {
		//fmt.Println(iterations)
		for _, ant := range ants {
			// TODO: add goroutine
			ant.MakePath(pathCh, &pathMatrix)
		}

		for j := 0; j < len(ants); j++ {
			path := <-pathCh

			if path.Len < bestPath.Len {
				bestPath = path
			}
		}

		VapePheromones(&pathMatrix)

		iterations--
	}

	fmt.Println(bestPath)

	for i := 1; i < len(bestPath.Points); i++ {
		fmt.Println(pathMatrix.Path(bestPath.Points[i-1], bestPath.Points[i]))
		out <- pathMatrix.Path(bestPath.Points[i-1], bestPath.Points[i])
	}

}
