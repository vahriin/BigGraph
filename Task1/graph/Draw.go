package graph

import (
	"sync"

	"github.com/vahriin/BigGraph/Task1/types"
	"github.com/vahriin/BigGraph/lib/coordinates"
)

// SVGImage draw the AdjList to .svg file
func SVGImage(al types.AdjList, filename string) {
	var wg sync.WaitGroup
	wg.Add(1)

	lineCh := make(chan [2]coordinates.EuclidCoords, 1000)
	pointCh := make(chan coordinates.EuclidCoords, len(al.AL))

	go writeSVG(lineCh, pointCh, filename, &wg)

	processed := make(map[uint64]bool)

	var line [2]coordinates.EuclidCoords

	for node, incidentNodes := range al.AL {
		for _, in := range incidentNodes {
			if _, ok := processed[in]; !ok {
				line[0] = al.Nodes[node].Euclid
				line[1] = al.Nodes[in].Euclid
				lineCh <- line
			}
		}
		processed[node] = false
		pointCh <- al.Nodes[node].Euclid
	}

	close(lineCh)
	close(pointCh)

	wg.Wait()
}
