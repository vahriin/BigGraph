package output

import (
	"sync"

	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
)

func DrawGraph(svgChan chan<- svg.SVGWriter, wg *sync.WaitGroup, al model.AdjList) {
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

	wg.Done()
}
