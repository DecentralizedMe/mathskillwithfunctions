package mathskill

import "math"

// calculateStandardDeviation calculates the standard deviation from the given variance.
func CalculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
