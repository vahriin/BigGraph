package graph

import (
	"bufio"
	"io"
	"os"

	"github.com/vahriin/BigGraph/csv"
)

type Vertex struct {
	NodeId uint64
	Distance float64
}

type AdjacencyList struct {
	List map[uint64][]Vertex
}

func (al AdjacencyList) Nodes() map[uint64][]uint64 {
	r := make(map[uint64][]uint64)
	for id, vertexes := range al.List {
		for _, vertex := range vertexes {
			r[id] = append(r[id], vertex.NodeId)
		}
	}
	return r
}

func (al AdjacencyList) ToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvTable := csv.NewCSV(file)

	defer csvTable.Close()

	for id, nodes := range al.Nodes() {
		csvTable.ALLine(id, nodes)
	}
}

/*func LoadAL(r io.Reader) AdjacencyList {
	bufreader := bufio.NewReader(r)

	var al AdjacencyList

	bufreader.ReadString('\n')
}*/