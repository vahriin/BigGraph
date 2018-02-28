package types

type Coordinates struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

func (c Coordinates) Xint() uint {
	return uint(c.Lon * 1000000)
}

func (c Coordinates) Yint() uint {
	return uint(c.Lat * 1000000)
}