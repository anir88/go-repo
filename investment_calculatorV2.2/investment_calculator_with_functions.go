// Golang program to get the investment values using functions

package main

import (
	"fmt"
	"math"
)

var investedAmount, rateOfInterest, durationOfInvestment float64

const minInflationRate = 2.5
const maxInflationRate = 8.5

func investmentCalculation(investedAmount, rateOfInterest, durationOfInvestment float64) (float64, float64, float64) {
	investedValueDerived := float64(investedAmount) * math.Pow(1+rateOfInterest/100, float64(durationOfInvestment))
	bestInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+minInflationRate/100, float64(durationOfInvestment))
	worstInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+maxInflationRate/100, float64(durationOfInvestment))
	return investedValueDerived, bestInflationCorrectedInvestedDerivedValue, worstInflationCorrectedInvestedDerivedValue

}

func main() {
	//var idealFutureValue, actualFutureValueWithBestInflationCase, actualFutureValueWithWorstInflationCase float64
	fmt.Println("Hello All!! Welcome to investment calculatorv2.2")
	// Getting the amount to be invested from the user
	fmt.Print("Please enter the desired amount to invest: ")
	fmt.Scan(&investedAmount)

	// Getting the rate of interest as desired by the user
	fmt.Print("Please enter the rate of interest: ")
	fmt.Scan(&rateOfInterest)

	// Getting the investment duration from the user

	fmt.Print("Please enter the desired duration of investment: ")
	fmt.Scan(&durationOfInvestment)
	idealFutureValue, actualFutureValueWithBestInflationCase, actualFutureValueWithWorstInflationCase := investmentCalculation(investedAmount, rateOfInterest, durationOfInvestment)
	fmt.Printf("The ideal derived valyue should be: %0.2f\n", idealFutureValue)
	fmt.Printf("The best case derived value as per inflation is: %0.2f\n", actualFutureValueWithBestInflationCase)
	fmt.Printf("The worst case derived value as per inflation is: %0.2f\n", actualFutureValueWithWorstInflationCase)

}
