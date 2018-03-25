package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/vahriin/BigGraph/Task1/graph"
	"github.com/vahriin/BigGraph/Task1/xmlparse"
)

func flags() (filename string) {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "input/Task1/map.osm"
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
	fmt.Println("File parsed. Time spent: ", time.Since(start), "\n")

	fmt.Println("Building AdjList...")
	adjList := doc.AdjList()
	doc = nil
	fmt.Println("Done, time spent ", time.Since(start), "\n")

	fmt.Println("Delete excess points...")
	fmt.Println(adjList.DropExcessPoints(), " excess points deleted. Time spent: ", time.Since(start), "\n")

	fmt.Println("Generate output...")

	graph.Output(adjList)

	fmt.Println("Output generated. Time spent total: ", time.Since(start), "\n")
	fmt.Println("Have a nice day!")
}
