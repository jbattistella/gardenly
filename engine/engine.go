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
	Msg3       string               `json:"msg3"`
	Masg4      error                `json:"Msg4"`
	Vegetables []database.Vegetable `json:"vegetables"`
}

func Engine(zip string) (AppResponse, error) {
	var a AppResponse

	zipCode, err := strconv.Atoi(zip)
	if err != nil {
		return AppResponse{}, err
	}

	postalInfo, err := client.GetPostalInfo(zipCode)
	if err != nil {
		return AppResponse{}, err
	}

	station, err := postalInfo.GetStation()
	if err != nil {
		return AppResponse{}, err
	}
	dTFrost, err := client.GetDatesByTemperature(station, 0)

	fmt.Println(dTFrost.LastFrost) //74

	if dTFrost.LastFrost > 0 {
		a = AppResponse{
			Msg1: fmt.Sprintf("%d days until the last frost.", int(dTFrost.LastFrost)),
		}
	}
	if dTFrost.LastFrost > 45 {
		daysToSeeding := int(dTFrost.LastFrost) - 45
		a = AppResponse{
			Msg1: fmt.Sprintf("%d days until the last frost. Check back in %d days.", int(dTFrost.LastFrost), daysToSeeding),
		}
	}
	if dTFrost.LastFrost < 45 && dTFrost.LastFrost > 0 {
		a = AppResponse{
			Msg1:       fmt.Sprintf("%d days until the last frost.", int(dTFrost.LastFrost)),
			Vegetables: getCropsToPlant(dTFrost.FirstFrost),
		}
		fmt.Println(getCropsToPlant(dTFrost.FirstFrost))
		fmt.Println(dTFrost.FirstFrost)
	}
	if dTFrost.LastFrost < 0 {
		a = AppResponse{
			Msg1:       fmt.Sprintf("%d days until the last frost.", int(dTFrost.FirstFrost)),
			Vegetables: getCropsToPlant(dTFrost.LastFrost),
		}
	}
	if dTFrost.LastFrost < 65 && dTFrost.LastFrost > 45 {
		a = AppResponse{
			Msg1: "Prepare for spring!",
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
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < (-30) {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		a = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	if dTFrost.FirstFrost < -60 {
		dTFrost, err = client.GetDatesByTemperature(station, 1)
		lastFrost := dTFrost.LastFrost
		a = AppResponse{
			Msg1: "Winter is coming",
			Msg2: "Garlic and Onions can be planted in the fall and winter. Order you seeds and prepare garden areas for spring planting",
			Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
		}
	}
	return a, nil
}

func getCropsToPlant(days float64) []database.Vegetable {
	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
	}
	var vegetables []database.Vegetable

	_ = DB.Where("dtm < ?", days).Find(&vegetables)

	return vegetables
}
