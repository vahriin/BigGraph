package csv

import (
	"io"
	"strconv"
	"strings"
)

type Line []uint64

func NewLine(f csv) (Line, error) {
	str, err := f.Buffer.ReadString('\n')
	if err != nil {
		return Line{}, err
	}

	l := strings.Split(strings.TrimSuffix(str, "\n"), ",")

	if l[0] == "" {
		return Line{}, io.EOF
	}

	line := make(Line, 0, 10)
	for _, id := range l {
		id, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return Line{}, err
		}

		line = append(line, id)
	}
	return line, nil
}

func (l Line) CSVWrite(f csv) {
	for i, id := range l {
		f.Buffer.WriteString(strconv.FormatUint(id, 10))
		if i != len(l)-1 {
			f.comma()
		}
	}
	f.Buffer.WriteRune('\n')
}
