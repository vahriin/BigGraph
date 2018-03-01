package graph

import (
	"os"

	"github.com/vahriin/BigGraph/csv"
	"github.com/vahriin/BigGraph/types"
)

func CSVNodeList(area *types.Area, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvTable := csv.NewCSV(file)

	defer csvTable.Close()

	for id, coords := range area.Points {
		csvTable.NLLine(id, &coords)
	}
}
