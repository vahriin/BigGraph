package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/vahriin/BigGraph/lib/algorithm"

	"github.com/vahriin/BigGraph/Task3/input"
	"github.com/vahriin/BigGraph/Task3/output"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
)

func main() {
	start := time.Now()

	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	test := pf()

	adjacencyList, travelPoints, startPoint := input.Input("input/Task3/")

	// file-writer goroutines
	outSVGChan, outCSVChan, wg := createOutChannel(test, 10000, len(travelPoints))

	// draw graph
	var dgwg sync.WaitGroup
	dgwg.Add(1)
	go output.DrawGraph(outSVGChan, &dgwg, adjacencyList)

	pathChan := make(chan model.Path, 10)

	go algorithm.NearestVertexAdd(pathChan, travelPoints, 0, adjacencyList)

	dgwg.Wait()

	output.ProcessPath(outCSVChan, outSVGChan, adjacencyList, pathChan)

	close(outCSVChan)

	outSVGChan <- svg.Circle{Center: adjacencyList.Nodes[startPoint].Euclid,
		Color: "green", Radius: svg.PointAttentionRadius}
	close(outSVGChan)

	wg.Wait()

	fmt.Println(time.Since(start))
}

func createOutChannel(test bool, svgCap int, csvCap int) (chan svg.SVGWriter, chan csv.CSVWriter, *sync.WaitGroup) {
	var wg sync.WaitGroup
	outMapChan := make(chan svg.SVGWriter, svgCap)
	outCSVChan := make(chan csv.CSVWriter, csvCap)
	wg.Add(2)

	if !test {
		os.MkdirAll("output/Task3/", 0777)
		go svg.ParallelWrite(outMapChan, &wg, "output/Task3/road_graph.svg")
		go csv.ParallelWrite(outCSVChan, &wg, "output/Task3/pathway.csv")
	} else {
		go svg.ParallelWrite(outMapChan, &wg, "/dev/null")
		go csv.ParallelWrite(outCSVChan, &wg, "/dev/null")
	}
	return outMapChan, outCSVChan, &wg
}

func pf() bool {
	test := flag.Bool("t", false, "no output file")
	flag.Parse()
	return *test
}
