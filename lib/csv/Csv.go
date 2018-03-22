package csv

import (
	"bufio"
	"io"
)

type csv struct {
	rw     io.ReadWriteCloser
	Buffer *bufio.ReadWriter
}

func newCSV(rwc io.ReadWriteCloser) csv {
	var f csv

	f.rw = rwc
	f.Buffer = bufio.NewReadWriter(bufio.NewReader(f.rw), bufio.NewWriter(f.rw))

	return f
}

func (f csv) close() {
	f.Buffer.WriteRune('\n')
	f.Buffer.Flush()
	f.rw.Close()
}

func (f csv) comma() {
	f.Buffer.WriteRune(',')
}
