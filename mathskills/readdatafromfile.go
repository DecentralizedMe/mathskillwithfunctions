package mathskill

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// readDataFromFile reads the data from the given file path and returns a slice of numbers.
func ReadDataFromFile(filePath string) ([]float64, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var data []float64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	return data, nil
}
