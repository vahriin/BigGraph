package graph

import (
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/csv"
)

func ProcessAlCSV(csvChan chan<- csv.CSVWriter, wg *sync.WaitGroup, al types.AdjList) {
	for node, incidentNodes := range al.AdjacencyList {
		csvChan <- csv.ALLine{VertexID: node, IncidentVertexesID: incidentNodes}
	}

	close(csvChan)
	wg.Done()
}
