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
	lastFrost, firstFrost, err := client.GetDatesByTemperature(station, 0)

	var datesToReference float64

	if lastFrost > 0 {
		datesToReference = lastFrost
		getCropsToPlant(lastFrost)
	}
	if firstFrost > 0 && lastFrost < 0 {
		datesToReference = firstFrost
		getCropsToPlant(firstFrost)
	}
	if firstFrost < 0 && firstFrost > -30 {
		fmt.Println("Winter is coming")
		fmt.Println("Garlic and Onions can be planted in the fall and winter for spring harvest")
		datesToReference, _, err = client.GetDatesByTemperature(station, 1)
		fmt.Printf("There are %0.0f days until the last frost \n", datesToReference)
		fmt.Printf("Check back in at %0.0f days", (datesToReference - 45))
	}
	if firstFrost < (-30) {
		fmt.Println("Winter is coming")
		datesToReference, _, err = client.GetDatesByTemperature(station, 1)
		fmt.Printf("There are %0.0f days until the last frost \n", datesToReference)
		fmt.Printf("Check back in at %0.0f", (datesToReference - 45))
	}
	if firstFrost < -60 {
		datesToReference, _, err = client.GetDatesByTemperature(station, 1)
		fmt.Printf("There are %0.0f days until the last frost \n", datesToReference)
		fmt.Printf("Check back in at %0.0f days", (datesToReference - 45))
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
