package types

// Area is main data structure contains the highway's node point and their coordinate
type Area struct {
	Highways []Highway
	Points   map[uint64]GeneralCoords
}

// PointsID return array contains ID of all point
func (area Area) PointsID() []uint64 {
	ids := make([]uint64, 0, len(area.Points))
	for id := range area.Points {
		ids = append(ids, id)
	}
	return ids
}
