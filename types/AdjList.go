package types

// AdjList is (TODO: Add doc)
type AdjList struct {
    AL map[uint64][]uint64
    Nodes map[uint64]GeneralCoords
}

