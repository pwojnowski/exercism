package darts

import "math"

func Score(x, y float64) int {
	var d = math.Sqrt(x*x + y*y)
	var points = 0

	switch {
	case 10 < d:
		points = 0
	case 5 < d:
		points = 1
	case 1 < d:
		points = 5
	case d <= 1:
		points = 10
	}

	return points
}
