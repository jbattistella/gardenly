package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func GetPostalInfo(zip int) (PostalCodeInfo, error) {

	url := fmt.Sprintf("http://api.zippopotam.us/us/%d", zip)

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

	return info, nil

}

func (pci *PostalCodeInfo) GetStation() (string, error) {

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
	fmt.Println(station)

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

	daysFromLastFrost := DaysFrom(springDates[0].Prob90, y)
	daysFromFirstFrost := DaysFrom(fallDates[0].Prob90, y)

	dff := DaysFromFrost{
		FirstFrost: daysFromFirstFrost,
		LastFrost:  daysFromLastFrost,
	}

	return dff, nil
}

func DaysFrom(t string, y int) float64 {
	year := (time.Now().Year()) + y
	dateString := fmt.Sprintf("%d%s", year, t)

	date, _ := time.Parse("20060102", dateString)
	daysFrom := math.RoundToEven(time.Until(date).Hours() / 24)

	return daysFrom
}
