package types

import (
	"math"
)

// radius of Earth sphere
const radius = 1

// GeographicCoords is coordinates of several point (node)
type GeographicCoords struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

// EuclidCoords is coordinates of point on the plain
type EuclidCoords struct {
	X float64
	Y float64
}

// GeneralCoords is container for Euclid and Geographic coordinates
type GeneralCoords struct {
	Euclid EuclidCoords
	Earth  GeographicCoords
}

// EuclidCoords converts geographic coordinates into Euclidian coordinates. 
// The transformation takes place on a sphere of radius 1
func (c GeographicCoords) EuclidCoords() EuclidCoords {
	var ec EuclidCoords

	latRad := c.Lat * math.Pi / 180
	lonRad := c.Lon * math.Pi / 180

	ec.X = radius * lonRad
	ec.Y = math.Log(math.Tan(math.Pi/4 + latRad/2))

	ec.X *= Multiplier
	ec.Y *= Multiplier

	return ec
}
