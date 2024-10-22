package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No arguments found.\nThe usage is:\n<program_name> <argument>")
		os.Exit(1)
	} else {
		fmt.Printf("Hello folks!! The output with only the supplied arguments are as follows:\n %v\n", args[1:])
	}
}
