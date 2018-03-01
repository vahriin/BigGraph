package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/vahriin/BigGraph/graph"
	"github.com/vahriin/BigGraph/xmlparse"
)

func Flags() (filename string) {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "map.osm"
	} else {
		return args[0]
	}
}

func main() {
	mapsrc := Flags()

	start := time.Now()

	fmt.Println("File parsing...")

	doc := xmlparse.XMLRead(mapsrc)
	area := doc.Graph()
	doc = nil

	fmt.Println("File parsed. Time spent: ", time.Since(start))
	fmt.Println("Generate output...")

	os.Mkdir("output", os.ModePerm)
	graph.SVGImage(&area, "output/road_graph.svg")
	graph.CSVNodeList(&area, "output/nodes_list.csv")
	graph.AdjList(&area, "output/adjacency_list.csv")

	fmt.Println("Output generated. Time spent total: ", time.Since(start))
	fmt.Println("Have a nice day!")
}
