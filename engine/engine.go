package engine

import (
	"fmt"
	"log"

	"github.com/jbattistella/capstone-project/client"
	"github.com/jbattistella/capstone-project/database"
)

type AppResponse struct {
	Msg1       string
	Msg2       string
	Msg3       string
	Vegetables []database.Vegetable
}

func (a AppResponse) Engine() {
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

	if dTFrost.LastFrost > 0 {
		a = AppResponse{
			Vegetables: getCropsToPlant(dTFrost.LastFrost),
		}
	}
	if dTFrost.FirstFrost > 0 && dTFrost.LastFrost < 0 {
		a = AppResponse{
			Vegetables: getCropsToPlant(dTFrost.FirstFrost),
		}
	}
	if dTFrost.FirstFrost < 0 && dTFrost.FirstFrost > -30 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost

		a = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < (-30) {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		a = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < -60 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		a = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter \n Order you seeds and prepare garden areas for spring planting",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost \n Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
}

func getCropsToPlant(days float64) []database.Vegetable {
	if err := database.ConnectDB(); err != nil {
		log.Fatal("error connecting to DB")
	}
	var veg []database.Vegetable

	if err := database.DB.Model(&veg).Where("dtm < ?", int(days)).Select(); err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close()

	return veg
}
