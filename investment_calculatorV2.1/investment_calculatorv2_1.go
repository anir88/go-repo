//Golang script to create a compund interest calculator and fetch data using user inputs

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello!! I am an investment calulator")
	// Getting the amount to be invested from the user
	var investedAmount float64
	fmt.Print("Please enter the desired amount to invest: ")
	fmt.Scan(&investedAmount)

	// Getting the rate of interest as desired by the user
	var rateOfInterest float64
	fmt.Print("PLease enter the rate of interest: ")
	fmt.Scan(&rateOfInterest)

	// Getting the investment suration from the user
	var durationOfInvestment float64
	fmt.Print("Please enter the desired duration of investment: ")
	fmt.Scan(&durationOfInvestment)
	//durationOfInvestment := 10
	const minInflationRate = 2.5
	const maxInflationRate = 8.5

	investedValueDerived := float64(investedAmount) * math.Pow(1+rateOfInterest/100, float64(durationOfInvestment))
	fmt.Println("The ideal derived investment is:", investedValueDerived)
	bestInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+minInflationRate/100, float64(durationOfInvestment))
	fmt.Println("The best case derived value is:", bestInflationCorrectedInvestedDerivedValue)
	worstInflationCorrectedInvestedDerivedValue := investedValueDerived / math.Pow(1+maxInflationRate/100, float64(durationOfInvestment))
	fmt.Println("The worst case derived value is:", worstInflationCorrectedInvestedDerivedValue)

}
