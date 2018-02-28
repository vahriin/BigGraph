package types

type Edge struct {
	NodesId []uint64
}

type Area struct {
	Border Bounds
	Edges  []Edge
	Points map[uint64]GeneralCoords
}
