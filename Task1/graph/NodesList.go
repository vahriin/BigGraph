package graph

import (
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/csv"
)

// ProcessNlCSV write list of nodes to filename file TODO: outdated
func ProcessNlCSV(csvChan chan<- csv.CSVWriter, wg *sync.WaitGroup, al types.AdjList) {
	for point, coords := range al.Nodes {
		csvChan <- csv.NLLine{VertexID: point, GeneralCoords: coords}
	}

	close(csvChan)
	wg.Done()
}
