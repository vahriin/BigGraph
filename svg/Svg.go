package svg

import (
	"bufio"
	"os"
	"strconv"
)

type SVG struct {
	file *os.File
	Buffer *bufio.Writer
}

func NewSVG(filename string, width, heigh uint) SVG {
	var svg SVG
	var err error

	svg.file, err = os.Create(filename + ".svg")

	if err != nil {
		panic(err)
	}

	svg.Buffer = bufio.NewWriter(svg.file)

	header := `
<svg version="1.1" baseProfile="full" width="` + strconv.Itoa(int(width)) + `" height="` + strconv.Itoa(int(heigh)) + `" xmlns="http://www.w3.org/2000/svg">`
	svg.Buffer.WriteString(header)
	svg.Buffer.WriteRune('\n')

	return svg
}

func (svg SVG) Close() {
	end := "</svg>\n"
	svg.Buffer.WriteString(end)
	svg.Buffer.Flush()
	svg.file.Close()
}

func (svg SVG) Circle(x, y uint, r uint) {
	svg.Buffer.WriteString("<circle cx=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(x)))
	svg.Buffer.WriteString("\" cy=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(y)))
	svg.Buffer.WriteString("\" r=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(r)))
	svg.Buffer.WriteString("\" fill=\"red\" />")
	svg.Buffer.WriteRune('\n')
}

func (svg SVG) Line (x1, y1, x2, y2 uint) {
	svg.Buffer.WriteString("<line x1=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(x1)))
	svg.Buffer.WriteString("\" x2=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(x2)))
	svg.Buffer.WriteString("\" y1=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(y1)))
	svg.Buffer.WriteString("\" y2=\"")
	svg.Buffer.WriteString(strconv.Itoa(int(y2)))
	svg.Buffer.WriteString("\" stroke=\"blue\" stroke-width=\"25\"/>")
	svg.Buffer.WriteRune('\n')
}