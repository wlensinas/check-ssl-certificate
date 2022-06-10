package main

import (
	"fmt"
	"os"
)

func main() {
	var url string = "www.google.com.ar"
	var port string = "443"
	expiration_date, err := getExpirationDate(url, port)
	if err != nil {
		fmt.Println("Problem with ssl certificate")
		os.Exit(1)
	}
	fmt.Println(expiration_date)
}
