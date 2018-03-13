package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/vahriin/BigGraph/graph"

	"github.com/vahriin/BigGraph/xmlparse"
)

func flags() (filename string) {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "map.osm"
	}
	return args[0]
}

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
	//runtime.GOMAXPROCS(1)

	mapsrc := flags()

	start := time.Now()

	fmt.Println("File parsing...")
	doc := xmlparse.XMLRead(mapsrc)
	adjList := doc.AdjList()
	doc = nil
	fmt.Println("File parsed. Time spent: ", time.Since(start), "\n")

	fmt.Println("Delete excess points...")
	fmt.Println(adjList.DropExcessPoints(), " excess points deleted. Time spent: ", time.Since(start), "\n")

	fmt.Println("Generate output...")

	os.Mkdir("output", os.ModePerm)

	var oh sync.WaitGroup
	oh.Add(2)

	go graph.CSVNodeList(adjList, "output/NL.csv", &oh)
	go graph.CSVAdjList(adjList, "output/AL.csv", &oh)

	graph.SVGImage(adjList, "output/viz.svg")

	oh.Wait()

	fmt.Println("Output generated. Time spent total: ", time.Since(start), "\n")
	fmt.Println("Have a nice day!")
}
