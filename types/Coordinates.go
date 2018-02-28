package types

import (
	"math"
)

const radius = 1
const multiplier = 1000000

type EarthCoords struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

func (c EarthCoords) EuclidCoords() EuclidCoords {
	var ec EuclidCoords

	latRad := c.Lat * math.Pi / 180
	lonRad := c.Lon * math.Pi / 180

	ec.X = radius * lonRad
	ec.Y = math.Log(math.Tan(math.Pi/4 + latRad/2))

	ec.X *= multiplier
	ec.Y *= multiplier

	return ec
}

type EuclidCoords struct {
	X float64
	Y float64
}

type GeneralCoords struct {
	Euclid EuclidCoords
	Earth  EarthCoords
}
