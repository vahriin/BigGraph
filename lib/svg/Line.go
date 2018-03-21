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

func (line Line) SVGWrite(f svg) {
	f.Buffer.WriteString("<line x1=\"")
	f.Buffer.WriteString(strconv.FormatFloat(line.Begin.X, 'f', -1, 64))
	f.Buffer.WriteString("\" x2=\"")
	f.Buffer.WriteString(strconv.FormatFloat(line.End.X, 'f', -1, 64))
	f.Buffer.WriteString("\" y1=\"")
	f.Buffer.WriteString(strconv.FormatFloat(line.Begin.Y, 'f', -1, 64))
	f.Buffer.WriteString("\" y2=\"")
	f.Buffer.WriteString(strconv.FormatFloat(line.End.Y, 'f', -1, 64))
	f.Buffer.WriteString("\" stroke=\"")
	f.Buffer.WriteString(line.Color)
	f.Buffer.WriteString("\" stroke-width=\"")
	f.Buffer.WriteString(strconv.FormatUint(uint64(line.Width), 10))
	f.Buffer.WriteString("\"/>")
	f.Buffer.WriteRune('\n')
}
