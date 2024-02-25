// Golang program to demonstrate the usage of variables

package main

import "fmt"

func main() {
	fmt.Println("Hello Folks!! This is a program to show a demo on Golang Variables")

	//Using the var keyword with data-type

	var int1 int = 1
	fmt.Println(int1)

	//Using the var keyword without the data-type

	var int2 = 2
	fmt.Println(int2)

	// Using the walrus operator to assign variable

	int3 := 3
	fmt.Println(int3)

}
