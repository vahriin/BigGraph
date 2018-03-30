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

	for w := range svgw {
		w.SVGWrite(svgImage)
	}
}
