package main

import (
	"flag"
	"fmt"
	"time"

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
	mapsrc := flags()

	start := time.Now()

	fmt.Println("File parsing...")

	doc := xmlparse.XMLRead(mapsrc)
	adjList := doc.AdjList()
	doc = nil

	fmt.Println(len(adjList.AL))
	fmt.Println(len(adjList.Nodes))

	a := adjList.DropExcessPoints()

	fmt.Println(a)

	fmt.Println("File parsed. Time spent: ", time.Since(start))
	/*fmt.Println("Generate output...")

	os.Mkdir("output", os.ModePerm)
	graph.SVGImage(&area, "output/viz.svg")
	graph.CSVNodeList(&area, "output/NL.csv")
	graph.AdjList(&area, "output/AL.csv")

	fmt.Println("Output generated. Time spent total: ", time.Since(start))
	fmt.Println("Have a nice day!")*/
}
