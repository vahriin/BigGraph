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
	var s svg

	s.writer = writer
	s.Buffer = bufio.NewWriter(writer)

	header := `<svg version="1.1" baseProfile="full" xmlns="http://www.w3.org/2000/svg">`
	s.Buffer.WriteString(header)
	s.Buffer.WriteRune('\n')

	return s
}

func (s svg) close() {
	end := "</svg>\n"
	s.Buffer.WriteString(end)
	s.Buffer.Flush()
	s.writer.Close()
}
