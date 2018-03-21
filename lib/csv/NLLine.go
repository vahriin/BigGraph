package csv

import (
	"strconv"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

type NLLine struct {
	VertexID uint64
	coordinates.GeneralCoords
}

func (nl NLLine) CSVWrite(f csv) {
	f.Buffer.WriteString(strconv.FormatUint(nl.VertexID, 10))
	f.comma()
	f.Buffer.WriteString(strconv.FormatFloat(nl.Earth.Lat, 'f', -1, 64))
	f.comma()
	f.Buffer.WriteString(strconv.FormatFloat(nl.Earth.Lon, 'f', -1, 64))
	f.comma()
	f.Buffer.WriteString(strconv.FormatFloat(nl.Euclid.X, 'f', -1, 64))
	f.comma()
	f.Buffer.WriteString(strconv.FormatFloat(nl.Euclid.Y, 'f', -1, 64))
	f.Buffer.WriteRune('\n')
}
