package main

import (
	"fmt"
	"log"

	"github.com/jbattistella/capstone-project/client"
	"github.com/jbattistella/capstone-project/database"
)

func main() {

	zipCode := 36525
	postalInfo, err := client.GetPostalInfo(zipCode)
	if err != nil {
		log.Fatal(err)
	}
	station, err := postalInfo.GetStation()
	if err != nil {
		log.Fatal(err)
	}
	dTFrost, err := client.GetDatesByTemperature(station, 0)

	// var datesToReference float64

	if dTFrost.LastFrost > 0 {
		// datesToReference = dTFrost.LastFrost
		getCropsToPlant(dTFrost.LastFrost)
	}
	if dTFrost.FirstFrost > 0 && dTFrost.LastFrost < 0 {
		// datesToReference = dTFrost.FirstFrost
		getCropsToPlant(dTFrost.FirstFrost)
	}
	if dTFrost.FirstFrost < 0 && dTFrost.FirstFrost > -30 {
		fmt.Println("Winter is coming")
		fmt.Println("Garlic and Onions can be planted in the fall and winter for spring harvest")
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		firstFrost := dTFrost.FirstFrost
		fmt.Printf("There are %0.0f days until the last frost \n", firstFrost)
		fmt.Printf("Check back in at %0.0f days", (firstFrost - 45))
	}
	if dTFrost.FirstFrost < (-30) {
		fmt.Println("Winter is coming")
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		firstFrost := dTFrost.FirstFrost
		fmt.Printf("There are %0.0f days until the last frost \n", firstFrost)
		fmt.Printf("Check back in at %0.0f days", (firstFrost - 45))
	}
	if dTFrost.FirstFrost < -60 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		firstFrost := dTFrost.FirstFrost
		fmt.Printf("There are %0.0f days until the last frost \n", firstFrost)
		fmt.Printf("Check back in at %0.0f days", (firstFrost - 45))
	}
}

func getCropsToPlant(days float64) {
	database.ConnectDB()
	var veg []database.Vegetable
	// var v models.Vegetable

	err := database.DB.Model(&veg).Where("dtm < ?", int(days)).Select()
	if err != nil {
		log.Fatal(err)
	}

	var vegNames []string

	for _, v := range veg {
		vegNames = append(vegNames, v.CommonName)
		fmt.Println(v.CommonName)
	}

}
