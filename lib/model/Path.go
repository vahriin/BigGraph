package model

import (
	"github.com/vahriin/BigGraph/lib/coordinates"
)

type Path struct {
	Points []uint64
	Len    float64
}

func (p Path) Start() uint64 {
	return p.Points[0]
}

func (p Path) End() uint64 {
	return p.Points[len(p.Points)-1]
}

func (p Path) Coordinates(al AdjList) []coordinates.EuclidCoords {
	coords := make([]coordinates.EuclidCoords, len(p.Points))

	for i, pointID := range p.Points {
		coords[i] = al.Nodes[pointID].Euclid
	}

	return coords
}

func (p Path) Copy() Path {
	cp := Path{
		Points: make([]uint64, len(p.Points)),
		Len:    p.Len,
	}

	copy(cp.Points, p.Points)
	return cp
}

func (p Path) Reverse() Path {
	cp := Path{
		Points: make([]uint64, len(p.Points)),
		Len:    p.Len,
	}

	for i, id := range p.Points {
		cp.Points[len(p.Points)-i-1] = id
	}

	return cp
}
