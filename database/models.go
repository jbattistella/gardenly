package database

type Vegetable struct {
	ID         int    `json:"id"`
	CommonName string `json:"common_name"`
	DTM        int    `json:"dtm"`
	DownToTemp int    `json:"frost_temp"`
}
