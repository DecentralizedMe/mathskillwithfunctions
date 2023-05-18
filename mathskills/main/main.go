package main

import (
	"fmt"
	"log"
	"mathskill"
	"os"
)

func main() {
	// Read the file path from command-line argument
	filePath := os.Args[1]

	// Read data from the file
	data, err := mathskill.ReadDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate average
	average := mathskill.CalculateAverage(data)

	// Sort the data
	sortedData := mathskill.SortData(data)

	// Calculate median
	median := mathskill.CalculateMedian(sortedData)

	// Calculate variance
	variance := mathskill.CalculateVariance(data, average)

	// Calculate standard deviation
	standardDeviation := mathskill.CalculateStandardDeviation(variance)

	// Print the results rounded to the nearest number
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", standardDeviation)
}
