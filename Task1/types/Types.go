package types

// Tag is data from <tag>
type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

// Nd is data from way's <tag>
type Nd struct {
	Ref uint64 `xml:"ref,attr"`
}

// Node is point on the Earth.
type Node struct {
	ID uint64 `xml:"id,attr"`
	GeographicCoords
}
