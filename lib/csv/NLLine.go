package csv

import (
	"io"
	"strconv"
	"strings"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

type NLLine struct {
	VertexID uint64
	coordinates.GeneralCoords
}

func NewNlLine(f csv) (NLLine, error) {
	str, err := f.Buffer.ReadString('\n')

	if err != nil {
		return NLLine{}, err
	}

	line := strings.Split(strings.TrimSuffix(str, "\n"), ",")

	if line[0] == "" {
		return NLLine{}, io.EOF
	}

	var nll NLLine
	nll.VertexID, err = strconv.ParseUint(line[0], 10, 64)
	if err != nil {
		return NLLine{}, err
	}

	nll.Earth.Lat, err = strconv.ParseFloat(line[1], 64)
	if err != nil {
		return NLLine{}, err
	}

	nll.Earth.Lon, err = strconv.ParseFloat(line[2], 64)
	if err != nil {
		return NLLine{}, err
	}

	nll.Euclid.X, err = strconv.ParseFloat(line[3], 64)
	if err != nil {
		return NLLine{}, err
	}

	nll.Euclid.Y, err = strconv.ParseFloat(line[4], 64)
	if err != nil {
		return NLLine{}, err
	}

	return nll, nil
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
