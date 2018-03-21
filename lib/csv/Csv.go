package csv

import (
	"bufio"
	"io"
)

type csv struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

func newCSV(wc io.WriteCloser) csv {
	var f csv

	f.writer = wc
	f.Buffer = bufio.NewWriter(f.writer)

	return f
}

func (f csv) close() {
	f.Buffer.WriteRune('\n')
	f.Buffer.Flush()
	f.writer.Close()
}

func (f csv) comma() {
	f.Buffer.WriteRune(',')
}
