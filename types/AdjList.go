package types

import (
	"math"
)

// AdjList contains the AL - adjacency list like a map where key is point id and value is incident points id array
// and Nodes - map of nodes, where key is point id
type AdjList struct {
	AL    map[uint64][]uint64
	Nodes map[uint64]GeneralCoords
}

func (al AdjList) onLine(left, center, right uint64) bool {
	areaOfTriangle := math.Abs(
		// calc the area of triangle on those three points
		(al.Nodes[left].Euclid.X-al.Nodes[right].Euclid.X)*(al.Nodes[center].Euclid.Y-al.Nodes[right].Euclid.Y)-
			(al.Nodes[center].Euclid.X-al.Nodes[right].Euclid.X)*(al.Nodes[left].Euclid.Y-al.Nodes[right].Euclid.Y)) / 2

	// if the area is negligible small then points are on the same straight line
	if areaOfTriangle < 1E-1 {
		return true
	}
	return false
}

// DropExcessPoints delete points that located on a straight line between two other points
func (al *AdjList) DropExcessPoints() int {
	counter := 0

	for point, incidentNodes := range al.AL {

		if len(incidentNodes) == 2 {
			// sum of distances between current point and those incident to it should not be much greater than
			// distance between two incident points of current point according to the rule of the triangle
			if al.onLine(incidentNodes[0], point, incidentNodes[1]) {

				al.AL[incidentNodes[0]][linearSearch(al.AL[incidentNodes[0]], point)] = incidentNodes[1]
				al.AL[incidentNodes[1]][linearSearch(al.AL[incidentNodes[1]], point)] = incidentNodes[0]

				delete(al.AL, point)
				delete(al.Nodes, point)
				counter++
			}
		}
	}
	return counter
}

func linearSearch(array []uint64, target uint64) (index int) {
	for i, p := range array {
		if p == target {
			return i
		}
	}
	return -1
}
