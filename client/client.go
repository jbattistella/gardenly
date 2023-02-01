package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// get hardiness zone
func GetZoneByZipcode(zipcode string) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("https://plant-hardiness-zone.p.rapidapi.com/zipcodes/%s", zipcode)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", "plant-hardiness-zone.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	type HardnessRes struct {
		Hardiness string `json:"hardiness_zone"`
		Zipcode   string `json:"zipcode"`
	}

	var hr HardnessRes

	json.NewDecoder(res.Body).Decode(&hr)

	log.Printf("zipcode:%s hardiness zone:%s", zipcode, hr.Hardiness)
}

func GetPostalInfo(zip int) (PostalCodeInfo, error) {
	var url string
	//handles postal codes that begin with 0
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PostalCodeInfo{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return PostalCodeInfo{}, err
	}
	defer res.Body.Close()
	var info PostalCodeInfo

	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		return PostalCodeInfo{}, err
	}
	if info.Country == "" {
		return PostalCodeInfo{}, errors.New("invalid zip code")

	}

	return info, nil

}

func (pci *PostalCodeInfo) GetStation() (string, error) {

	latInt, _ := strconv.ParseFloat(pci.Places[0].Latitude, 64)

	if latInt < 31.0 {
		return "", errors.New("zipcode currently not supported")
	}

	url := fmt.Sprintf("https://api.farmsense.net/v1/frostdates/stations/?lat=%s&lon=%s", pci.Places[0].Latitude, pci.Places[0].Longitude)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var info StationInfo
	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		log.Fatal(err)
	}
	station := info[0].ID

	return station, nil
}

func GetDatesByTemperature(id string, y int) (DaysFromFrost, error) {

	url1 := fmt.Sprintf("https://api.farmsense.net/v1/frostdates/probabilities/?station=%s&season=%s", id, "1")
	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var springDates FrostDates
	if err = json.NewDecoder(res.Body).Decode(&springDates); err != nil {
		log.Fatal(err)
	}

	url2 := fmt.Sprintf("https://api.farmsense.net/v1/frostdates/probabilities/?station=%s&season=%s", id, "2")
	req2, err := http.NewRequest("GET", url2, nil)
	if err != nil {
		log.Fatal(err)
	}

	res2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer res2.Body.Close()

	var fallDates FrostDates
	if err = json.NewDecoder(res2.Body).Decode(&fallDates); err != nil {
		log.Fatal(err)
	}

	if springDates[0].Prob90 == "0000" || fallDates[0].Prob90 == "0000" {
		return DaysFromFrost{}, errors.New("this zipcode is not supported by gardenly")
	}

	daysFromLastFrost, _ := DaysFrom(springDates[0].Prob90, y)
	daysFromFirstFrost, fallDate := DaysFrom(fallDates[0].Prob90, y)
	_, equinox := DaysFrom("0621", y)

	var dff DaysFromFrost

	if fallDate.Before(equinox) {
		daysFromFirstFrost, _ := DaysFrom(springDates[0].Prob90, 1)
		dff = DaysFromFrost{
			FirstFrost: daysFromLastFrost,
			LastFrost:  daysFromFirstFrost,
		}
	} else {
		dff = DaysFromFrost{
			FirstFrost: daysFromFirstFrost,
			LastFrost:  daysFromLastFrost,
		}

	}

	return dff, nil
}

func DaysFrom(t string, y int) (float64, time.Time) {
	year := (time.Now().Year()) + y
	dateString := fmt.Sprintf("%d%s", year, t)

	date, _ := time.Parse("20060102", dateString)
	daysFrom := math.RoundToEven(time.Until(date).Hours() / 24)

	return daysFrom, date
}
