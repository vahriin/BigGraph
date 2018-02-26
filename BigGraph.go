package main

import (
	"fmt"
	"time"

	"github.com/vahriin/BigGraph/xmlpars"
)

func main() {
	start := time.Now()
	doc := xmlpars.XMLRead("/home/vahriin/Downloads/map")
	graph := doc.Graph()
	fmt.Println(len(graph.Edges))
	fmt.Println(graph.Edges[55])
	fmt.Println(time.Since(start))
}
