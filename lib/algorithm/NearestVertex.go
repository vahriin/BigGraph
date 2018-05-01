package algorithm

import (
	"github.com/vahriin/BigGraph/lib/model"
)

func NearestVertexAdd(out chan<- model.Path, points map[uint64]struct{}, start uint64, al model.AdjList) {
	defer close(out)

	var iterate = true
	startSave := start

	for len(points) != 0 {
		pathCh := make(chan model.Path, len(points))

		go Astar(pathCh, points, start, al)

		shortestPath := <-pathCh // Astar guarantees that the shortest path will be return first

		// throw out the garbage
		go func() {
			for range pathCh {
			}
		}()

		start = shortestPath.End()
		delete(points, start)

		if len(points) == 0 && iterate {
			iterate = false
			points[startSave] = struct{}{}
		}

		out <- shortestPath
	}
}
