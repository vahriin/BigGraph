package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/vahriin/BigGraph/Task2/algorithm"
	"github.com/vahriin/BigGraph/Task2/input"
	"github.com/vahriin/BigGraph/Task2/output"
	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
)

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	start := time.Now()

	oneStartPointChannel := make(chan coordinates.GeneralCoords)
	destinationPointsChannel := make(chan []uint64, 10)
	oneAdjacencyListChannel := make(chan model.AdjList)

	input.ParallelInput(oneStartPointChannel, destinationPointsChannel, oneAdjacencyListChannel, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/input/Task2/")

	// getting of general variables
	startPoint := <-oneStartPointChannel
	destinationPoints := make(map[uint64]struct{})
	for pointsID := range destinationPointsChannel {
		destinationPoints[pointsID[0]] = struct{}{}
	}
	adjacencyList := <-oneAdjacencyListChannel

	// modify adjacency list
	nearestPointID := algorithm.Search(adjacencyList, startPoint)
	adjacencyList.Nodes[0] = startPoint
	adjacencyList.AdjacencyList[0] = []uint64{nearestPointID}

	// file-writer goroutines
	os.MkdirAll("output/Task2", 0777)
	var wg sync.WaitGroup
	outMapChan := make(chan svg.SVGWriter, 10000)
	outCSVChan := make(chan csv.CSVWriter, len(destinationPoints))
	wg.Add(2)
	go svg.ParallelWrite(outMapChan, &wg, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/output/Task2/road_graph.svg")
	go csv.ParallelWrite(outCSVChan, &wg, "/home/vahriin/Projects/GO/src/github.com/vahriin/BigGraph/output/Task2/pathways.csv")

	// draw graph
	wg.Add(1)
	go output.DrawGraph(outMapChan, &wg, adjacencyList)

	// algorithms
	pathChan := make(chan model.Path)
	go algorithm.Dijkstra(pathChan, destinationPoints, 0, adjacencyList)

	//
	output.ProcessPath(outCSVChan, outMapChan, adjacencyList, pathChan)
	close(outCSVChan)

	outMapChan <- svg.Circle{Center: startPoint.Euclid, Color: "green", Radius: svg.PointAttentionRadius}
	close(outMapChan)

	wg.Wait()

	log.Println("End; time spent:", time.Since(start))
}
