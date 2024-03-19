//This is the program to demmonstrate the investment calculator

package main

import (
	"fmt"
	"math"
)

func main() {
	var amountInvested int
	var rateOfInterest float64
	var durationOfInvestment int
	fmt.Println("Please enter the amount to invest:")
	fmt.Scan(&amountInvested)

	fmt.Println("Please enter the rate of interest:")
	fmt.Scan(&rateOfInterest)

	fmt.Println("Please enter the time of investment:")
	fmt.Scan(&durationOfInvestment)

	amountDerived := float64(amountInvested) + math.Pow((1+(rateOfInterest/100)), float64(durationOfInvestment))

	fmt.Println("The derived amount is: ", amountDerived)

}
