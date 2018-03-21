package csv

import "strconv"

type ALLine struct {
	VertexID           uint64
	IncidentVertexesID []uint64
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
