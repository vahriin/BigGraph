package aco

import (
	"github.com/vahriin/BigGraph/lib/model"
)

type PathMatrix struct {
	pathways   [][]model.Path
	pheromones [][]float64
	index      map[uint64]int
}

func NewPathMatrix(start uint64, travelPoints map[uint64]struct{}) PathMatrix {
	var pm PathMatrix
	pm.index = make(map[uint64]int, len(travelPoints)+1)
	pm.index[start] = 0
	matrixIndex := 1
	for tp := range travelPoints {
		pm.index[tp] = matrixIndex
		matrixIndex++
	}

	pm.pathways = make([][]model.Path, len(travelPoints))
	for i := 0; i < len(travelPoints); i++ {
		pm.pathways[i] = make([]model.Path, len(travelPoints))
	}
	pm.pheromones = make([][]float64, len(travelPoints))
	for i := 0; i < len(travelPoints); i++ {
		pm.pheromones[i] = make([]float64, len(travelPoints))
	}
	return pm
}

func (pm PathMatrix) Pheromone(i, j uint64) float64 {
	return pm.pheromones[pm.index[i]][pm.index[j]]
}

func (pm *PathMatrix) SetPheromone(i, j uint64, value float64) {
	pm.pheromones[pm.index[i]][pm.index[j]] = value
}

func (pm PathMatrix) Distance(i, j uint64) float64 {
	return pm.pathways[pm.index[i]][pm.index[j]].Len
}

func (pm PathMatrix) Path(i, j uint64) model.Path {
	return pm.pathways[pm.index[i]][pm.index[j]]
}

func (pm *PathMatrix) SetPath(i, j uint64, path model.Path) {
	pm.pathways[pm.index[i]][pm.index[j]] = path
}
