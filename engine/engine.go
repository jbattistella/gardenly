package engine

import (
	"errors"
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
	Masg4      error                `json:"msg4"`
	Vegetables []database.Vegetable `json:"vegetables"`
}

func Engine(zip string) (AppResponse, error) {
	var a AppResponse

	var dTFrost client.DaysFromFrost
	var station string

	//non test code
	if zip[0] != 't' {
		zipCode, err := strconv.Atoi(zip)
		if err != nil {
			err = errors.New("re-enter zip code as a five digit integer")
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
		dTFrost, err = client.GetDatesByTemperature(station, 0)
	}
	//test cases and vars
	var LastFrostPlus float64
	if zip[0] == 't' {
		switch zip[1] {
		case '1':
			dTFrost.LastFrost = 60
			dTFrost.FirstFrost = 240
		case '2':
			dTFrost.LastFrost = 20
			dTFrost.FirstFrost = 200
		case '3':
			dTFrost.LastFrost = -10
			dTFrost.FirstFrost = 170
		case '4':
			dTFrost.LastFrost = -110
			dTFrost.FirstFrost = 70
		case '5':
			dTFrost.LastFrost = -170
			dTFrost.FirstFrost = 30
		case '6':
			dTFrost.LastFrost = -220
			dTFrost.FirstFrost = -20
			LastFrostPlus = 160
		case '7':
			dTFrost.LastFrost = -250
			dTFrost.FirstFrost = -50
			LastFrostPlus = 130
		case '8':
			dTFrost.LastFrost = 100
			dTFrost.FirstFrost = 280
		}
	}

	daysToSeeding := int(dTFrost.LastFrost) - 45

	if dTFrost.LastFrost > 0 {
		switch dTFrost.LastFrost < 45 {
		case true:
			a = AppResponse{
				Msg1:       fmt.Sprintf("%d days until the last frost.", int(dTFrost.LastFrost)),
				Vegetables: getCropsToPlant(dTFrost.FirstFrost),
			}
		case false:
			a.Msg1 = fmt.Sprintf("%d days until the last frost. Check back in %d days.", int(dTFrost.LastFrost), daysToSeeding)
		}
	}

	//Last frost
	if dTFrost.FirstFrost > 0 && dTFrost.LastFrost < 1 {
		a = AppResponse{
			Msg1:       "We are into the growing season. Plant slow and steady.",
			Vegetables: getCropsToPlant(dTFrost.FirstFrost),
		}
	}

	if dTFrost.FirstFrost > 45 && dTFrost.FirstFrost < 95 {
		a = AppResponse{
			Msg1:       fmt.Sprintf("%d days until the first frost.", int(dTFrost.FirstFrost)),
			Vegetables: getCropsToPlant(dTFrost.FirstFrost - 15),
		}
	}

	if dTFrost.FirstFrost < 45 {
		a.Msg1 = fmt.Sprintf("%d days until the first frost. \n\n", int(dTFrost.FirstFrost))
		a.Msg2 = fmt.Sprintln("Prepare your garden for winter.")
		a.Vegetables = nil
	}

	//First Frost
	if dTFrost.FirstFrost < 0 {
		switch {
		case dTFrost.FirstFrost > -30:
			if zip[0] == 't' {
				a = AppResponse{
					Msg1: "Winter is coming!\n",
					Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
					Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", LastFrostPlus, (LastFrostPlus - 45)),
				}
				return a, nil
			}

			dTFrost, err := client.GetDatesByTemperature(station, 1)
			if err != nil {
				return AppResponse{}, err
			}
			lastFrost := dTFrost.LastFrost

			a = AppResponse{
				Msg1: "Winter is coming!\n",
				Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
				Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
			}
		case dTFrost.FirstFrost < -30:
			if zip[0] == 't' {
				a = AppResponse{
					Msg1: "Winter is coming!\n",
					Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
					Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", LastFrostPlus, (LastFrostPlus - 45)),
				}
				return a, nil
			}

			dTFrost, err := client.GetDatesByTemperature(station, 1)
			if err != nil {
				return AppResponse{}, err
			}
			lastFrost := dTFrost.LastFrost

			a = AppResponse{
				Msg1: "Winter is coming!\n",
				Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
				Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
			}

		case dTFrost.FirstFrost < -60:
			if zip[0] == 't' {
				a = AppResponse{
					Msg1: "Winter is coming!\n",
					Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
					Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", LastFrostPlus, (LastFrostPlus - 45)),
				}
				return a, nil
			}

			dTFrost, err := client.GetDatesByTemperature(station, 1)
			if err != nil {
				return AppResponse{}, err
			}
			lastFrost := dTFrost.LastFrost

			a = AppResponse{
				Msg1: "Winter is coming!\n",
				Msg2: "Garlic and Onions can be planted in the fall and winter for spring harvest \n\n",
				Msg3: fmt.Sprintf("There are %0.0f days until the last frost. Check back in at %0.0f days", lastFrost, (lastFrost - 45)),
			}
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
