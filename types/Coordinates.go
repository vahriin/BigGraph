package types

type EarthCoords struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

func (c EarthCoords) IntLon() uint64 {
	return uint64(c.Lon * Multiplier)
}

func (c EarthCoords) IntLat() uint64 {
	return uint64(c.Lat * Multiplier)
}

func (c EarthCoords) EuclidCoords() EuclidCoords {

}

type EuclidCoords struct {
	X uint64
	Y uint64
}
