package csv

import (
	"io"
	"os"
	"sync"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

func ParallelWrite(csvw <-chan CSVWriter, wg *sync.WaitGroup, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer func() {
		csvFile.close()
		wg.Done()
	}()

	for w := range csvw {
		w.CSVWrite(csvFile)
	}
}

func ReadNodeList(nlChan chan<- map[uint64]coordinates.GeneralCoords, filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer func() {
		csvFile.close()
		close(nlChan)
	}()

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

	nlChan <- nl
}

func ReadAdjacencyList(alChan chan<- map[uint64][]uint64, filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer func() {
		csvFile.close()
		close(alChan)
	}()

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

	alChan <- al
}

func ReadPointsID(lChan chan<- []uint64, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	csvFile := newCSV(file)

	defer func() {
		csvFile.close()
		close(lChan)
	}()

	for {
		line, err := NewLine(csvFile)

		if err == nil {
			lChan <- line
		} else {
			break
		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}
}
