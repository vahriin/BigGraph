package model

import (
	"github.com/vahriin/BigGraph/lib/coordinates"
)

type Path struct {
	Points []uint64
	Len    float64
}

func (p Path) Coordinates(al AdjList) []coordinates.EuclidCoords {
	coords := make([]coordinates.EuclidCoords, len(p.Points))

	for i, pointID := range p.Points {
		coords[i] = al.Nodes[pointID].Euclid
	}

	return coords
}
