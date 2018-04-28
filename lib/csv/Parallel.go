package csv

import (
	"os"
	"sync"
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
