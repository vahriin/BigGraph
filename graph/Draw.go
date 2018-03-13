package graph

import (
	"os"
	"sync"

	"github.com/vahriin/BigGraph/svg"
	"github.com/vahriin/BigGraph/types"
)

// SVGImage draw the AdjList to .svg file
func SVGImage(al types.AdjList, filename string) {
	var wg sync.WaitGroup
	wg.Add(1)

	lineCh := make(chan [2]types.EuclidCoords, 1000)
	pointCh := make(chan types.EuclidCoords, len(al.AL))

	go writeSVG(lineCh, pointCh, filename, &wg)

	processed := make(map[uint64]bool)

	var line [2]types.EuclidCoords

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

func writeSVG(lines <-chan [2]types.EuclidCoords, points <-chan types.EuclidCoords, filename string, wg *sync.WaitGroup) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	svgImage := svg.NewSVG(file)

	defer func() {
		svgImage.Close()
		wg.Done()
	}()

	for line := range lines {
		svgImage.Line(line[0], line[1], 1)
	}

	for point := range points {
		svgImage.Circle(point, 1)
	}
}
