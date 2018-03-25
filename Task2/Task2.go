package main

import (
	"fmt"

	"github.com/vahriin/BigGraph/Task2/algo"
	"github.com/vahriin/BigGraph/Task2/input"
	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
)

func main() {
	dp := make(chan []uint64, 10)
	go csv.ReadPointsID(dp, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/input/destination_points.csv")

	p := make(chan coordinates.GeneralCoords)
	go input.ReadPoint(p, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/input/point.xml")

	startPoint := <-p

	destinationPoints := make(map[uint64]struct{})
	for pointsID := range dp {
		destinationPoints[pointsID[0]] = struct{}{}
	}

	al := model.NewAdjList("/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/input/adjacency_list.csv", "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/input/nodes_list.csv")

	nearestPointID := algo.Search(al, startPoint)

	al.Nodes[0] = startPoint
	al.AdjacencyList[0] = []uint64{nearestPointID}

	pathChan := make(chan algo.Path)

	go algo.Dijkstra(pathChan, destinationPoints, 0, al)

	fmt.Printf("%v\n", <-pathChan)

}
