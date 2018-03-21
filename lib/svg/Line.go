package svg

import (
	"strconv"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

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
