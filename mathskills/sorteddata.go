package mathskill

import "sort"

// sortData sorts the given data in ascending order.
func SortData(data []float64) []float64 {
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	return sortedData
}
