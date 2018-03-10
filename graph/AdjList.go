package graph

import (
	"os"

	"github.com/vahriin/BigGraph/csv"
	"github.com/vahriin/BigGraph/types"
)

func AdjList(area *types.Area, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvAL := csv.NewCSV(file)

	defer csvAL.Close()

	al := make(map[uint64][]uint64)

	for id := range area.Points {
		al[id] = make([]uint64, 0, 4)
	}

	for _, edge := range area.Highways {
		for i := 1; i < len(edge.NodesID); i++ {
			al[edge.NodesID[i]] = append(al[edge.NodesID[i]], edge.NodesID[i-1])
		}
		for i := 0; i < len(edge.NodesID)-1; i++ {
			al[edge.NodesID[i]] = append(al[edge.NodesID[i]], edge.NodesID[i+1])
		}
	}

	for id, nodes := range al {
		csvAL.ALLine(id, nodes)
	}
}
