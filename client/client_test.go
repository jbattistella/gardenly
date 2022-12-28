package client

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPostalInfo(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		res, err := GetPostalInfo(36525)
		assert.NoError(t, err)
		assert.Equal(t, PostalCodeInfo{
			PostCode:            "36525",
			Country:             "United States",
			CountryAbbreviation: "US",
			Places: []PlaceInfo{PlaceInfo{
				PlaceName:         "Creola",
				Longitude:         "-88.0174",
				State:             "Alabama",
				StateAbbreviation: "AL",
				Latitude:          "30.9013",
			}}}, res)
	})
}

func TestGetStation(t *testing.T) {

	t.Run("OK", func(t *testing.T) {

		testInfo := PostalCodeInfo{
			PostCode:            "36525",
			Country:             "United States",
			CountryAbbreviation: "US",
			Places: []PlaceInfo{PlaceInfo{
				PlaceName:         "Creola",
				Longitude:         "-88.0174",
				State:             "Alabama",
				StateAbbreviation: "AL",
				Latitude:          "30.9013",
			}}}
		res, err := testInfo.GetStation()
		assert.NoError(t, err)
		assert.Equal(t, "10583", res)
	})
}

func TestDaysFrom(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		testDate := "1027"
		dateString := "20221027"
		dateTime, _ := time.Parse("20060102", dateString)
		daysFrom := math.RoundToEven(time.Until(dateTime).Hours() / 24)

		res := DaysFrom(testDate, 0)

		if res != daysFrom {
			t.Errorf("calculating days from (2022-10-27) FAILED. Wanted %v got %v", daysFrom, res)
		} else {
			t.Logf("Calculate days from (2022-10-27) PASSED")
		}

	})
}

//Testing this would need to incorporate a function that calculates daty to, bc
// func TestGetDatesByTemperature(t *testing.T) {

// 	t.Run("OK", func(t *testing.T) {
// 		tdff := DaysFromFrost{
// 			FirstFrost: -33.96006275064815,
// 			LastFrost:  -294.9600627506134,
// 		}

// 		res, err := GetDatesByTemperature("10583", 0)
// 		assert.NoError(t, err)
// 		assert.Equal(t, tdff, res)
// 	})

// }
