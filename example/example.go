package main

import (
	"fmt"
	"github.com/mattn/go-jmardb"
	"log"
)

func main() {
	datas, err := jmardb.NewClient().Area(35.6199114, 139.7512247, -1)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range datas {
		fmt.Println(data.JISX0402 + ": " + data.Code + ", " + data.Name)
		for _, forecast := range data.Forecast {
			fmt.Println("\t" + forecast.Code + ": " + forecast.Name)
		}
	}
}
