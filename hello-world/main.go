//First Golang program!! All the best!!

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello world!! I am new golang program")
	fmt.Printf("Hello Folks!! I am a golang program with arguments.\nThe argument is: %v\n", args)
	fmt.Printf("Hello folks!! The output with only the supplied arguments are as follows:\n %v\n", args[1:])

}
