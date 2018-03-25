package svg

import (
	"bufio"
	"io"
)

type svg struct {
	rw     io.ReadWriteCloser
	Buffer *bufio.ReadWriter
}

func newSVG(writer io.ReadWriteCloser) svg {
	var f svg

	f.rw = writer
	f.Buffer = bufio.NewReadWriter(bufio.NewReader(writer), bufio.NewWriter(writer))

	header := `<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">`
	f.Buffer.WriteString(header)
	f.Buffer.WriteRune('\n')

	return f
}

func (f svg) close() {
	end := "</svg>\n"
	f.Buffer.WriteString(end)
	f.Buffer.Flush()
	f.rw.Close()
}
