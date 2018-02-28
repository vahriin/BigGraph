package graph

import (
	"log"
	"os"

	"github.com/vahriin/BigGraph/svg"
	"github.com/vahriin/BigGraph/types"
)

func SVGImage(area types.Area, filename string) {
	file, _ := os.Create("logger")
	log.SetOutput(file)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	svgImage := svg.NewSVG(file)

	defer svgImage.Close()

	rectMin := area.Border.Mins()
	rectMax := area.Border.Maxs()

	for _, way := range area.Edges {
		polyline := make([]types.EuclidCoords, 0, len(way.Nodes))
		for _, node := range way.Nodes {
			nodeEC := node.EuclidCoords()
			nodeEC.X -= rectMin.X
			nodeEC.Y = rectMax.Y - nodeEC.Y

			polyline = append(polyline, nodeEC)
			svgImage.Circle(nodeEC, 2)
		}
		svgImage.Polyline(polyline, 1)
	}
}
