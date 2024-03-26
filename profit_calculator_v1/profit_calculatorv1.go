package main

import "fmt"

func main() {
	fmt.Println("Hello!! I am a simple profit calculator")

	//Accepting the user inputs

	var costPrice, sellingPrice, salesTax, revenue, taxedAmountOnrevenue, profitPercentage float64

	fmt.Print("Please enter the costPrice: ")
	fmt.Scan(&costPrice)

	fmt.Print("Please enter the selling price: ")
	fmt.Scan(&sellingPrice)

	fmt.Print("Please enter the sales tax as applicable in your region/location: ")
	fmt.Scan(&salesTax)

	revenue = sellingPrice - costPrice
	fmt.Println("The revenue generated is: ", revenue)

	taxedAmountOnrevenue = revenue * (salesTax / 100)
	fmt.Println("The taxed amount on revenue is: ", taxedAmountOnrevenue)

	profitPercentage = (revenue - taxedAmountOnrevenue) / 100
	fmt.Println("The profit percentage is: ", profitPercentage)

}
