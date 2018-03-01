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

	//AdjMatrix(area, al, "AM.csv")

	for _, edge := range area.Edges {
		for i := 1; i < len(edge.NodesId); i++ {
			al[edge.NodesId[i]] = append(al[edge.NodesId[i]], edge.NodesId[i-1])
		}
		for i := 0; i < len(edge.NodesId)-1; i++ {
			al[edge.NodesId[i]] = append(al[edge.NodesId[i]], edge.NodesId[i+1])
		}
	}

	for id, nodes := range al {
		csvAL.ALLine(id, nodes)
	}
}
