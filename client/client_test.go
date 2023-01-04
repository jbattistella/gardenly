package client

import (
	"fmt"
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
		year := time.Now().Year()
		dateString := fmt.Sprintf("%d%s", year, testDate)
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
