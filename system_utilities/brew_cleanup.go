//Golang Program to cleanup homebrew files as per a cron job

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("starting the homebrew cleanup job")
	cmd := exec.Command("brew", "cleanup")
	out1, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out1))

}
