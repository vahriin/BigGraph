package main

import (
	"fmt"
	"time"

	"github.com/vahriin/BigGraph/graph"
	"github.com/vahriin/BigGraph/xmlparse"
)

func main() {
	start := time.Now()
	doc := xmlparse.XMLRead("/home/vahriin/Downloads/map")
	startComp := time.Now()
	area := doc.Graph()
	graph.SVGImage(area, "output1.svg")
	fmt.Println(time.Since(start))
	fmt.Println(time.Since(startComp))
}
