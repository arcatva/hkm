package lib

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadAllData(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return make([][]string, 0), err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	allData, err := reader.ReadAll()
	if err != nil {
		return make([][]string, 0), err
	}
	return allData, nil
}

func ReadColumn(column int, allData [][]string) ([]float64, error) {
	result := make([]float64, 0)
	for i, row := range allData {
		if i == 0 {
			continue
		}

		value, err := strconv.ParseFloat(row[column], 32)
		if err != nil {
			return result, nil
		}
		result = append(result, value)
	}
	return result, nil
}
