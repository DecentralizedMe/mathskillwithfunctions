package mathskill

// calculateVariance calculates the variance of the given data, given the average.
func CalculateVariance(data []float64, average float64) float64 {
	sumSquaredDiffs := 0.0
	for _, num := range data {
		diff := num - average
		sumSquaredDiffs += diff * diff
	}
	return sumSquaredDiffs / float64(len(data))
}
