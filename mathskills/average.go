package mathskill

// calculateAverage calculates the average of the given data.
func CalculateAverage(data []float64) float64 {
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}
