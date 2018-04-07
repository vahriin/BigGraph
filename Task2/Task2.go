package main

import (
	"flag"
	"os"
	"runtime"
	"sync"

	"github.com/vahriin/BigGraph/Task2/algorithm"
	"github.com/vahriin/BigGraph/Task2/input"
	"github.com/vahriin/BigGraph/Task2/output"
	"github.com/vahriin/BigGraph/lib/coordinates"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
)

func pf() (bool, bool) {
	hell := flag.Bool("through_the_gates_of_hell", false, "")
	test := flag.Bool("t", false, "no output file")
	flag.Parse()
	return *hell, *test
}

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	broden, test := pf()

	oneStartPointChannel := make(chan coordinates.GeneralCoords)
	destinationPointsChannel := make(chan []uint64, 10)
	oneAdjacencyListChannel := make(chan model.AdjList)

	input.ParallelInput(oneStartPointChannel, destinationPointsChannel, oneAdjacencyListChannel, "input/Task2/")

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
	var wg sync.WaitGroup
	outMapChan := make(chan svg.SVGWriter, 10000)
	outCSVChan := make(chan csv.CSVWriter, len(destinationPoints))
	wg.Add(2)

	directory := "output/Task2/"
	rg := directory + "road_graph.svg"
	pw := directory + "pathways.csv"

	if test {
		directory = ""
		rg = "/dev/null"
		pw = "/dev/null"
	}
	os.MkdirAll(directory, 0777)
	go svg.ParallelWrite(outMapChan, &wg, rg)
	go csv.ParallelWrite(outCSVChan, &wg, pw)

	// draw graph
	var dgwg sync.WaitGroup
	dgwg.Add(1)
	go output.DrawGraph(outMapChan, &dgwg, adjacencyList)

	pathChan := make(chan model.Path, 10)

	// algorithms
	if !broden {
		//go algorithm.Dijkstra(pathChan, destinationPoints, 0, adjacencyList)
		go algorithm.Levit(pathChan, destinationPoints, 0, adjacencyList)
		//go algorithm.Astar(pathChan, destinationPoints, 0, adjacencyList)
	} else {
		pathChan <- output.BrodenPath()
		close(pathChan)
	}

	dgwg.Wait()

	output.ProcessPath(outCSVChan, outMapChan, adjacencyList, pathChan)

	close(outCSVChan)

	outMapChan <- svg.Circle{Center: startPoint.Euclid, Color: "green", Radius: svg.PointAttentionRadius}
	close(outMapChan)

	wg.Wait()
}
