package model

import "github.com/vahriin/BigGraph/lib/coordinates"

type AdjList struct {
	AdjacencyList map[uint64][]uint64
	Nodes         map[uint64]coordinates.GeneralCoords
}
