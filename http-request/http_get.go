//Golang program to invoke http-get request from external url

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Welcome to the HTTP GET program!!")
	args := os.Args

	if len(args) < 2 {
		fmt.Println("URL missing. The correct usage is <program><url>")
		os.Exit(1)

	}

	_, err1 := url.ParseRequestURI(args[1])

	if err1 != nil {
		fmt.Printf("The specified link:%s is not a valid url\n", args[1])
	}

	response1, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response1.Body.Close()

	body, err := io.ReadAll(response1.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The HTTP code is: %d\nThe response body is:%s\n", response1.StatusCode, body)

}
