package graph

import (
	"os"
	"sync"

	"github.com/vahriin/BigGraph/csv"
	"github.com/vahriin/BigGraph/types"
)

// CSVNodeList write list of nodes to filename file
func CSVNodeList(al types.AdjList, filename string, wg *sync.WaitGroup) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvTable := csv.NewCSV(file)

	defer func() {
		csvTable.Close()
		wg.Done()
	}()

	for point, coords := range al.Nodes {
		csvTable.NLLine(point, &coords)
	}
}
