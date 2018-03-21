package svg

import (
	"strconv"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

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
