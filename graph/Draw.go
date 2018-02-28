package graph

import (
	"fmt"

	"github.com/vahriin/BigGraph/svg"
	"github.com/vahriin/BigGraph/xmlparse"
)

func SVGImage(area xmlparse.Area) {
	ymin := uint(area.Border.Minlat * 1000000)
	ymax := uint(area.Border.Maxlat * 1000000)
	xmax := uint(area.Border.Maxlon * 1000000)
	xmin := uint(area.Border.Minlon * 1000000)
	widht := ymax - ymin
	height := xmax - xmin

	fmt.Println(widht, height)

	file := svg.NewSVG("output", widht, height)

	defer file.Close()

	badcounter := 0

	for _, way := range area.Edges {
		for i, node := range way.Nodes {
			file.Circle(node.Xint() - xmin, ymax - node.Yint(), 50)
			if i != len(way.Nodes) - 1 {
				file.Line(way.Nodes[i].Xint() - xmin, ymax - way.Nodes[i].Yint(),
					way.Nodes[i+1].Xint() - xmin, ymax - way.Nodes[i+1].Yint())
			}
		}
	}

	fmt.Println(badcounter)
}
