package svg

import (
	"strconv"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

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
