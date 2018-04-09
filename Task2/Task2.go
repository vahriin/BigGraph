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

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	broden, test := pf()

	oneStartPointChannel := make(chan coordinates.GeneralCoords)
	destinationPointsChannel := make(chan []uint64, 10)
	oneAdjacencyListChannel := make(chan model.AdjList)

	input.ParallelInput(oneStartPointChannel, destinationPointsChannel, oneAdjacencyListChannel, "input/Task2/")

	// getting of general variables
	startPoint, destinationPoints, adjacencyList := getInput(destinationPointsChannel, oneStartPointChannel, oneAdjacencyListChannel)

	// modify adjacency list
	modifyAL(&adjacencyList, startPoint)

	// file-writer goroutines
	outSVGChan, outCSVChan, wg := createInputChannel(test, 10000, len(destinationPoints))

	// draw graph
	var dgwg sync.WaitGroup
	dgwg.Add(1)
	go output.DrawGraph(outSVGChan, &dgwg, adjacencyList)

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

	output.ProcessPath(outCSVChan, outSVGChan, adjacencyList, pathChan)

	close(outCSVChan)

	outSVGChan <- svg.Circle{Center: startPoint.Euclid, Color: "green", Radius: svg.PointAttentionRadius}
	close(outSVGChan)

	wg.Wait()

}

func pf() (bool, bool) {
	hell := flag.Bool("through_the_gates_of_hell", false, "")
	test := flag.Bool("t", false, "no output file")
	flag.Parse()
	return *hell, *test
}

func getInput(dpCh <-chan []uint64, spCh <-chan coordinates.GeneralCoords, alCh <-chan model.AdjList) (coordinates.GeneralCoords, map[uint64]struct{}, model.AdjList) {
	destinationPoints := make(map[uint64]struct{})
	startPoint := <-spCh
	for pointsID := range dpCh {
		destinationPoints[pointsID[0]] = struct{}{}
	}
	adjacencyList := <-alCh
	return startPoint, destinationPoints, adjacencyList
}

func modifyAL(al *model.AdjList, sp coordinates.GeneralCoords) {
	nearestPointID := algorithm.Search(al, sp)
	al.Nodes[0] = sp
	al.AdjacencyList[0] = []uint64{nearestPointID}
}

func createInputChannel(test bool, svgCap int, csvCap int) (chan svg.SVGWriter, chan csv.CSVWriter, *sync.WaitGroup) {
	var wg sync.WaitGroup
	outMapChan := make(chan svg.SVGWriter, svgCap)
	outCSVChan := make(chan csv.CSVWriter, csvCap)
	wg.Add(2)

	if !test {
		os.MkdirAll("output/Task2/", 0777)
		go svg.ParallelWrite(outMapChan, &wg, "output/Task2/road_graph.svg")
		go csv.ParallelWrite(outCSVChan, &wg, "output/Task2/pathways.csv")
	} else {
		go svg.ParallelWrite(outMapChan, &wg, "/dev/null")
		go csv.ParallelWrite(outCSVChan, &wg, "/dev/null")
	}
	return outMapChan, outCSVChan, &wg
}
