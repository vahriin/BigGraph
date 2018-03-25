package main

import (
	"sync"

	"github.com/vahriin/BigGraph/Task2/algo"
	"github.com/vahriin/BigGraph/Task2/input"
	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
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

	outMapChan := make(chan svg.SVGWriter, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go svg.ParallelWrite(outMapChan, &wg, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/output/road_graph2.svg")

	for path := range pathChan {
		polyline := svg.Polyline{Points: make([]coordinates.EuclidCoords, 0, len(path.Points)), Width: svg.WidthPolyline, Color: svg.PolylineColor}

		for _, id := range path.Points {
			polyline.Points = append(polyline.Points, al.Nodes[id].Euclid)
		}

		outMapChan <- polyline
	}
	close(outMapChan)
	wg.Wait()
}
