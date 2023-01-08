package engine

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jbattistella/capstone-project/client"
	"github.com/jbattistella/capstone-project/database"
)

type AppResponse struct {
	Msg1       string               `json:"msg1"`
	Msg2       string               `json:"msg2"`
	Msg3       string               `json:"Msg3"`
	Vegetables []database.Vegetable `json:"vegetables"`
}

func Engine(zip string) AppResponse {
	// zipCode := 36525
	zipCode, err := strconv.Atoi(zip)
	if err != nil {
		log.Fatal(err)
	}
	postalInfo, err := client.GetPostalInfo(zipCode)
	if err != nil {
		log.Fatal(err)
	}
	station, err := postalInfo.GetStation()
	if err != nil {
		log.Fatal(err)
	}
	dTFrost, err := client.GetDatesByTemperature(station, 0)

	//create var of type response
	var response AppResponse

	if dTFrost.LastFrost > 0 {
		response = AppResponse{
			Vegetables: getCropsToPlant(dTFrost.LastFrost),
		}

	}
	if dTFrost.FirstFrost > 0 && dTFrost.LastFrost < 0 {
		response = AppResponse{
			Vegetables: getCropsToPlant(dTFrost.FirstFrost),
		}
	}
	if dTFrost.FirstFrost < 0 && dTFrost.FirstFrost > -30 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost

		response = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < (-30) {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		response = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < -60 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		response = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter \n Order you seeds and prepare garden areas for spring planting",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}

	return response
}

func getCropsToPlant(days float64) []database.Vegetable {
	database.ConnectDB()
	var veg []database.Vegetable
	// var v models.Vegetable

	err := database.DB.Model(&veg).Where("dtm < ?", int(days)).Select()
	if err != nil {
		log.Fatal(err)
	}

	return veg
}
