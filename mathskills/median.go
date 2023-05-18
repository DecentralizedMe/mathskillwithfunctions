package mathskill

// calculateMedian calculates the median of the given sorted data.
func CalculateMedian(sortedData []float64) float64 {
	length := len(sortedData)
	middle := length / 2
	if length%2 == 0 {
		return (sortedData[middle-1] + sortedData[middle]) / 2.0
	}
	return sortedData[middle]
}
