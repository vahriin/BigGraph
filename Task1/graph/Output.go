package graph

import (
	"os"
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/svg"
)

func Output(al types.AdjList) {
	var wg sync.WaitGroup

	os.Mkdir("output", os.ModePerm)

	svgChan := make(chan svg.SVGWriter, 1000)
	wg.Add(2)
	go ProcessSVG(svgChan, &wg, al)
	go svg.ParallelWrite(svgChan, &wg, "output/road_graph.svg")

	alChan := make(chan csv.CSVWriter, 1000)
	wg.Add(2)
	go ProcessAlCSV(alChan, &wg, al)
	go csv.ParallelWrite(alChan, &wg, "output/adjacency_list.csv")

	nlChan := make(chan csv.CSVWriter, 1000)
	wg.Add(2)
	go ProcessAlCSV(nlChan, &wg, al)
	go csv.ParallelWrite(nlChan, &wg, "output/nodes_list.csv")

	wg.Wait()
}
