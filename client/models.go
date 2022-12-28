package client

type PostalCodeInfo struct {
	PostCode            string      `json:"post code"`
	Country             string      `json:"country"`
	CountryAbbreviation string      `json:"country abbreviation"`
	Places              []PlaceInfo `json:"places"`
}

type PlaceInfo struct {
	PlaceName         string `json:"place name"`
	Longitude         string `json:"longitude"`
	State             string `json:"state"`
	StateAbbreviation string `json:"state abbreviation"`
	Latitude          string `json:"latitude"`
}

type StationInfo []struct {
	ID string `json:"id"`
}

type FrostDates []struct {
	SeasonID             string `json:"season_id"`
	TemperatureThreshold string `json:"temperature_threshold"`
	Prob90               string `json:"prob_90"`
}

type DaysFromFrost struct {
	FirstFrost float64
	LastFrost  float64
}
