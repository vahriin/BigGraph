package graph

import (
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/svg"
)

// ProcessSVG draw the AdjList to .svg file TODO: outdated
func ProcessSVG(svgChan chan<- svg.SVGWriter, wg *sync.WaitGroup, al types.AdjList) {
	processed := make(map[uint64]struct{})

	for node, incidentNodes := range al.AdjacencyList {
		for _, in := range incidentNodes {
			if _, ok := processed[in]; !ok {
				l := svg.Line{Begin: al.Nodes[node].Euclid, End: al.Nodes[in].Euclid,
					Color: svg.LineColor, Width: svg.LineWidth}
				svgChan <- l
			}
		}
		processed[node] = struct{}{}
		svgChan <- svg.Circle{Center: al.Nodes[node].Euclid, Color: svg.CircleColor, Radius: svg.PointRadius}
	}

	close(svgChan)

	wg.Done()
}
