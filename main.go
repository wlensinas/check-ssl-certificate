package main

import (
	"fmt"
	"os"
)

const url = "www.google.com.ar"
const port = "443"

func main() {
	expirationDate, err := getExpirationDate(url, port)

	if err != nil {
		fmt.Println(fmt.Sprintf("Problem with ssl certificate - err: %v", err))
		os.Exit(1)
	}

	fmt.Println(expirationDate)
}
