package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/engine"
)

func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	UserId := params["zipcode"]

	res := engine.Engine(UserId)

	var display string

	if res.Msg1 != "" {
		display = res.Msg1
	}
	if res.Msg2 != "" {
		display = display + res.Msg2
	}
	if res.Msg3 != "" {
		display = display + res.Msg3
	}
	if res.Vegetables != nil {
		for _, x := range res.Vegetables {
			display = display + x.CommonName
		}

	}

	if err := json.NewEncoder(w).Encode(display); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler).Methods("GET")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
