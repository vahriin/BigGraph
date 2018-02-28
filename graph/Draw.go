package graph

import (
	"os"

	"github.com/vahriin/BigGraph/svg"
	"github.com/vahriin/BigGraph/types"
)

func SVGImage(area types.Area, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	svgImage := svg.NewSVG(file)

	defer svgImage.Close()

	for _, way := range area.Edges {
		polyline := make([]types.EuclidCoords, 0, len(way.NodesId))
		for _, nodeId := range way.NodesId {
			polyline = append(polyline, area.Points[nodeId].Euclid)
		}
		svgImage.Polyline(polyline, 1)
	}

	for _, coord := range area.Points {
		svgImage.Circle(coord.Euclid, 1)
	}
}
