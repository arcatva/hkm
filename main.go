package main

import (
	"fmt"
	"github.com/sajari/regression"
	"hkm/lib"
)

func main() {
	allData, _ := lib.ReadAllData("He_Kelly_Manela_Factors_And_Test_Assets_monthly.csv")
	slopeShockList := make([]float64, 0)
	slopeMarketList := make([]float64, 0)
	for i := 9; i <= 33; i++ {
		coli, _ := lib.ReadColumn(i, allData)
		col3, _ := lib.ReadColumn(3, allData)
		col5, _ := lib.ReadColumn(5, allData)

		r := new(regression.Regression)
		r.SetObserved(fmt.Sprintf("FF25-%v", i-8))
		r.SetVar(0, "intermediary_capital_risk_factor")
		r.SetVar(1, "mkt_rf")
		for j, _ := range col3 {
			r.Train(
				regression.DataPoint(coli[j], []float64{col3[j], col5[j]}),
			)
		}

		r.Run()
		fmt.Printf("FF25-%v\n", i-8)
		fmt.Printf("R2:%v\n", r.R2)
		slopeShock := r.GetCoeffs()[1]
		slopeShockList = append(slopeShockList, slopeShock)
		slopeMarket := r.GetCoeffs()[2]
		slopeMarketList = append(slopeMarketList, slopeMarket)
		fmt.Printf("slopeShock: %v\nslopeMarket: %v\n", slopeShock, slopeMarket)
		fmt.Printf("=======================\n")

	}
	fmt.Printf("Finished beta couting...\n")
	fmt.Printf("=======================\n")
	lib.CalculateAverage(1, 516, 10, 34, allData, slopeShockList, slopeMarketList)
	lib.CalculateAverage(1, 516, 35, 54, allData, slopeShockList, slopeMarketList)
}
