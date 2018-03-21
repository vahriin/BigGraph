package graph

import (
	"os"
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/csv"
)

// CSVAdjList write adjacency list to csv file with 'filename' name.
func CSVAdjList(al types.AdjList, filename string, wg *sync.WaitGroup) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvTable := csv.NewCSV(file)

	defer func() {
		csvTable.Close()
		wg.Done()
	}()

	for node, incidentNodes := range al.AL {
		csvTable.ALLine(node, incidentNodes)
	}
}
