package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

type Data struct {
	Sites []string `json:"sites"`
}

const port = "443"

func main() {
	var dateNow = time.Now().Format(time.RFC3339)
	file, err := ioutil.ReadFile("sites.json")
	if err != nil {
		fmt.Println("Error open sites.json files")
		panic(err)
	}
	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error copying data to variable")
		panic(err)
	}
	fmt.Printf("URL\t Certificate Date\t Days until expiration\t\n")
	for _, url := range data.Sites {
		expirationDate := getExpirationDate(url, port)
		if expirationDate == "" {
			fmt.Printf("%v\t 0\t 0\t\n", url)
		} else {
			d, err := getDurationFromDates(dateNow, expirationDate)
			if err != nil {
				panic(err)
			}
			duration := d.Hours() / 24
			days := int(math.Round(duration))
			fmt.Printf("%v\t %v\t %v\t\n", url, expirationDate, days)
		}

	}
	os.Exit(1)
}
