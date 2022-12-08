package client

type PostalCodeInfo struct {
	PostCode string `json:"post code"`
	Places   []struct {
		Long string `json:"longitude"`
		Lat  string `json:"latitude"`
	} `json:"places"`
}

type StationInfo []struct {
	ID string `json:"id"`
}

type FrostDates []struct {
	SeasonID             string `json:"season_id"`
	TemperatureThreshold string `json:"temperature_threshold"`
	Prob90               string `json:"prob_90"`
}
