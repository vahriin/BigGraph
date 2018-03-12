package types

// AdjList is (TODO: Add doc)
type AdjList struct {
	AL    map[uint64][]uint64
	Nodes map[uint64]GeneralCoords
}

// DropExcessPoints delete points that located on a straight line between two other points
func (al *AdjList) DropExcessPoints() int {
	counter := 0

	for point, incidentNodes := range al.AL {

		if len(incidentNodes) == 2 {
			// sum of distances between current point and those incident to it should not be much greater than
			// distance between two incident points of current point according to the rule of the triangle
			if Distance(al.Nodes[incidentNodes[0]].Euclid, al.Nodes[point].Euclid)+
				Distance(al.Nodes[incidentNodes[1]].Euclid, al.Nodes[point].Euclid)-
				Distance(al.Nodes[incidentNodes[0]].Euclid, al.Nodes[incidentNodes[1]].Euclid) < 1E-3 {

				delete(al.AL, point)
				counter++
			}
		}
	}
	return counter
}
