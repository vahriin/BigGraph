package output

import (
	"fmt"

	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
	"github.com/vahriin/BigGraph/lib/svg"
)

func ProcessPath(csvChan chan<- csv.CSVWriter, svgChan chan<- svg.SVGWriter, al model.AdjList, pathChan <-chan model.Path) {
	var pathLen float64

	outLine := csv.Line(make([]uint64, 0, 1000))

	circles := make([]svg.Circle, 22)

	colors := []string{"rgb(0,0,0)", "rgb(20,20,20)", "rgb(40,40,40)", "rgb(60,60,60)", "rgb(80,80,80)",
		"rgb(100,100,100)", "rgb(120,120,120)", "rgb(140,140,140)", "rgb(160,160,160)", "rgb(180,180,180)",
		"rgb(200,200,200)"}

	colorIndex := 0
	for path := range pathChan {
		polyline := svg.Polyline{Width: svg.PolylineWidth, Color: "red", Points: path.Coordinates(al)}
		circles = append(circles, svg.Circle{Center: polyline.Points[len(polyline.Points)-1],
			Color: colors[colorIndex], Radius: svg.PointAttentionRadius})

		svgChan <- polyline

		outLine = append(outLine, path.Points...)
		pathLen += path.Len

		colorIndex++
	}

	// delete duplicate id
	for i := 1; i < len(outLine); i++ {
		if outLine[i-1] == outLine[i] {
			outLine = append(outLine[:i-1], outLine[i:]...)
		}
	}

	csvChan <- outLine

	for _, circle := range circles {
		svgChan <- circle
	}

	fmt.Printf("Длина пути составила %d метров\n", int(pathLen))
}
