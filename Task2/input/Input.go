package input

import (
	"github.com/vahriin/BigGraph/lib/csv"
	"github.com/vahriin/BigGraph/lib/model"
)

func Input(directory string) (al model.AdjList, destPoints map[uint64]struct{}, startPoint uint64) {
	al = model.NewAdjList(directory+"adjacency_list.csv", directory+"nodes_list.csv")
	dest := csv.ReadPointsID(directory + "destination_points.csv")

	destPoints = make(map[uint64]struct{}, 10)
	for _, dp := range dest {
		destPoints[dp] = struct{}{}
	}

	start := ReadPoint(directory + "point.xml")

	al.AddPoint(start)

	return al, destPoints, 0
}
