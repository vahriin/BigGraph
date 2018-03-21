package svg

import (
	"bufio"
	"io"
)

type svg struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

func newSVG(writer io.WriteCloser) svg {
	var f svg

	f.writer = writer
	f.Buffer = bufio.NewWriter(writer)

	header := `<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">`
	f.Buffer.WriteString(header)
	f.Buffer.WriteRune('\n')

	return f
}

func (f svg) close() {
	end := "</svg>\n"
	f.Buffer.WriteString(end)
	f.Buffer.Flush()
	f.writer.Close()
}
