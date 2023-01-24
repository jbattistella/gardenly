package database

type Vegetable struct {
	CommonName string `json:"common_name"`
	DTM        int    `json:"dtm"`
	FrostTemp  int    `json:"frost_temp"`
}
