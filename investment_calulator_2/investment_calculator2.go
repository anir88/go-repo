//Golang script to create a compund interest calculator

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello!! I am an investment calulator")
	investedAmount := 100000
	rateOfInterest := 5.5
	durationOfInvestment := 10
	const minInflationRate = 2.5
	const maxInflationRate = 8.5

	investedValueDerived := float64(investedAmount) * math.Pow(1+rateOfInterest/100, float64(durationOfInvestment))
	fmt.Println("The ideal derived investment is:", investedValueDerived)
	bestInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+minInflationRate/100, float64(durationOfInvestment))
	fmt.Println("The best case derived value is:", bestInflationCorrectedInvestedDerivedValue)
	worstInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+maxInflationRate/100, float64(durationOfInvestment))
	fmt.Println("The worst case derived value is:", worstInflationCorrectedInvestedDerivedValue)

}
