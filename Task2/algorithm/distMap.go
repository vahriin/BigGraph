package algorithm

import "sync"

type distMap struct {
	mutex     *sync.RWMutex
	distances map[uint64]float64
}

func (dm distMap) distance(id uint64) float64 {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()

	return dm.distances[id]
}

func (dm *distMap) setDistance(id uint64, d float64) bool {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	if dm.distances[id] > d {
		dm.distances[id] = d
		return true
	}
	return false
}
