package svg

import (
	"os"
	"sync"
)

func ParallelWrite(svgw <-chan SVGWriter, wg *sync.WaitGroup, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	svgImage := newSVG(file)

	defer func() {
		svgImage.close()
		wg.Done()
	}()

	circles := make([]Circle, 0, 10000)
	polylines := make([]Polyline, 0, 10)

	for w := range svgw {
		switch w.(type) {
		case Circle:
			circles = append(circles, w.(Circle))
		case Polyline:
			polylines = append(polylines, w.(Polyline))
		case Line:
			w.SVGWrite(svgImage)
		}
	}

	for _, poly := range polylines {
		poly.SVGWrite(svgImage)
	}

	for _, circle := range circles {
		circle.SVGWrite(svgImage)
	}
}
