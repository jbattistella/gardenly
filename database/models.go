package database

type Vegetable struct {
	// ID         int    `json:"id"`
	CommonName string `json:"common_name"`
	DTM        int    `json:"days_to_maturity"`
	DownToTemp int    `json:"down_to_temp"`
}
