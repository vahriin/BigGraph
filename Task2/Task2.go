package main

import (
	"fmt"

	"github.com/vahriin/BigGraph/lib/model"
)

func main() {
	fmt.Printf("%d\n", len(model.NewAdjList("input/adjacency_list.csv", "input/nodes_list.csv").Nodes))
}
