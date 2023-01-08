package main

import (
	"encoding/json"
	"fmt"
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
}
func main() {


	var res = engine.AppResponse{}

	res = engine.Engine()

	if res.Msg1 != "" {
		fmt.Println(res.Msg1)
	}
	if res.Msg2 != "" {
		fmt.Println(res.Msg2)
	}
	if res.Msg3 != "" {
		fmt.Println(res.Msg3)
	}
	if res.Vegetables != nil {
		for veg := range res.Vegetables {
			fmt.Println(res.Vegetables[veg].CommonName)
		}
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	// res := engine.Engine()

	//set up router
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	//add handler func

	//update DB funcs
}
