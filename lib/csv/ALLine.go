package csv

import (
	"io"
	"strconv"
	"strings"
)

type ALLine struct {
	VertexID           uint64
	IncidentVertexesID []uint64
}

func NewAlLine(f csv) (ALLine, error) {
	str, err := f.Buffer.ReadString('\n')

	if err != nil {
		return ALLine{}, err
	}

	line := strings.Split(strings.TrimSuffix(str, "\n"), ",")

	if line[0] == "" {
		return ALLine{}, io.EOF
	}

	var all ALLine
	all.IncidentVertexesID = make([]uint64, 0, 4)

	all.VertexID, err = strconv.ParseUint(line[0], 10, 64)
	if err != nil {
		return ALLine{}, err
	}

	for i := 1; i < len(line); i++ {
		node, err := strconv.ParseUint(line[i], 10, 64)
		if err != nil {
			return ALLine{}, err
		}
		all.IncidentVertexesID = append(all.IncidentVertexesID, node)
	}

	return all, nil
}

func (al ALLine) CSVWrite(f csv) {
	f.Buffer.WriteString(strconv.FormatUint(al.VertexID, 10))
	f.comma()
	for i, child := range al.IncidentVertexesID {
		f.Buffer.WriteString(strconv.FormatUint(child, 10))
		if i != len(al.IncidentVertexesID)-1 {
			f.comma()
		}
	}
	f.Buffer.WriteRune('\n')
}
