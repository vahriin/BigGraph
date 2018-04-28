package algorithm

import (
	"math"

	"github.com/vahriin/BigGraph/lib/coordinates"
)

func getHF() func(coordinates.EuclidCoords, coordinates.EuclidCoords) float64 {
	return l2Dist
}

func l1Dist(begin, end coordinates.EuclidCoords) float64 {
	return math.Abs(begin.X-end.X) + math.Abs(begin.Y-end.Y)
}

func lInfDist(begin, end coordinates.EuclidCoords) float64 {
	return math.Max(math.Abs(begin.X-end.X), math.Abs(begin.Y-end.Y))
}

func l2Dist(begin, end coordinates.EuclidCoords) float64 {
	return coordinates.Distance(begin, end)
}
