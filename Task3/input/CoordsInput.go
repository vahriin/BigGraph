package input

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

func ReadPoint(filename string) coordinates.GeneralCoords {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var gc coordinates.GeographicCoords

	err = xml.Unmarshal(file, &gc)
	if err != nil {
		panic(err)
	}
	c := coordinates.GeneralCoords{Earth: gc, Euclid: gc.EuclidCoords()}
	c.Euclid.Y = -c.Euclid.Y
	return c
}
