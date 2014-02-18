package main

import (
	"fmt"
	"github.com/mattn/go-jmardb"
	"log"
)

func main() {
	client := jmardb.NewClient()

	areas, err := client.Area(35.6199114, 139.7512247, -1)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range areas {
		fmt.Println(area.JISX0402 + ": " + area.Code + ", " + area.Name)
		for _, forecast := range area.Forecast {
			fmt.Println("\t" + forecast.Code + ": " + forecast.Name)
		}
	}
}
