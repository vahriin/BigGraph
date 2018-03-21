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

func (poly Polyline) SVGWrite(f svg) {
	f.Buffer.WriteString("<polyline stroke=\"")
	f.Buffer.WriteString(poly.Color)
	f.Buffer.WriteString("\" stroke-width=\"")
	f.Buffer.WriteString(strconv.FormatUint(uint64(poly.Width), 10))
	f.Buffer.WriteString("\" fill=\"none\" points=\" ")
	for i, c := range poly.Points {
		f.Buffer.WriteString(strconv.FormatFloat(c.X, 'f', -1, 64))
		f.Buffer.WriteRune(',')
		f.Buffer.WriteString(strconv.FormatFloat(c.Y, 'f', -1, 64))

		if i != len(poly.Points)-1 {
			f.Buffer.WriteRune(' ')
		}
	}

	f.Buffer.WriteString("\" />")
	f.Buffer.WriteRune('\n')
}
