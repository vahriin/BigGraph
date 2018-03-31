package algorithm

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/model"
)

func Levit(out chan<- model.Path, endpoints map[uint64]struct{}, start uint64, al model.AdjList) {
	distances, processingStatus, previousVertices := lPrepare(al) // nonProcessed is set M_2
	var mainQueue queue                                           // M'_1
	var priorityQueue queue                                       // M''_1

	defer close(out)

	distances[start] = 0
	mainQueue.push(start)
	processingStatus[start] = 1

	for !(mainQueue.empty() && priorityQueue.empty()) {
		var currentVertex uint64
		if priorityQueue.empty() {
			currentVertex = mainQueue.pop()
		} else {
			currentVertex = priorityQueue.pop()
		}

		for _, incidentVertex := range al.AdjacencyList[currentVertex] {
			ps, ok := processingStatus[incidentVertex]
			if !ok {
				panic("This is baaaaaad")
			}

			switch ps {
			case 2:
				mainQueue.push(incidentVertex)
				processingStatus[incidentVertex] = 1
				newDist := coordinates.Distance(al.Nodes[currentVertex].Euclid, al.Nodes[incidentVertex].Euclid)
				if distances[incidentVertex] > distances[currentVertex]+newDist {
					distances[incidentVertex] = distances[currentVertex] + newDist
					previousVertices[incidentVertex] = currentVertex
				}

			case 1:
				newDist := coordinates.Distance(al.Nodes[currentVertex].Euclid, al.Nodes[incidentVertex].Euclid)
				if distances[incidentVertex] > distances[currentVertex]+newDist {
					distances[incidentVertex] = distances[currentVertex] + newDist
					previousVertices[incidentVertex] = currentVertex
				}
			case 0:
				newDist := coordinates.Distance(al.Nodes[currentVertex].Euclid, al.Nodes[incidentVertex].Euclid)
				if distances[incidentVertex] > distances[currentVertex]+newDist {
					priorityQueue.push(incidentVertex)
					processingStatus[incidentVertex] = 1
					distances[incidentVertex] = distances[currentVertex] + newDist
					previousVertices[incidentVertex] = currentVertex
				}
			}
		}

		processingStatus[currentVertex] = 0
	}

	for endpoint, _ := range endpoints {
		path := model.Path{Len: distances[endpoint], Points: make([]uint64, 0, 20)}
		tempID := endpoint

		for tempID != start {
			path.Points = append(path.Points, tempID)
			tempID = previousVertices[tempID]
		}

		path.Points = append(path.Points, tempID)
	}

	for len(endpoints) > 0 {
		minimum := math.MaxFloat64
		var minimumID uint64
		for id, _ := range endpoints {
			if distances[id] < minimum {
				minimum = distances[id]
				minimumID = id
			}
		}

		path := model.Path{Len: distances[minimumID], Points: make([]uint64, 0, 20)}
		tempID := minimumID

		for tempID != start {
			path.Points = append(path.Points, tempID)
			tempID = previousVertices[tempID]
		}

		path.Points = append(path.Points, tempID)

		out <- path
		delete(endpoints, minimumID)
	}

}

func lPrepare(al model.AdjList) (map[uint64]float64, map[uint64]uint8, map[uint64]uint64) {
	distances := make(map[uint64]float64, len(al.Nodes))
	processed := make(map[uint64]uint8, len(al.Nodes))

	var w sync.WaitGroup
	w.Add(2)

	go func() {
		for key := range al.Nodes {
			distances[key] = math.MaxFloat64
		}
		w.Done()
	}()

	go func() {
		for key := range al.Nodes {
			processed[key] = 2
		}
		w.Done()
	}()

	previousVertices := make(map[uint64]uint64)

	w.Wait()

	return distances, processed, previousVertices
}
