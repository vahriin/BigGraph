package coordinates

import "math"

// Distance calc the distance between two point on Euclid metric
func Distance(one, two EuclidCoords) float64 {
	return math.Sqrt((one.X-two.X)*(one.X-two.X) + (one.Y-two.Y)*(one.Y-two.Y))
}
