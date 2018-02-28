package csv

import (
	"bufio"
	"io"
	"strconv"

	"github.com/vahriin/BigGraph/types"
)

type CSV struct {
	writer io.WriteCloser
	Buffer *bufio.Writer
}

func NewCSV(closer io.WriteCloser) CSV {
	var csv CSV

	csv.writer = closer
	csv.Buffer = bufio.NewWriter(csv.writer)

	return csv
}

func (csv CSV) Close() {
	csv.Buffer.WriteRune('\n')
	csv.Buffer.Flush()
	csv.writer.Close()
}

func (csv CSV) ALLine(root uint64, children []uint64) {
	csv.Buffer.WriteString(strconv.FormatUint(root, 10))
	csv.comma()
	for i, child := range children {
		csv.Buffer.WriteString(strconv.FormatUint(child, 10))
		if i != len(children)-1 {
			csv.comma()
		}
	}
	csv.Buffer.WriteRune('\n')
}

func (csv CSV) AMLine(distances []float64) {
	for i, dist := range distances {
		csv.Buffer.WriteString(strconv.FormatFloat(dist, 'f', -1, 64))
		if i != len(distances)-1 {
			csv.comma()
		}
	}
	csv.Buffer.WriteRune('\n')
}

func (csv CSV) NLLine(node *types.Node) {
	csv.Buffer.WriteString(strconv.FormatUint(node.Id, 10))
	csv.comma()
	csv.Buffer.WriteString(strconv.FormatFloat(node.Lat, 'f', -1, 64))
	csv.comma()
	csv.Buffer.WriteString(strconv.FormatFloat(node.Lon, 'f', -1, 64))
	csv.comma()
	ec := node.EuclidCoords()
	csv.Buffer.WriteString(strconv.FormatUint(ec.X, 10))
	csv.comma()
	csv.Buffer.WriteString(strconv.FormatUint(ec.Y, 10))
}

func (csv CSV) comma() {
	csv.Buffer.WriteRune(',')
}
