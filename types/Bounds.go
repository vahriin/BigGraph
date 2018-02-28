package types

type Bounds struct {
	Minlat float64 `xml:"minlat,attr"`
	Maxlat float64 `xml:"maxlat,attr"`
	Minlon float64 `xml:"minlon,attr"`
	Maxlon float64 `xml:"maxlon,attr"`
}

func (b Bounds) Maxs() EuclidCoords {
	return EarthCoords{Lon: b.Maxlon, Lat: b.Maxlat}.EuclidCoords()
}

func (b Bounds) Mins() EuclidCoords {
	return EarthCoords{Lon: b.Minlon, Lat: b.Minlat}.EuclidCoords()
}
