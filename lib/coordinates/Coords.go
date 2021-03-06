package coordinates

import (
	"math"
)

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
func (c GeographicCoords) EuclidCoords() EuclidCoords {
	var ec EuclidCoords

	latRad := c.Lat * math.Pi / 180
	lonRad := c.Lon * math.Pi / 180

	ec.X = Radius * lonRad
	ec.Y = Radius * math.Log(math.Tan(math.Pi/4+latRad/2))

	ec.X *= Multiplier
	ec.Y *= Multiplier

	return ec
}
