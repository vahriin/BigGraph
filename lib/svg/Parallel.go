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

	for w := range svgw {
		if circle, ok := w.(Circle); ok {
			circles = append(circles, circle)
		}
		w.SVGWrite(svgImage)
	}

	for _, circle := range circles {
		circle.SVGWrite(svgImage)
	}
}
