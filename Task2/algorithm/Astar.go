package algorithm

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"

	"github.com/vahriin/BigGraph/lib/model"
)

func Astar(out chan<- model.Path, endpoints map[uint64]struct{}, start uint64, al model.AdjList) {
	dm := aPrepare(al)
	dm.setDistance(start, 0)

	defer close(out)

	resultChan := make(chan model.Path, len(endpoints))

	for endID := range endpoints {
		go astar(resultChan, start, endID, al, dm)
	}

	paths := make([]model.Path, len(endpoints))
	for i := 0; i < len(endpoints); i++ {
		paths[i] = <-resultChan
	}
	close(resultChan)

	for i := 0; i < len(paths); i++ {
		minimum := math.MaxFloat64
		var minimumID int
		for index, p := range paths {
			if p.Len < minimum {
				minimum = p.Len
				minimumID = index
			}
		}

		// if point is unreachable
		if minimumID == 0 {
			return
		}

		out <- paths[minimumID]

		paths[minimumID].Len = math.MaxFloat64
	}
}

func astar(out chan<- model.Path, start uint64, end uint64, al model.AdjList, dm *distMap) {
	heuristicValues := make(map[uint64]float64) // contains only unvisited vertices
	previousVertices := make(map[uint64]uint64)
	processed := make(map[uint64]struct{})

	path := model.Path{Len: math.MaxFloat64, Points: make([]uint64, 0, 20)}

	defer func() {
		out <- path
	}()

	heuristicValues[start] = getHF()(al.Nodes[start].Euclid, al.Nodes[end].Euclid)

	for len(heuristicValues) != 0 {
		currentVertex := searchMinHeuristic(heuristicValues)

		if currentVertex == end {
			path.Len = dm.distance(currentVertex)
			tempID := currentVertex

			for tempID != start {
				path.Points = append(path.Points, tempID)
				tempID = previousVertices[tempID]
			}

			path.Points = append(path.Points, tempID)
			break
		}

		delete(heuristicValues, currentVertex)
		processed[currentVertex] = struct{}{}

		for _, incidentVertex := range al.AdjacencyList[currentVertex] {
			tentativeScore := dm.distance(currentVertex) +
				coordinates.Distance(al.Nodes[currentVertex].Euclid, al.Nodes[incidentVertex].Euclid)

			if _, ok := processed[incidentVertex]; ok && tentativeScore >= dm.distance(incidentVertex) {
				continue
			}
			if _, ok := processed[incidentVertex]; !ok || tentativeScore < dm.distance(incidentVertex) {
				previousVertices[incidentVertex] = currentVertex
				dm.setDistance(incidentVertex, tentativeScore)
				heuristicValues[incidentVertex] = dm.distance(incidentVertex) +
					getHF()(al.Nodes[incidentVertex].Euclid, al.Nodes[end].Euclid)
			}
		}

	}
}

func aPrepare(al model.AdjList) *distMap {
	distances := make(map[uint64]float64, len(al.Nodes))

	var w sync.WaitGroup
	w.Add(1)

	go func() {
		for key := range al.Nodes {
			distances[key] = math.MaxFloat64
		}
		w.Done()
	}()

	dm := new(distMap)
	dm.mutex = new(sync.RWMutex)
	w.Wait()

	dm.distances = distances

	return dm
}

func searchMinHeuristic(heuristic map[uint64]float64) uint64 {
	minimum := math.MaxFloat64
	var minimumID uint64
	for id, hv := range heuristic {
		if hv < minimum {
			minimum = hv
			minimumID = id
		}
	}
	return minimumID
}
