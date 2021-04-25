package go_helper

import (
	"math"
)

const float64EqualityThreshold = 1e-9

func Float64CompareAlmostEqual(f1 float64, f2 float64) bool {
	return math.Abs(f1-f2) <= float64EqualityThreshold
}
