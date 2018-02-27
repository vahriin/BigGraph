package xmlparse

type Tag struct {
	Key string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

type Nd struct {
	Ref uint `xml:"ref,attr"`
}

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

type Node struct {
	Id uint `xml:"id,attr"`
	Coordinates
}

type Bounds struct {
	Minlat float64 `xml:"minlat,attr"`
	Maxlat float64 `xml:"maxlat,attr"`
	Minlon float64 `xml:"minlon,attr"`
	Maxlon float64 `xml:"maxlon,attr"`
}





