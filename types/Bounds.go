package types

// Bounds is container for latitude and longitude of angles of rectangle area
type Bounds struct {
	Minlat float64 `xml:"minlat,attr"`
	Maxlat float64 `xml:"maxlat,attr"`
	Minlon float64 `xml:"minlon,attr"`
	Maxlon float64 `xml:"maxlon,attr"`
}

// Maxs returns EuclidCoords of area's top right corner
func (b Bounds) Maxs() EuclidCoords {
	return GeographicCoords{Lon: b.Maxlon, Lat: b.Maxlat}.EuclidCoords()
}

// Mins returns EuclidCoords of area's lower left corner
func (b Bounds) Mins() EuclidCoords {
	return GeographicCoords{Lon: b.Minlon, Lat: b.Minlat}.EuclidCoords()
}
