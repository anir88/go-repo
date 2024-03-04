//This is the program to demmonstrate the investment calculator

package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := 1000
	rateOfInterest := 5.5
	numberOfYears := 10

	amountDerived := float64(investmentAmount) * math.Pow((1+rateOfInterest/100), float64(numberOfYears))
	fmt.Print("The amount which would be derived is: ", amountDerived)

}
