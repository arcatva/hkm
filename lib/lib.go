package lib

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
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
			value = 0
		}
		result = append(result, value)
	}

	return result, nil

}

func CalculateAverage(rowStart, rowEnd, columnStart, columnEnd int, allData [][]string, slopeShockList, slopeMarketList []float64) {
	allR2 := make([]float64, rowEnd-rowStart+1)
	allSlope := make([]float64, rowEnd-rowStart+1)
	allSlopeShock := make([]float64, rowEnd-rowStart+1)
	allSlopeMarket := make([]float64, rowEnd-rowStart+1)

	for i := rowStart; i <= rowEnd; i++ {
		row, _ := ReadRow(i, columnStart-1, columnEnd-1, allData)
		r := new(regression.Regression)
		r.SetObserved(allData[i][0])
		r.SetVar(0, "slopeShock")
		r.SetVar(1, "slopeMarket")

		for j := 0; j <= columnEnd-columnStart; j++ {
			r.Train(
				regression.DataPoint(row[j], []float64{slopeShockList[j], slopeMarketList[j]}),
			)
		}
		r.Run()

		fmt.Printf("Row %v 's slope: %v\n", i, r.GetCoeffs()[0])
		fmt.Printf("Row %v 's slopeShock: %v\n", i, r.GetCoeffs()[1])
		fmt.Printf("Row %v 's slopeMarket: %v\n", i, r.GetCoeffs()[2])
		fmt.Printf("Row %v 's R2: %v\n", i, r.R2)
		allR2[i-1] = r.R2
		allSlope[i-1] = r.GetCoeffs()[0]
		allSlopeShock[i-1] = r.GetCoeffs()[1]
		allSlopeMarket[i-1] = r.GetCoeffs()[2]
	}

	var allR2Aggregated float64

	for _, v := range allR2 {
		allR2Aggregated += v
	}
	averageR2 := allR2Aggregated / float64(rowEnd-rowStart+1)
	fmt.Printf("=======================\n")
	fmt.Println("Average R2 is:")
	fmt.Println(averageR2)

	var allSlopeAggregated float64

	for _, v := range allSlope {
		allSlopeAggregated += v
	}
	averageSlope := allSlopeAggregated / float64(rowEnd-rowStart+1)
	fmt.Printf("=======================\n")
	fmt.Println("Average Slope is:")
	fmt.Println(averageSlope)

	var allSlopeShockAggregated float64

	for _, v := range allSlopeShock {
		allSlopeShockAggregated += v
	}
	averageSlopeShock := allSlopeShockAggregated / float64(rowEnd-rowStart+1)
	fmt.Printf("=======================\n")
	fmt.Println("Average SlopeShock is:")
	fmt.Println(averageSlopeShock)

	var allSlopeMarketAggregated float64

	for _, v := range allSlopeMarket {
		allSlopeMarketAggregated += v
	}
	averageSlopeMarket := allSlopeMarketAggregated / float64(rowEnd-rowStart+1)
	fmt.Printf("=======================\n")
	fmt.Println("Average SlopeMarket is:")
	fmt.Println(averageSlopeMarket)

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
