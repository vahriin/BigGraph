package csv

import (
	"io"
	"os"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

func ReadNodeList(filename string) map[uint64]coordinates.GeneralCoords {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer csvFile.close()

	nl := make(map[uint64]coordinates.GeneralCoords)

	var nlLine NLLine
	for {
		nlLine, err = NewNlLine(csvFile)
		if err == nil {
			nl[nlLine.VertexID] = nlLine.GeneralCoords
		} else {
			break
		}
	}

	if err != io.EOF {
		panic(err)
	}

	return nl
}

func ReadAdjacencyList(filename string) map[uint64][]uint64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer csvFile.close()

	al := make(map[uint64][]uint64)

	var alLine ALLine
	for {
		alLine, err = NewAlLine(csvFile)

		if err == nil {
			al[alLine.VertexID] = alLine.IncidentVertexesID
		} else {
			break
		}
	}

	if err != io.EOF {
		panic(err)
	}

	return al
}

// ReadPointsID return ID of points from destination_points.csv
func ReadPointsID(filename string) []uint64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer csvFile.close()

	res := make([]uint64, 0, 10)

	for {
		line, err := NewLine(csvFile)

		if err == nil {
			res = append(res, line[0])
		} else {
			break
		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}
	return res
}
