package svg

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

type svg struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

type Polyline struct {
	Points []coordinates.EuclidCoords
	Width  uint8
	Color  string
}

func (poly Polyline) SVGWrite(s svg) {
	s.Buffer.WriteString("<polyline stroke=\"")
	s.Buffer.WriteString(poly.Color)
	s.Buffer.WriteString("\" stroke-width=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(poly.Width), 10))
	s.Buffer.WriteString("\" fill=\"none\" points=\" ")
	for i, c := range poly.Points {
		s.Buffer.WriteString(strconv.FormatFloat(c.X, 'f', -1, 64))
		s.Buffer.WriteRune(',')
		s.Buffer.WriteString(strconv.FormatFloat(c.Y, 'f', -1, 64))

		if i != len(poly.Points)-1 {
			s.Buffer.WriteRune(' ')
		}
	}

	s.Buffer.WriteString("\" />")
	s.Buffer.WriteRune('\n')
}

type Circle struct {
	Center coordinates.EuclidCoords
	Radius uint8
	Color  string
}

func (circle Circle) SVGWrite(s svg) {
	s.Buffer.WriteString("<circle cx=\"")
	s.Buffer.WriteString(strconv.FormatFloat(circle.Center.X, 'f', -1, 64))
	s.Buffer.WriteString("\" cy=\"")
	s.Buffer.WriteString(strconv.FormatFloat(circle.Center.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" r=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(circle.Radius), 10))
	s.Buffer.WriteString("\" fill=\"")
	s.Buffer.WriteString(circle.Color)
	s.Buffer.WriteString("\" />")
	s.Buffer.WriteRune('\n')
}

type Line struct {
	Begin coordinates.EuclidCoords
	End   coordinates.EuclidCoords
	Width uint8
	Color string
}

func (line Line) SVGWrite(s svg) {
	s.Buffer.WriteString("<line x1=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.Begin.X, 'f', -1, 64))
	s.Buffer.WriteString("\" x2=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.End.X, 'f', -1, 64))
	s.Buffer.WriteString("\" y1=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.Begin.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" y2=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.End.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" stroke=\"")
	s.Buffer.WriteString(line.Color)
	s.Buffer.WriteString("\" stroke-width=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(line.Width), 10))
	s.Buffer.WriteString("\"/>")
	s.Buffer.WriteRune('\n')
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

func (s svg) Circle(circle Circle) {
	s.Buffer.WriteString("<circle cx=\"")
	s.Buffer.WriteString(strconv.FormatFloat(circle.Center.X, 'f', -1, 64))
	s.Buffer.WriteString("\" cy=\"")
	s.Buffer.WriteString(strconv.FormatFloat(circle.Center.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" r=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(circle.Radius), 10)) //TODO:
	s.Buffer.WriteString("\" fill=\"")
	s.Buffer.WriteString(circle.Color)
	s.Buffer.WriteString("\" />")
	s.Buffer.WriteRune('\n')
}

func (s svg) Line(line Line) {
	s.Buffer.WriteString("<line x1=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.Begin.X, 'f', -1, 64))
	s.Buffer.WriteString("\" x2=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.End.X, 'f', -1, 64))
	s.Buffer.WriteString("\" y1=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.Begin.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" y2=\"")
	s.Buffer.WriteString(strconv.FormatFloat(line.End.Y, 'f', -1, 64))
	s.Buffer.WriteString("\" stroke=\"")
	s.Buffer.WriteString(line.Color)
	s.Buffer.WriteString("\" stroke-width=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(line.Width), 10))
	s.Buffer.WriteString("\"/>")
	s.Buffer.WriteRune('\n')
}

func (s svg) Polyline(poly Polyline) {
	s.Buffer.WriteString("<polyline stroke=\"")
	s.Buffer.WriteString(poly.Color)
	s.Buffer.WriteString("\" stroke-width=\"")
	s.Buffer.WriteString(strconv.FormatUint(uint64(poly.Width), 10))
	s.Buffer.WriteString("\" fill=\"none\" points=\" ")
	for i, c := range poly.Points {
		s.Buffer.WriteString(strconv.FormatFloat(c.X, 'f', -1, 64))
		s.Buffer.WriteRune(',')
		s.Buffer.WriteString(strconv.FormatFloat(c.Y, 'f', -1, 64))

		if i != len(poly.Points)-1 {
			s.Buffer.WriteRune(' ')
		}
	}

	s.Buffer.WriteString("\" />")
	s.Buffer.WriteRune('\n')
}
