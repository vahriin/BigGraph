package svg

import (
	"bufio"
	"io"
	"strconv"

	"github.com/vahriin/BigGraph/types"
)

type SVG struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

func NewSVG(writer io.WriteCloser, width, height uint64) SVG {
	var svg SVG

	svg.writer = writer
	svg.Buffer = bufio.NewWriter(writer)

	header := `
<svg version="1.1" baseProfile="full" width="` + strconv.FormatUint(width, 10) + `" height="` + strconv.FormatUint(height, 10) + `" xmlns="http://www.w3.org/2000/svg">`
	svg.Buffer.WriteString(header)
	svg.Buffer.WriteRune('\n')

	return svg
}

func (svg SVG) Close() {
	end := "</svg>\n"
	svg.Buffer.WriteString(end)
	svg.Buffer.Flush()
	svg.writer.Close()
}

func (svg SVG) Circle(c types.EuclidCoords, r uint64) {
	svg.Buffer.WriteString("<circle cx=\"")
	svg.Buffer.WriteString(strconv.FormatUint(c.X, 10))
	svg.Buffer.WriteString("\" cy=\"")
	svg.Buffer.WriteString(strconv.FormatUint(c.Y, 10))
	svg.Buffer.WriteString("\" r=\"")
	svg.Buffer.WriteString(strconv.FormatUint(r, 10))
	svg.Buffer.WriteString("\" fill=\"red\" />")
	svg.Buffer.WriteRune('\n')
}

func (svg SVG) Line(begin, end types.EuclidCoords) {
	svg.Buffer.WriteString("<line x1=\"")
	svg.Buffer.WriteString(strconv.FormatUint(begin.X, 10))
	svg.Buffer.WriteString("\" x2=\"")
	svg.Buffer.WriteString(strconv.FormatUint(end.X, 10))
	svg.Buffer.WriteString("\" y1=\"")
	svg.Buffer.WriteString(strconv.FormatUint(begin.Y, 10))
	svg.Buffer.WriteString("\" y2=\"")
	svg.Buffer.WriteString(strconv.FormatUint(end.Y, 10))
	svg.Buffer.WriteString("\" stroke=\"blue\" stroke-width=\"25\"/>")
	svg.Buffer.WriteRune('\n')
}

func (svg SVG) Polyline(points []types.EuclidCoords) {
	svg.Buffer.WriteString("<polyline stroke=\"blue\" stroke-width=\"25\" fill=\"none\" points=\" ")
	for i, c := range points {
		svg.Buffer.WriteString(strconv.FormatUint(c.X, 10))
		svg.Buffer.WriteRune(',')
		svg.Buffer.WriteString(strconv.FormatUint(c.Y, 10))

		if i != len(points)-1 {
			svg.Buffer.WriteRune(' ')
		}
	}

	svg.Buffer.WriteString("\" />")
	svg.Buffer.WriteRune('\n')
}
