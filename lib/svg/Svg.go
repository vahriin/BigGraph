package svg

import (
	"bufio"
	"io"
	"os"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

type svg struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

func ParallelWrite(lines <-chan [2]coordinates.EuclidCoords, points <-chan coordinates.EuclidCoords, filename string, wg *sync.WaitGroup) {
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

func NewSVG(writer io.WriteCloser) svg {
	var s svg

	s.writer = writer
	s.Buffer = bufio.NewWriter(writer)

	header := `<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">`
	s.Buffer.WriteString(header)
	s.Buffer.WriteRune('\n')

	return s
}

func (s svg) Close() {
	end := "</svg>\n"
	s.Buffer.WriteString(end)
	s.Buffer.Flush()
	s.writer.Close()
}
