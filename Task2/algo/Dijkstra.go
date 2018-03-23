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

	currentVertexID := start
	distances[currentVertexID] = 0

	for foo := 0; foo < len(al.Nodes); foo++ {
		if len(endpoints) == 0 || (distances[currentVertexID]-math.MaxFloat64) < 1E-10 {
			break
		}

		minimum := math.MaxFloat64 // minimum distance to incident point
		var minimumVertexID uint64 // ID of incident point at which the distance is minimum

		for _, incidentVertexID := range al.AdjacencyList[currentVertexID] {
			d := coordinates.Distance(al.Nodes[currentVertexID].Euclid, al.Nodes[incidentVertexID].Euclid)
			if d < distances[currentVertexID]+distances[incidentVertexID] && !processed[incidentVertexID] {
				distances[incidentVertexID] = distances[currentVertexID] + d // relaxation
				previousVertices[incidentVertexID] = currentVertexID
			}
			if minimum > d {
				minimum = d
				minimumVertexID = incidentVertexID
			}
		}

		// make path to point in endpoints
		if _, ok := endpoints[currentVertexID]; ok {
			path := make(Path, 20)
			tempID := currentVertexID

			for tempID != start {
				path = append(path, tempID)
				tempID = previousVertices[tempID]
			}

			out <- path

			delete(endpoints, currentVertexID)
		}

		// point is processed
		processed[currentVertexID] = true
		currentVertexID = minimumVertexID
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
