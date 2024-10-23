//Golang program to invoke http-get request from external url

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Data struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
}

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

	if response1.StatusCode != 200 {
		fmt.Printf("Invalid Status code:%d\n", response1.StatusCode)
	}

	var data Data

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed.\nThe fistname is:%v\nThe lastname is:%v\nThe DoB is:%v\nThe gender of %v is:%v\n", data.FirstName, data.LastName, data.Birthday, data.Birthday, data.Gender)

}
