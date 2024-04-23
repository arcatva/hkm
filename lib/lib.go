package lib

import (
	"encoding/csv"
	"math"
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

func ReadRow(rowNum int, start int, end int, allData [][]string) ([]float64, error) {
	result := make([]float64, 0)

	for current := start; current <= end; current++ {
		value, err := strconv.ParseFloat(allData[rowNum][current], 32)
		if err != nil {
			return result, nil
		}
		result = append(result, value)
	}

	return result, nil

}

func CalculateMAPE(actual, predicted []float64) float64 {
	if len(actual) != len(predicted) {
		panic("Lengths of actual and predicted slices must be equal")
	}

	var sumPercentageError float64

	for i := 0; i < len(actual); i++ {
		actualValue := actual[i]
		predictedValue := predicted[i]

		// Avoid division by zero
		if actualValue == 0 {
			continue
		}

		// Calculate absolute percentage error
		absPercentageError := math.Abs((actualValue - predictedValue) / actualValue)

		// Sum up absolute percentage errors
		sumPercentageError += absPercentageError
	}

	// Calculate MAPE
	mape := (sumPercentageError / float64(len(actual))) * 100
	return mape
}
