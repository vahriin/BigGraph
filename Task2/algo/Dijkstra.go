package algo

import (
	"math"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"

	"github.com/vahriin/BigGraph/lib/model"
)

func Dijkstra(out chan<- Path, endpoints map[uint64]struct{}, start uint64, al model.AdjList) {
	distances, processed, previousVertices := dPrepare(al)

	defer close(out)

	distances[start] = 0

	for foo := 0; foo < len(al.Nodes); foo++ {
		minimumDistance := math.MaxFloat64 // minimum distance to incident point
		var currentVertexID uint64         // ID of incident point at which the distance is minimum

		for id, dist := range distances {
			if dist < minimumDistance && !processed[id] {
				minimumDistance = dist
				currentVertexID = id
			}
		}

		if math.Abs(minimumDistance-math.MaxFloat64) < 1E-6 { // unreachable point
			break
		}

		// make path to point in endpoints
		if _, ok := endpoints[currentVertexID]; ok {
			path := Path{Len: distances[currentVertexID], Points: make([]uint64, 0, 20)}
			tempID := currentVertexID

			for tempID != start {
				path.Points = append(path.Points, tempID)
				tempID = previousVertices[tempID]
			}

			path.Points = append(path.Points, tempID)

			out <- path

			delete(endpoints, currentVertexID)

			if len(endpoints) == 0 {
				break
			}
		}

		for _, incidentVertexID := range al.AdjacencyList[currentVertexID] {
			if !processed[incidentVertexID] {
				d := coordinates.Distance(al.Nodes[currentVertexID].Euclid, al.Nodes[incidentVertexID].Euclid)
				if distances[currentVertexID]+d < distances[incidentVertexID] {
					distances[incidentVertexID] = distances[currentVertexID] + d
					previousVertices[incidentVertexID] = currentVertexID
				}
			}
		}

		// point is processed
		processed[currentVertexID] = true
	}
}

func dPrepare(al model.AdjList) (map[uint64]float64, map[uint64]bool, map[uint64]uint64) {
	distances := make(map[uint64]float64, len(al.Nodes))

	var w sync.WaitGroup
	w.Add(2)

	go func() {
		for key := range al.Nodes {
			distances[key] = math.MaxFloat64
		}
		w.Done()
	}()

	processed := make(map[uint64]bool, len(al.Nodes))

	go func() {
		for key := range al.Nodes {
			processed[key] = false
		}
		w.Done()
	}()

	previousVertices := make(map[uint64]uint64, len(al.Nodes))

	w.Wait()

	return distances, processed, previousVertices
}
