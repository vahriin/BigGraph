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

func (circle Circle) SVGWrite(f svg) {
	f.Buffer.WriteString("<circle cx=\"")
	f.Buffer.WriteString(strconv.FormatFloat(circle.Center.X, 'f', -1, 64))
	f.Buffer.WriteString("\" cy=\"")
	f.Buffer.WriteString(strconv.FormatFloat(circle.Center.Y, 'f', -1, 64))
	f.Buffer.WriteString("\" r=\"")
	f.Buffer.WriteString(strconv.FormatUint(uint64(circle.Radius), 10))
	f.Buffer.WriteString("\" fill=\"")
	f.Buffer.WriteString(circle.Color)
	f.Buffer.WriteString("\" />")
	f.Buffer.WriteRune('\n')
}
