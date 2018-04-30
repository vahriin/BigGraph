package aco

import (
	"sync"

	"github.com/vahriin/BigGraph/lib/model"
)

type PathMatrix struct {
	mutex      *sync.RWMutex
	pathways   [][]model.Path
	pheromones [][]float64
	cities     map[uint64]int
}

func NewPathMatrix(start uint64, travelPoints map[uint64]struct{}) PathMatrix {
	var pm PathMatrix
	pm.mutex = new(sync.RWMutex)
	pm.cities = make(map[uint64]int, len(travelPoints)+1)
	pm.cities[start] = 0
	matrixIndex := 1
	for tp := range travelPoints {
		pm.cities[tp] = matrixIndex
		matrixIndex++
	}

	pm.pathways = make([][]model.Path, len(pm.cities))
	for i := 0; i < len(pm.cities); i++ {
		pm.pathways[i] = make([]model.Path, len(pm.cities))
	}
	pm.pheromones = make([][]float64, len(pm.cities))
	for i := 0; i < len(pm.cities); i++ {
		pm.pheromones[i] = make([]float64, len(pm.cities))
	}

	return pm
}

func (pm PathMatrix) Pheromone(i, j uint64) float64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.pheromones[pm.cities[i]][pm.cities[j]]
}

func (pm *PathMatrix) UpdatePheromone(i, j uint64, value float64) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.pheromones[pm.cities[i]][pm.cities[j]] += value
	pm.pheromones[pm.cities[j]][pm.cities[i]] += value
}

func (pm *PathMatrix) SetPheromone(i, j uint64, value float64) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.pheromones[pm.cities[i]][pm.cities[j]] = value
	pm.pheromones[pm.cities[j]][pm.cities[i]] = value
}

func (pm PathMatrix) Distance(i, j uint64) float64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.pathways[pm.cities[i]][pm.cities[j]].Len
}

func (pm PathMatrix) Path(i, j uint64) model.Path {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.pathways[pm.cities[i]][pm.cities[j]]
}

func (pm *PathMatrix) SetPath(i, j uint64, path model.Path) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.pathways[pm.cities[i]][pm.cities[j]] = path
	pm.pathways[pm.cities[j]][pm.cities[i]] = path
}
