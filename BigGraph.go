package main

import (
	"fmt"
	"time"

	"github.com/vahriin/BigGraph/xml"
)

func main() {
	start := time.Now()
	doc := xml.XMLRead("/home/vahriin/Downloads/map")
	graph := doc.Graph()
	fmt.Println(len(graph.Edges))
	fmt.Println(graph.Edges[55])
	fmt.Println(time.Since(start))
}
