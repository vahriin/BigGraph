package types

type Edge struct {
	NodesId []uint64
}

type Area struct {
	Edges  []Edge
	Points map[uint64]GeneralCoords
}

func (area Area) PointsId() []uint64 {
	keys := make([]uint64, 0, len(area.Points))
	for k := range area.Points {
		keys = append(keys, k)
	}
	return keys
}
