package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//client := &http.Client{}

	resp, err := http.DefaultClient.Get("https://academy.golang-ninja.com/")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.StatusCode)
}
