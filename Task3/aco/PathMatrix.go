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

	pm.pathways = make([][]model.Path, len(pm.cities)-1)
	for i := 0; i < len(pm.cities)-1; i++ {
		pm.pathways[i] = make([]model.Path, len(pm.cities)-1-i)
	}
	pm.pheromones = make([][]float64, len(pm.cities)-1)
	for i := 0; i < len(pm.cities)-1; i++ {
		pm.pheromones[i] = make([]float64, len(pm.cities)-1-i)
	}

	return pm
}

func (pm PathMatrix) Pheromone(i, j uint64) float64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	if i == j {
		return 0
	}

	iN, jN := pm.normalIndexes(i, j)
	return pm.pheromones[iN][jN]
}

func (pm *PathMatrix) UpdatePheromone(i, j uint64, value float64) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if i == j {
		return
	}

	iN, jN := pm.normalIndexes(i, j)

	pm.pheromones[iN][jN] += value
}

func (pm *PathMatrix) SetPheromone(i, j uint64, value float64) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if i == j {
		return
	}

	iN, jN := pm.normalIndexes(i, j)

	pm.pheromones[iN][jN] = value

}

func (pm PathMatrix) Distance(i, j uint64) float64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	if i == j {
		return 0
	}

	iN, jN := pm.normalIndexes(i, j)

	return pm.pathways[iN][jN].Len
}

func (pm PathMatrix) Path(i, j uint64) model.Path {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	if i == j {
		rp := model.Path{
			Len:    0,
			Points: make([]uint64, 1),
		}
		rp.Points[0] = i
		return rp
	}

	iN, jN := pm.normalIndexes(i, j)

	return pm.pathways[iN][jN]
}

func (pm *PathMatrix) SetPath(i, j uint64, path model.Path) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if i == j {
		return
	}

	iN, jN := pm.normalIndexes(i, j)

	pm.pathways[iN][jN] = path
}

func (pm PathMatrix) normalIndexes(i, j uint64) (iN, jN int) {
	if pm.cities[j] < pm.cities[i] {
		i, j = j, i
	}

	return pm.cities[i], pm.cities[j] - pm.cities[i] - 1
}
