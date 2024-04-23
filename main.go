package main

import (
	"fmt"
	"github.com/sajari/regression"
	"hkm/lib"
)

func main() {
	allData, _ := lib.ReadAllData("He_Kelly_Manela_Factors_And_Test_Assets_monthly.csv")

	col3, _ := lib.ReadColumn(3, allData)
	col5, _ := lib.ReadColumn(5, allData)
	col9, _ := lib.ReadColumn(9, allData)

	r := new(regression.Regression)
	r.SetObserved("FF25_01")
	r.SetVar(0, "intermediary_capital_risk_factor")
	r.SetVar(1, "mkt_rf")
	for i, _ := range col3 {
		r.Train(
			regression.DataPoint(col9[i], []float64{col3[i], col5[i]}),
		)
	}

	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

}
