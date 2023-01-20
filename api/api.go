package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/engine"
)

func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages   string
		VegHeader  string
		Vegetables []string
	}

	type ErrPage struct {
		Message error
	}

	params := mux.Vars(r)
	zipCode := params["zipcode"]

	res, err := engine.Engine(zipCode)
	if err != nil {
		er := ErrPage{Message: err}
		t, _ := template.ParseFiles("html/errpage.html")
		if err := t.Execute(w, er); err != nil {
			log.Fatal(err)
		}
		return
	}

	var rep Reply

	msg := res.Msg1 + res.Msg2 + res.Msg3

	var vegNames []string
	for _, v := range res.Vegetables {

		vegNames = append(vegNames, v.CommonName)
	}

	if len(vegNames) == 0 {
		rep = Reply{
			Messages: msg,
		}
	} else {
		rep = Reply{
			Messages:   msg,
			VegHeader:  "You can seed the following:",
			Vegetables: vegNames,
		}
	}

	t, _ := template.ParseFiles("html/gardenly.html")
	if err := t.Execute(w, rep); err != nil {
		log.Fatal(err)
	}
}

func GardenlyHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/homepage.html")
}

func GardenlyHomeSubmission(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	zipcode := r.FormValue("zipcode")

	http.Redirect(w, r, fmt.Sprintf("/gardenly/%s", zipcode), http.StatusFound)

}

func notZero(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s
}

func GardenlyAPI() {

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler).Methods("GET")
	r.HandleFunc("/", GardenlyHome).Methods("GET")
	r.HandleFunc("/", GardenlyHomeSubmission).Methods("POST")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
