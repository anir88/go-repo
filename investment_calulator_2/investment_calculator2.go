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

	investedValueDerived := float64(investedAmount) * math.Pow(1+rateOfInterest/100, float64(durationOfInvestment))

	fmt.Println("The derived investment is:", investedValueDerived)

}
