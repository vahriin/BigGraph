package types

type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

type Nd struct {
	Ref uint64 `xml:"ref,attr"`
}

type Node struct {
	Id uint64 `xml:"id,attr"`
	EarthCoords
}
